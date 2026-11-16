package proto

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"testing"

	"github.com/xdrpp/goxdr/xdr"
)

func committedPayloadGoldenFixture() IndexerCommittedBlockPayload {
	var payload IndexerCommittedBlockPayload
	payload.Format_version = 2

	value := &payload.Consensus_value
	value.Protocol = 0x01020304
	value.Prev_hash[0] = 0x10
	value.Prev_hash[31] = 0x2f
	value.Proposal_header.Network_hash[0] = 0x30
	value.Proposal_header.Network_hash[31] = 0x4f
	value.Proposal_header.Config_hash[0] = 0x70
	value.Proposal_header.Config_hash[31] = 0x8f
	value.Proposal_header.Proposer[0] = 0x50
	value.Proposal_header.Proposer[31] = 0x6f
	value.Proposal_header.Block_timepoint = 0x0102030405060708
	value.Proposal_header.Block_height = 123456789
	value.Proposal_header.Storage_hash[0] = 0x90
	value.Proposal_header.Storage_hash[31] = 0xaf
	value.Upgrades = []Upgrade{{0xde, 0xad, 0xbe, 0xef}, {1, 2, 3}}

	var event TxEventLog
	event.Exec_result = ExecResultTxRootFrameReverted
	event.Gas_consumed = 0x0f0e0d0c0b0a0908
	event.Root_output.Status_code = Abort
	*event.Root_output.Abort_buffer() = []byte{0xfa, 0xce, 0xb0, 0x0c, 0x01}
	var emitted TxEmittedEvent
	emitted.Contract_address[0] = 0x41
	emitted.Contract_address[31] = 0x60
	emitted.Num_topics = 1
	emitted.Serialized_event = make([]byte, 35)
	emitted.Serialized_event[0] = 0x61
	emitted.Serialized_event[31] = 0x80
	emitted.Serialized_event[32] = 0xca
	emitted.Serialized_event[33] = 0xfe
	emitted.Serialized_event[34] = 0x01
	event.Events = []TxEmittedEvent{emitted}
	payload.Event_logs = []TxEventLog{event}

	var body ProposedTransaction
	body.Signed_tx.Tx.Invocation.Invoked_address[0] = 0x81
	body.Signed_tx.Tx.Invocation.Invoked_address[31] = 0xa0
	body.Signed_tx.Tx.Invocation.Entrypoint = ContractEntrypointID{'g', 'o', 'l', 'd', 'e', 'n', '0', '1'}
	body.Signed_tx.Tx.Invocation.Gas_limit = 0x1122334455667788
	body.Signed_tx.Tx.Invocation.Calldata = []byte{0x00, 0x11, 0x22, 0x33, 0x44}
	body.Signed_tx.Tx.Uniquifier.Block_expiry_height = 987654321
	body.Signed_tx.Tx.Uniquifier.Nonce = 0x8877665544332211
	body.Signed_tx.Tx.Auth_claim.Source_address_claim().Allowed_keys = 0xa5a5
	body.Signed_tx.Tx.Auth_claim.Source_address_claim().Required_weight = 9
	body.Nondeterministic_proposal_data.Nondeterministic_bits = []byte{0xa5, 0x03}
	body.Nondeterministic_proposal_data.Nondeterministic_bit_count = 10
	body.Nondeterministic_proposal_data.Observed_queue_entry_sequence = []QueueEntry{{0x11, 0x22}, {0x33}}
	payload.Tx_bodies = []ProposedTransaction{body}

	txBytes := marshalGoldenXdr(XDR_Transaction(&body.Signed_tx.Tx))
	txHash := TransactionHash(sha256.Sum256(txBytes))
	payload.Event_logs[0].Tx_hash = txHash
	// ProposalTxs names each transaction by its exact signed encoding, while
	// Txs_hash below stays the txset root over TransactionHash.
	signedTxBytes := marshalGoldenXdr(XDR_SignedTransaction(&body.Signed_tx))
	signedTxHash := SignedTransactionHash(sha256.Sum256(signedTxBytes))
	signedHashBytes := marshalGoldenXdr(XDR_SignedTransactionHash(&signedTxHash))
	nondeterministicBytes := marshalGoldenXdr(
		XDR_NondeterministicProposalData(&body.Nondeterministic_proposal_data),
	)
	proposalTxs := ProposalTxs{}
	proposalTxs.Signed_hash = ParallelBatch{
		Segment_tokens: []uint32{1},
		Segment_bytes:  []uint32{uint32(len(signedHashBytes))},
		Xdr:            signedHashBytes,
	}
	proposalTxs.Nondeterministic_data = ParallelBatch{
		Segment_tokens: []uint32{1},
		Segment_bytes:  []uint32{uint32(len(nondeterministicBytes))},
		Xdr:            nondeterministicBytes,
	}
	value.Proposal_header.Txs_hash = goldenSingleEntryRoot(txHash, nil)
	eventBytes := marshalGoldenXdr(XDR_TxEventLog(&payload.Event_logs[0]))
	value.Proposal_header.Events_hash = goldenSingleEntryRoot(txHash,
		append(eventBytes[sha256.Size:], nondeterministicBytes...))
	value.Proposal_txs_hash = goldenHash("130a7db4c5adbcd1a285f2412a2b226f40115b896835f8b5bc792743e07020c2")
	if got := goldenProposalTxsHash(proposalTxs); got != value.Proposal_txs_hash {
		panic("proposal tx compatibility hash mismatch")
	}

	return payload
}

// Reproduces the fixed shape of a one-entry hogd MerkleTriePrefix. The leaf
// commits directly to its complete 32-byte key and raw value in one SHA-256.
func goldenSingleEntryRoot(key TransactionHash, value []byte) SHA256Hash {
	leaf := []byte{2}
	leaf = binary.LittleEndian.AppendUint64(leaf, 64)
	leaf = append(leaf, key[:]...)
	leaf = binary.LittleEndian.AppendUint64(leaf, uint64(len(value)))
	leaf = append(leaf, value...)

	hashBranch := func(children [16]SHA256Hash) SHA256Hash {
		encoding := []byte{0, 0}
		for _, child := range children {
			encoding = append(encoding, child[:]...)
		}
		return sha256.Sum256(encoding)
	}
	emptyRoot := hashBranch([16]SHA256Hash{})
	var inner, outer [16]SHA256Hash
	inner[key[0]&15] = sha256.Sum256(leaf)
	for i := range outer {
		outer[i] = emptyRoot
	}
	outer[key[0]>>4] = hashBranch(inner)
	return hashBranch(outer)
}

func marshalGoldenXdr(value xdr.XdrType) []byte {
	var encoded bytes.Buffer
	out := xdr.XdrOut{Out: &encoded}
	out.Marshal("", value)
	return encoded.Bytes()
}

func goldenHash(encoded string) SHA256Hash {
	decoded, err := hex.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	var hash SHA256Hash
	copy(hash[:], decoded)
	return hash
}

// appendGoldenCxsUint32 appends n as a cxs unsigned int, in its shortest form:
// n itself up to 62, otherwise a header byte (63, 127 or 191 for one, two or
// four more bytes) and n in that many little-endian bytes.
func appendGoldenCxsUint32(b []byte, n uint32) []byte {
	switch {
	case n <= 62:
		return append(b, byte(n))
	case n <= 0xff:
		return append(b, 63, byte(n))
	case n <= 0xffff:
		return binary.LittleEndian.AppendUint16(append(b, 127), uint16(n))
	default:
		return binary.LittleEndian.AppendUint32(append(b, 191), n)
	}
}

// appendGoldenCxsUint32s appends ns as a cxs array of unsigned ints: its
// length, then each element.
func appendGoldenCxsUint32s(b []byte, ns []uint32) []byte {
	b = appendGoldenCxsUint32(b, uint32(len(ns)))
	for _, n := range ns {
		b = appendGoldenCxsUint32(b, n)
	}
	return b
}

// Reproduces hogd's hash_parallel_batch_compatibility: the hash of
// cxs(DomainSeparator("ParallelBatch")), cxs(segment_tokens),
// cxs(segment_bytes), and the hash of the segment hashes in order.
func goldenParallelBatchHash(batch ParallelBatch) SHA256Hash {
	segmentHashes := sha256.New()
	offset := 0
	for _, segmentBytes := range batch.Segment_bytes {
		end := offset + int(segmentBytes)
		segmentHash := sha256.Sum256(batch.Xdr[offset:end])
		segmentHashes.Write(segmentHash[:])
		offset = end
	}
	preimage := []byte("\x00\x0dParallelBatch")
	preimage = appendGoldenCxsUint32s(preimage, batch.Segment_tokens)
	preimage = appendGoldenCxsUint32s(preimage, batch.Segment_bytes)
	preimage = segmentHashes.Sum(preimage)
	return SHA256Hash(sha256.Sum256(preimage))
}

// Reproduces hogd's proposal_txs_compatibility_hash: the hash of
// cxs(DomainSeparator("ProposalTxs")), the hash of the base with both batches
// emptied, and the hashes of the two batches, in that order.
func goldenProposalTxsHash(txs ProposalTxs) ProposalTxsHash {
	base := txs
	base.Signed_hash = ParallelBatch{}
	base.Nondeterministic_data = ParallelBatch{}
	baseHash := sha256.Sum256(marshalGoldenXdr(XDR_ProposalTxs(&base)))
	signedHash := goldenParallelBatchHash(txs.Signed_hash)
	nondeterministicHash := goldenParallelBatchHash(txs.Nondeterministic_data)
	h := sha256.New()
	h.Write([]byte("\x00\x0bProposalTxs"))
	h.Write(baseHash[:])
	h.Write(signedHash[:])
	h.Write(nondeterministicHash[:])
	return ProposalTxsHash(h.Sum(nil))
}

func TestIndexerCommittedBlockPayloadXdrGolden(t *testing.T) {
	var encoded bytes.Buffer
	out := xdr.XdrOut{Out: &encoded}
	payload := committedPayloadGoldenFixture()
	out.Marshal("", XDR_IndexerCommittedBlockPayload(&payload))

	digest := sha256.Sum256(encoded.Bytes())
	// 588 since SignedTransaction gained the optional pubkey_claim_key: four
	// more bytes for the absent key at the end of the one transaction body.
	// The digest moved again, at the same length, when ProposalTxs began
	// naming transactions by SignedTransactionHash, which changed the
	// fixture's Proposal_txs_hash, again when that hash began combining its
	// parts in order instead of by XOR (hogd#4508), and again when a batch
	// began hashing the cxs of its counts with the hash of its segment
	// hashes. Config_hash, Gas_consumed, and direct key/value commitments now
	// match hogd's committed payload fixture.
	const expectedByteLength = 588
	const expectedDigest = "ea976a0aef83867e589632ab7a175712bf085b08d14a799c089ffe13c89bf643"
	if encoded.Len() != expectedByteLength || hex.EncodeToString(digest[:]) != expectedDigest {
		t.Fatalf("golden mismatch: length=%d digest=%x", encoded.Len(), digest)
	}
}

func TestIndexerOutboxOriginXdrGolden(t *testing.T) {
	origin := IndexerOutboxOrigin{
		Protocol:       ProtocolCommit([]byte("hogp/0123456789abcde")),
		Epoch_sequence: 0x0102030405060708,
		First_height:   123456789,
	}
	origin.First_prev_hash[0] = 0x10
	origin.First_prev_hash[31] = 0x2f
	encoded := marshalGoldenXdr(XDR_IndexerOutboxOrigin(&origin))
	digest := sha256.Sum256(encoded)

	const expectedByteLength = 68
	const expectedDigest = "c31dedbeba469212cb8aee98c74bdd53810031fdd5f2e5b97f68193b50b74fca"
	if len(encoded) != expectedByteLength || hex.EncodeToString(digest[:]) != expectedDigest {
		t.Fatalf("golden mismatch: length=%d digest=%x", len(encoded), digest)
	}
}

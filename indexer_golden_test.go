package proto

import (
	"bytes"
	"crypto/sha256"
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
	event.Root_output.Status_code = Abort
	*event.Root_output.Abort_buffer() = []byte{0xfa, 0xce, 0xb0, 0x0c, 0x01}
	event.Gas_consumed = 0x0f0e0d0c0b0a0908
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
	value.Proposal_header.Txs_hash = goldenHash("cd70edd1d81090195ba2f433aaf7f54454ad1c497be6d16361c2fcadbaac2587")
	value.Proposal_header.Events_hash = goldenHash("285d63c4908a92247c98f3077f3edbc84c2e39c4cb975431f48e4714c9e1b6dc")
	value.Proposal_txs_hash = goldenHash("10326f5184cf53e2ebe46ebecf03650fbf5dd5afb20a245829eb7ad3247a0aa5")
	if got := goldenProposalTxsHash(proposalTxs); got != value.Proposal_txs_hash {
		panic("proposal tx compatibility hash mismatch")
	}

	return payload
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

func goldenParallelBatchHash(batch ParallelBatch) SHA256Hash {
	var result SHA256Hash
	offset := 0
	for _, segmentBytes := range batch.Segment_bytes {
		end := offset + int(segmentBytes)
		segmentHash := sha256.Sum256(batch.Xdr[offset:end])
		for i := range result {
			result[i] ^= segmentHash[i]
		}
		offset = end
	}
	batch.Xdr = nil
	indexHash := sha256.Sum256(marshalGoldenXdr(XDR_ParallelBatch(&batch)))
	for i := range result {
		result[i] ^= indexHash[i]
	}
	return result
}

func goldenProposalTxsHash(txs ProposalTxs) ProposalTxsHash {
	base := txs
	base.Signed_hash = ParallelBatch{}
	base.Nondeterministic_data = ParallelBatch{}
	result := ProposalTxsHash(sha256.Sum256(marshalGoldenXdr(XDR_ProposalTxs(&base))))
	signedHash := goldenParallelBatchHash(txs.Signed_hash)
	nondeterministicHash := goldenParallelBatchHash(txs.Nondeterministic_data)
	for i := range result {
		result[i] ^= signedHash[i] ^ nondeterministicHash[i]
	}
	return result
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
	// fixture's Proposal_txs_hash. Full-key Merkle leaves changed both roots
	// again; Config_hash and Gas_consumed now match the hogd fixture too.
	const expectedByteLength = 588
	const expectedDigest = "0b7f7125e715a913839e4cb34edc793c9083824fc3eeca8a6ed0c19b1cd8efaa"
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

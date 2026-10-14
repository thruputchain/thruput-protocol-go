package proto

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"testing"

	"github.com/xdrpp/goxdr/xdr"
)

// unmarshalXdr decodes data into value, failing the test if the decode panics
// or leaves bytes unread.
func unmarshalXdr(t *testing.T, data []byte, value xdr.XdrType) {
	t.Helper()
	in := bytes.NewReader(data)
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Fatalf("unmarshal %s: %v", value.XdrTypeName(), r)
			}
		}()
		xdr.XdrIn{In: in}.Marshal("", value)
	}()
	if in.Len() != 0 {
		t.Fatalf("unmarshal %s: %d trailing bytes", value.XdrTypeName(), in.Len())
	}
}

// assertRoundTrip checks that decoding encoded into fresh and re-encoding it
// reproduces encoded exactly.
func assertRoundTrip(t *testing.T, encoded []byte, fresh xdr.XdrType) {
	t.Helper()
	unmarshalXdr(t, encoded, fresh)
	if again := marshalGoldenXdr(fresh); !bytes.Equal(again, encoded) {
		t.Fatalf("%s did not round-trip:\n got %x\nwant %x", fresh.XdrTypeName(), again, encoded)
	}
}

func filledBytes(n int, seed byte) []byte {
	out := make([]byte, n)
	for i := range out {
		out[i] = seed + byte(i)
	}
	return out
}

func mlDsa65TestKey() PublicKey {
	var key PublicKey
	key.Scheme = MlDsa65
	copy(key.Mldsa65()[:], filledBytes(1952, 0x40))
	return key
}

func ed25519TestKey() PublicKey {
	var key PublicKey
	key.Scheme = Ed25519
	copy(key.Ed25519()[:], filledBytes(32, 0x20))
	return key
}

func xdrU32(v uint32) []byte {
	return binary.BigEndian.AppendUint32(nil, v)
}

// domainSeparated computes sha256(cxs(DomainSeparator(tag)) || encoded), the
// construction preimage.x describes.
func domainSeparated(tag string, encoded []byte) [32]byte {
	sep := DomainSeparator{Domain_separator: []byte(tag)}
	h := sha256.New()
	h.Write(marshalGoldenXdr(XDR_DomainSeparator(&sep)))
	h.Write(encoded)
	var out [32]byte
	copy(out[:], h.Sum(nil))
	return out
}

func TestPublicKeyMlDsa65Arm(t *testing.T) {
	key := mlDsa65TestKey()
	encoded := marshalGoldenXdr(XDR_PublicKey(&key))
	want := append(xdrU32(1), filledBytes(1952, 0x40)...)
	if !bytes.Equal(encoded, want) {
		t.Fatalf("ML-DSA-65 PublicKey encoding: got %d bytes, want scheme 1 then the 1952-byte key", len(encoded))
	}
	var decoded PublicKey
	assertRoundTrip(t, encoded, &decoded)
	if decoded.Scheme != MlDsa65 || *decoded.Mldsa65() != *key.Mldsa65() {
		t.Fatal("decoded ML-DSA-65 PublicKey differs")
	}
}

func TestSignatureMlDsa65Arm(t *testing.T) {
	var sig Signature
	sig.Scheme = MlDsa65
	// 3309 is not a multiple of 4, so this also covers the XDR padding.
	copy(sig.Mldsa65()[:], filledBytes(3309, 0x07))
	encoded := marshalGoldenXdr(XDR_Signature(&sig))
	want := append(xdrU32(1), filledBytes(3309, 0x07)...)
	want = append(want, 0, 0, 0)
	if !bytes.Equal(encoded, want) {
		t.Fatalf("ML-DSA-65 Signature encoding: got %d bytes, want %d", len(encoded), len(want))
	}
	var decoded Signature
	assertRoundTrip(t, encoded, &decoded)
}

func TestPubkeyAccountPreimage(t *testing.T) {
	for _, key := range []PublicKey{ed25519TestKey(), mlDsa65TestKey()} {
		t.Run(key.Scheme.String(), func(t *testing.T) {
			keyHash := PublicKeyHash(domainSeparated("PublicKey", marshalGoldenXdr(XDR_PublicKey(&key))))

			var claim AuthClaim
			claim.Claim_type = KPubkeySignature
			preimage := claim.Pubkey_preimage()
			preimage.Pubkey_hash = keyHash
			preimage.Nonce = 0x0102030405060708

			// The preimage is the same 40 bytes whatever the scheme, so the
			// claim contracts read is too.
			preimageBytes := marshalGoldenXdr(XDR_PubkeyAccountPreimage(preimage))
			wantPreimage := append(append([]byte{}, keyHash[:]...), 1, 2, 3, 4, 5, 6, 7, 8)
			if !bytes.Equal(preimageBytes, wantPreimage) {
				t.Fatalf("PubkeyAccountPreimage encoding: got %x, want %x", preimageBytes, wantPreimage)
			}
			var decodedPreimage PubkeyAccountPreimage
			assertRoundTrip(t, preimageBytes, &decodedPreimage)

			claimBytes := marshalGoldenXdr(XDR_AuthClaim(&claim))
			if !bytes.Equal(claimBytes, append(xdrU32(uint32(KPubkeySignature)), wantPreimage...)) {
				t.Fatalf("AuthClaim encoding: got %x", claimBytes)
			}
			var decodedClaim AuthClaim
			assertRoundTrip(t, claimBytes, &decodedClaim)

			// Keys of different schemes must not share a hash, which the
			// scheme tag inside cxs(PublicKey) guarantees.
			other := key
			other.Scheme = Ed25519 + MlDsa65 - key.Scheme
			if PublicKeyHash(domainSeparated("PublicKey", marshalGoldenXdr(XDR_PublicKey(&other)))) == keyHash {
				t.Fatal("PublicKeyHash does not commit to the scheme")
			}
		})
	}
}

func TestSignedTransactionPubkeyClaimKey(t *testing.T) {
	var tx SignedTransaction
	tx.Tx.Auth_claim.Claim_type = KPubkeySignature
	tx.Tx.Auth_claim.Pubkey_preimage().Nonce = 7
	var sig DecoratedSignature
	sig.Signature.Scheme = MlDsa65
	copy(sig.Signature.Mldsa65()[:], filledBytes(3309, 0x99))
	tx.Signatures = []DecoratedSignature{sig}
	txHashBefore := sha256.Sum256(marshalGoldenXdr(XDR_Transaction(&tx.Tx)))

	absent := marshalGoldenXdr(XDR_SignedTransaction(&tx))
	if !bytes.HasSuffix(absent, xdrU32(0)) {
		t.Fatal("absent pubkey_claim_key must encode as a trailing zero flag")
	}
	var decodedAbsent SignedTransaction
	assertRoundTrip(t, absent, &decodedAbsent)
	if decodedAbsent.Pubkey_claim_key != nil {
		t.Fatal("absent pubkey_claim_key decoded as present")
	}

	key := mlDsa65TestKey()
	tx.Pubkey_claim_key = &key
	present := marshalGoldenXdr(XDR_SignedTransaction(&tx))
	keyBytes := marshalGoldenXdr(XDR_PublicKey(&key))
	wantSuffix := append(xdrU32(1), keyBytes...)
	if !bytes.Equal(present[:len(absent)-4], absent[:len(absent)-4]) || !bytes.HasSuffix(present, wantSuffix) {
		t.Fatal("present pubkey_claim_key must encode as a trailing one flag then the PublicKey")
	}
	var decodedPresent SignedTransaction
	assertRoundTrip(t, present, &decodedPresent)
	if decodedPresent.Pubkey_claim_key == nil || *decodedPresent.Pubkey_claim_key.Mldsa65() != *key.Mldsa65() {
		t.Fatal("decoded pubkey_claim_key differs")
	}

	// The key sits outside the signed body, so attaching it leaves the
	// transaction hash alone.
	if sha256.Sum256(marshalGoldenXdr(XDR_Transaction(&decodedPresent.Tx))) != txHashBefore {
		t.Fatal("pubkey_claim_key changed the transaction hash")
	}
}

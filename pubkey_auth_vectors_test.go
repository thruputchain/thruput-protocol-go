package proto

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"strconv"
	"testing"
)

// The shared cross-language vectors for pubkey accounts, from the protocol
// submodule. test_vectors/README.md there documents the fields.
const pubkeyAuthVectorsPath = "protocol/test_vectors/pubkey_auth.json"

type hexBytes []byte

func (h *hexBytes) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	decoded, err := hex.DecodeString(s)
	*h = decoded
	return err
}

// decimalU64 is an integer the file writes as a decimal string because it
// can exceed 2^53.
type decimalU64 uint64

func (d *decimalU64) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	v, err := strconv.ParseUint(s, 10, 64)
	*d = decimalU64(v)
	return err
}

type pubkeyAuthVectors struct {
	FormatVersion    int `json:"format_version"`
	DomainSeparators map[string]struct {
		Name string   `json:"name"`
		Cxs  hexBytes `json:"cxs"`
	} `json:"domain_separators"`
	Keys []struct {
		Name          string   `json:"name"`
		Scheme        string   `json:"scheme"`
		SchemeValue   int32    `json:"scheme_value"`
		PublicKey     hexBytes `json:"public_key"`
		PublicKeyCxs  hexBytes `json:"public_key_cxs"`
		PublicKeyXdr  hexBytes `json:"public_key_xdr"`
		PublicKeyHash hexBytes `json:"public_key_hash"`
		Accounts      []struct {
			Nonce                    decimalU64 `json:"nonce"`
			PubkeyAccountPreimageCxs hexBytes   `json:"pubkey_account_preimage_cxs"`
			Address                  hexBytes   `json:"address"`
		} `json:"accounts"`
	} `json:"keys"`
	Signing []struct {
		Name        string   `json:"name"`
		Key         string   `json:"key"`
		NetworkHash hexBytes `json:"network_hash"`
		Transaction struct {
			InvokedAddress    hexBytes   `json:"invoked_address"`
			Entrypoint        hexBytes   `json:"entrypoint"`
			GasLimit          decimalU64 `json:"gas_limit"`
			Calldata          hexBytes   `json:"calldata"`
			BlockExpiryHeight decimalU64 `json:"block_expiry_height"`
			UniquifierNonce   decimalU64 `json:"uniquifier_nonce"`
			AuthClaimType     string     `json:"auth_claim_type"`
			PubkeyHash        hexBytes   `json:"pubkey_hash"`
			AccountNonce      decimalU64 `json:"account_nonce"`
		} `json:"transaction"`
		TransactionXdr          hexBytes `json:"transaction_xdr"`
		SigningPreimage         hexBytes `json:"signing_preimage"`
		TransactionHash         hexBytes `json:"transaction_hash"`
		Signature               hexBytes `json:"signature"`
		SignedTransactionXdr    hexBytes `json:"signed_transaction_xdr"`
		SignedTransactionSha256 hexBytes `json:"signed_transaction_sha256"`
	} `json:"signing"`
	Resigning []struct {
		Name                 string   `json:"name"`
		Of                   string   `json:"of"`
		Signature            hexBytes `json:"signature"`
		SignedTransactionXdr hexBytes `json:"signed_transaction_xdr"`
		TransactionHash      hexBytes `json:"transaction_hash"`
	} `json:"resigning"`
	Negative []struct {
		Name                 string   `json:"name"`
		SignedTransactionXdr hexBytes `json:"signed_transaction_xdr"`
		TransactionHash      hexBytes `json:"transaction_hash"`
	} `json:"negative"`
}

func loadPubkeyAuthVectors(t *testing.T) *pubkeyAuthVectors {
	t.Helper()
	data, err := os.ReadFile(pubkeyAuthVectorsPath)
	if err != nil {
		t.Fatalf("read vectors: %v", err)
	}
	var v pubkeyAuthVectors
	if err := json.Unmarshal(data, &v); err != nil {
		t.Fatalf("parse vectors: %v", err)
	}
	if v.FormatVersion != 1 {
		t.Fatalf("vectors format_version %d, this test reads 1", v.FormatVersion)
	}
	return &v
}

func signatureBytes(sig *Signature) []byte {
	switch sig.Scheme {
	case Ed25519:
		return sig.Ed25519()[:]
	case MlDsa65:
		return sig.Mldsa65()[:]
	}
	return nil
}

func TestPubkeyAuthVectorsKeys(t *testing.T) {
	v := loadPubkeyAuthVectors(t)

	for name, sep := range v.DomainSeparators {
		if !bytes.Equal(cxsDomainSeparator(sep.Name), sep.Cxs) {
			t.Errorf("cxs(DomainSeparator(%q)) = %x, vectors say %x", name, cxsDomainSeparator(sep.Name), sep.Cxs)
		}
	}

	for _, k := range v.Keys {
		t.Run(k.Name, func(t *testing.T) {
			var key PublicKey
			assertRoundTrip(t, k.PublicKeyXdr, &key)
			if int32(key.Scheme) != k.SchemeValue || key.Scheme.String() != k.Scheme {
				t.Fatalf("scheme %v, vectors say %s/%d", key.Scheme, k.Scheme, k.SchemeValue)
			}
			if got := cxsPublicKey(&key); !bytes.Equal(got[1:], k.PublicKey) {
				t.Fatalf("raw key differs from public_key")
			}
			if got := cxsPublicKey(&key); !bytes.Equal(got, k.PublicKeyCxs) {
				t.Fatalf("cxs(PublicKey) = %x, vectors say %x", got, k.PublicKeyCxs)
			}
			keyHash := publicKeyHash(&key)
			if !bytes.Equal(keyHash[:], k.PublicKeyHash) {
				t.Fatalf("PublicKeyHash = %x, vectors say %x", keyHash, k.PublicKeyHash)
			}
			for _, account := range k.Accounts {
				preimage := PubkeyAccountPreimage{Pubkey_hash: keyHash, Nonce: uint64(account.Nonce)}
				if got := cxsPubkeyAccountPreimage(&preimage); !bytes.Equal(got, account.PubkeyAccountPreimageCxs) {
					t.Errorf("nonce %d: cxs(PubkeyAccountPreimage) = %x, vectors say %x", account.Nonce, got, account.PubkeyAccountPreimageCxs)
				}
				if got := pubkeyAccountAddress(&preimage); !bytes.Equal(got[:], account.Address) {
					t.Errorf("nonce %d: address = %x, vectors say %x", account.Nonce, got, account.Address)
				}
			}
		})
	}
}

func TestPubkeyAuthVectorsSigning(t *testing.T) {
	v := loadPubkeyAuthVectors(t)

	keysXdr := map[string][]byte{}
	for _, k := range v.Keys {
		keysXdr[k.Name] = k.PublicKeyXdr
	}
	txHashes := map[string][]byte{}

	for _, s := range v.Signing {
		t.Run(s.Name, func(t *testing.T) {
			var tx Transaction
			assertRoundTrip(t, s.TransactionXdr, &tx)
			want := s.Transaction
			if !bytes.Equal(tx.Invocation.Invoked_address[:], want.InvokedAddress) ||
				!bytes.Equal(tx.Invocation.Entrypoint[:], want.Entrypoint) ||
				tx.Invocation.Gas_limit != uint64(want.GasLimit) ||
				!bytes.Equal(tx.Invocation.Calldata, want.Calldata) ||
				uint64(tx.Uniquifier.Block_expiry_height) != uint64(want.BlockExpiryHeight) ||
				tx.Uniquifier.Nonce != uint64(want.UniquifierNonce) {
				t.Fatalf("decoded transaction differs from the vector's fields: %+v", tx.Invocation)
			}
			if tx.Auth_claim.Claim_type.String() != want.AuthClaimType {
				t.Fatalf("claim type %v, vectors say %s", tx.Auth_claim.Claim_type, want.AuthClaimType)
			}
			preimage := tx.Auth_claim.Pubkey_preimage()
			if !bytes.Equal(preimage.Pubkey_hash[:], want.PubkeyHash) || preimage.Nonce != uint64(want.AccountNonce) {
				t.Fatalf("claim preimage differs from the vector's fields")
			}
			if address := pubkeyAccountAddress(preimage); !bytes.Equal(address[:], want.InvokedAddress) {
				t.Fatalf("claim address %x, invoked address %x", address, want.InvokedAddress)
			}

			signingPreimage := append(cxsDomainSeparator("Transaction"), s.NetworkHash...)
			signingPreimage = append(signingPreimage, marshalGoldenXdr(XDR_Transaction(&tx))...)
			if !bytes.Equal(signingPreimage, s.SigningPreimage) {
				t.Fatalf("signing preimage differs")
			}
			txHash := sha256.Sum256(marshalGoldenXdr(XDR_Transaction(&tx)))
			if !bytes.Equal(txHash[:], s.TransactionHash) {
				t.Fatalf("transaction hash %x, vectors say %x", txHash, s.TransactionHash)
			}
			txHashes[s.Name] = s.TransactionHash

			var signed SignedTransaction
			assertRoundTrip(t, s.SignedTransactionXdr, &signed)
			if sum := sha256.Sum256(s.SignedTransactionXdr); !bytes.Equal(sum[:], s.SignedTransactionSha256) {
				t.Fatalf("signed_transaction_sha256 does not match signed_transaction_xdr")
			}
			if !bytes.Equal(marshalGoldenXdr(XDR_Transaction(&signed.Tx)), s.TransactionXdr) {
				t.Fatalf("signed transaction carries a different Transaction")
			}
			if len(signed.Signatures) != 1 || signed.Signatures[0].Public_key_id != 0 {
				t.Fatalf("want one signature with key id 0, got %+v", signed.Signatures)
			}
			if !bytes.Equal(signatureBytes(&signed.Signatures[0].Signature), s.Signature) {
				t.Fatalf("signature differs from the vector's")
			}
			if signed.Pubkey_claim_key == nil {
				t.Fatalf("pubkey_claim_key absent")
			}
			if !bytes.Equal(marshalGoldenXdr(XDR_PublicKey(signed.Pubkey_claim_key)), keysXdr[s.Key]) {
				t.Fatalf("pubkey_claim_key is not key %s", s.Key)
			}
			if keyHash := publicKeyHash(signed.Pubkey_claim_key); keyHash != preimage.Pubkey_hash {
				t.Fatalf("pubkey_claim_key hashes to %x, claim names %x", keyHash, preimage.Pubkey_hash)
			}
		})
	}

	for _, r := range v.Resigning {
		t.Run(r.Name, func(t *testing.T) {
			var signed SignedTransaction
			assertRoundTrip(t, r.SignedTransactionXdr, &signed)
			if len(signed.Signatures) != 1 || !bytes.Equal(signatureBytes(&signed.Signatures[0].Signature), r.Signature) {
				t.Fatalf("re-signed transaction does not carry the vector's signature")
			}
			txHash := sha256.Sum256(marshalGoldenXdr(XDR_Transaction(&signed.Tx)))
			if !bytes.Equal(txHash[:], r.TransactionHash) || !bytes.Equal(r.TransactionHash, txHashes[r.Of]) {
				t.Fatalf("re-signing changed the transaction hash")
			}
		})
	}
}

// The negatives are rejected by a verifier, not by the bindings, but every one
// is still well-formed XDR that must decode and re-encode unchanged.
func TestPubkeyAuthVectorsNegativeDecode(t *testing.T) {
	v := loadPubkeyAuthVectors(t)
	for _, n := range v.Negative {
		t.Run(n.Name, func(t *testing.T) {
			var signed SignedTransaction
			assertRoundTrip(t, n.SignedTransactionXdr, &signed)
			txHash := sha256.Sum256(marshalGoldenXdr(XDR_Transaction(&signed.Tx)))
			if !bytes.Equal(txHash[:], n.TransactionHash) {
				t.Fatalf("transaction hash %x, vectors say %x", txHash, n.TransactionHash)
			}
		})
	}
}

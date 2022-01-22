package peer

import (
	"bytes"
	"testing"
)

// TestEncrypt tests encrypt/decrypt with multiple key types
func TestEncrypt(t *testing.T) {
	keys := BuildMockKeys(t)
	for ki, privKey := range keys {
		pubKey := privKey.GetPublic()

		peerID, err := IDFromPublicKey(pubKey)
		if err != nil {
			t.Fatal(err.Error())
		}

		// super-secret message
		msg := "Hello to " + peerID.Pretty() + "!"
		dat, err := EncryptToPubKey(pubKey, []byte(msg))
		if err != nil {
			t.Fatal(err.Error())
		}
		t.Logf("keys[%d]: encrypted len: %d", ki, len(dat))

		dec, err := DecryptWithPrivKey(privKey, dat)
		if err != nil {
			t.Fatal(err.Error())
		}
		if bytes.Compare(dec, []byte(msg)) != 0 {
			t.Fatalf("keys[%d]: data did not match: %v != expected %v", ki, dec, []byte(msg))
		}
	}
}

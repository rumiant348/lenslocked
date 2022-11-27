package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"testing"
)

func TestHMAC_Hash(t *testing.T) {
	toHash := []byte("this is my string to hash")
	hashingFunc := hmac.New(sha256.New, []byte("my-secret-key"))
	hashingFunc.Write(toHash)
	b := hashingFunc.Sum(nil)
	expected := base64.URLEncoding.EncodeToString(b)

	h := NewHMAC("my-secret-key")
	actual := h.Hash("this is my string to hash")

	if expected != actual {
		t.Errorf("Hashes not equal, got %s, want %s\n", actual, expected)
	}

}

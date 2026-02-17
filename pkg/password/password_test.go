package password_test

import (
	"testing"

	"github.com/DiaaSaada/crypto-payment-gateway/pkg/password"
)

func TestHash(t *testing.T) {
	pwd := "mysecretpassword"

	hash, err := password.Hash(pwd)
	if err != nil {
		t.Errorf("Hash() unexpected error = %v", err)
	}

	if hash == "" {
		t.Error("Hash() returned empty string")
	}

	if hash == pwd {
		t.Error("Hash() returned plain password instead of hash")
	}
}

func TestVerify_ValidPassword(t *testing.T) {
	pwd := "mysecretpassword"
	hash, _ := password.Hash(pwd)

	if !password.Verify(pwd, hash) {
		t.Error("Verify() failed for valid password")
	}
}

func TestVerify_InvalidPassword(t *testing.T) {
	pwd := "mysecretpassword"
	hash, _ := password.Hash(pwd)

	if password.Verify("wrongpassword", hash) {
		t.Error("Verify() succeeded for invalid password")
	}
}

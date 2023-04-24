package cryptographer_test

import (
	"crypto/aes"
	"testing"

	"github.com/BiPwL/cryptographer-rest-api/internal/app/cryptographer"
)

func TestCryptographer_Encrypt(t *testing.T) {
	testCases := []struct {
		name      string
		text      string
		secretKey string
		bytes     []byte
		errorName error
	}{
		{
			name:      "Valid",
			text:      "hi",
			secretKey: "1234567891011121",
			bytes:     []byte("user@example.org")[:16],
			errorName: nil,
		},
		{
			name:      "short key",
			text:      "hi",
			secretKey: "123",
			bytes:     []byte("user@example.org")[:16],
			errorName: aes.KeySizeError(secretKey),
		},
		{
			name:      "short mail",
			text:      "hi",
			secretKey: "123",
			bytes:     []byte("u@mail.uk")[:16],
			errorName: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encryptedText, err := cryptographer.Encrypt(tc.name, tc.secretKey, tc.bytes)

		})
	}
}

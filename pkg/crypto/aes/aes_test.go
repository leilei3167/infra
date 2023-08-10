package aes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryptor_Encrypt(t *testing.T) {
	salt := "hellpo"
	e, err := NewEncryptor(salt)
	require.NoError(t, err)

	plainText := []byte("hello world")
	encrypted, err := e.Encrypt(plainText)
	require.NoError(t, err)
	require.NotEmpty(t, encrypted)
	t.Logf("encrypted: %x", string(encrypted))
}

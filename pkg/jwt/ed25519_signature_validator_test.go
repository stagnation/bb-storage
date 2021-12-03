package jwt_test

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/buildbarn/bb-storage/pkg/jwt"
	"github.com/stretchr/testify/require"
)

func TestEd25519SignatureValidator(t *testing.T) {
	block, _ := pem.Decode([]byte(`-----BEGIN PUBLIC KEY-----
MCowBQYDK2VwAyEA7fySb/9h7hVH8j1paD5IoLfXj4prjfNLwOPUYKvsTOc=
-----END PUBLIC KEY-----`))
	require.NotNil(t, block)
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	require.NoError(t, err)
	signatureValidator, err := jwt.NewEd25519SignatureValidator(key.(ed25519.PublicKey))
	require.NoError(t, err)

	// Algorithm "HS256" uses HMAC; not Ed25519. Validation should fail.
	require.False(t, signatureValidator.ValidateSignature(
		"HS256",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ",
		[]byte{
			0xb3, 0x57, 0x72, 0xdf, 0xc5, 0xc6, 0x74, 0xba,
			0x79, 0xcc, 0xb6, 0x04, 0x07, 0x65, 0x0e, 0xb7,
			0xd6, 0x65, 0x06, 0x1a, 0x09, 0xed, 0x97, 0xeb,
			0x35, 0x80, 0x06, 0x26, 0xdc, 0x19, 0xec, 0x61,
		}))

	// Ed25519, both with a valid and invalid signature.
	require.True(t, signatureValidator.ValidateSignature(
		"EdDSA",
		"eyJhbGciOiJFZERTQSJ9.eyJpZCI6MX0",
		[]byte{
			0x44, 0x0c, 0x41, 0x01, 0x03, 0xc5, 0x3b, 0x1a,
			0xc2, 0x7e, 0x0a, 0x9b, 0xe7, 0xa7, 0x9a, 0x03,
			0x3f, 0x6e, 0xda, 0x50, 0x72, 0x8f, 0xe5, 0x84,
			0x3b, 0xe2, 0x56, 0x80, 0x91, 0xf7, 0x0b, 0x2c,
			0x75, 0xa3, 0x51, 0xf0, 0x8e, 0x7d, 0x69, 0x03,
			0x63, 0x38, 0x36, 0x02, 0x5e, 0xa1, 0xbf, 0x6f,
			0x6f, 0x63, 0x9a, 0xc3, 0x81, 0x4e, 0x79, 0x81,
			0x6b, 0xeb, 0xfd, 0xf7, 0x5c, 0xc4, 0xe7, 0x01,
		}))
	require.False(t, signatureValidator.ValidateSignature(
		"EdDSA",
		"eyJhbGciOiJFZERTQSJ9.eyJpZCI6MX0",
		[]byte{
			0x04, 0x16, 0xeb, 0x4f, 0xfc, 0x5d, 0x6f, 0x39,
			0xc6, 0xe6, 0xa6, 0xf1, 0x41, 0xa2, 0x55, 0xd1,
			0x37, 0xf3, 0x6a, 0x08, 0x6d, 0x4f, 0x72, 0x73,
			0x11, 0xbe, 0xb1, 0xff, 0x0e, 0x5d, 0x4e, 0x2c,
			0xfe, 0x2f, 0x86, 0xc2, 0x53, 0x2c, 0xb0, 0xa6,
			0x37, 0xe8, 0x32, 0xed, 0x7f, 0x89, 0x5a, 0x70,
			0x30, 0x7a, 0x6f, 0x70, 0x8e, 0xac, 0xda, 0xce,
			0x3a, 0x55, 0x52, 0xa5, 0x1e, 0xdb, 0x07, 0x5c,
		}))
}
package ed25519

import (
	"korthochain/pkg/crypto"
	"korthochain/pkg/crypto/sigs"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundtrip(t *testing.T) {
	priv, err := sigs.Generate(crypto.ED25519)
	require.NoError(t, err)

	msg := []byte("kortho")
	sig, err := sigs.Sign(crypto.ED25519, priv, msg)
	require.NoError(t, err)

	pk, err := sigs.ToPublic(crypto.ED25519, priv)
	require.NoError(t, err)

	err = sigs.Verify(sig, pk, msg)
	require.NoError(t, err)
}

func TestUncompressedFails(t *testing.T) {
	var err error
	//comparessed
	err = sigs.Verify(&crypto.Signature{
		SigType: crypto.ED25519,
		Data:    []byte{0xdf, 0xdf, 0x5e, 0x54, 0xcf, 0xbc, 0x78, 0xc0, 0xe2, 0x3f, 0xca, 0x2f, 0x37, 0x70, 0xc, 0xb2, 0x69, 0xd6, 0x2f, 0x95, 0x5e, 0xff, 0xd5, 0xb2, 0x50, 0x77, 0x52, 0xec, 0x56, 0xc, 0xbb, 0xb2, 0xdf, 0x43, 0xed, 0x4a, 0x7f, 0x6b, 0xe9, 0xe2, 0x35, 0xe1, 0xe6, 0x4d, 0xb4, 0xf7, 0x7b, 0xe7, 0x4f, 0x1, 0x3c, 0x97, 0x51, 0x17, 0xe, 0x8d, 0x79, 0xf5, 0xc7, 0x9f, 0x12, 0xe1, 0x68, 0x7},
	},
		[]byte{0x68, 0x80, 0x73, 0x99, 0x76, 0x33, 0xe6, 0x16, 0x22, 0x59, 0x9e, 0x49, 0x53, 0x35, 0x5b, 0x50, 0xdc, 0x20, 0xda, 0x1d, 0x14, 0xb1, 0x75, 0xff, 0x2e, 0x73, 0x0, 0x6b, 0x82, 0xf3, 0x73, 0x66},
		[]byte{0x6b, 0x6f, 0x72, 0x74, 0x68, 0x6f})

	require.NoError(t, err)

	//comparessed byte changed
	err = sigs.Verify(&crypto.Signature{
		SigType: crypto.ED25519,
		Data:    []byte{0xdf, 0xdf, 0x5e, 0x54, 0xcf, 0xbc, 0x78, 0xc0, 0xe2, 0x3f, 0xca, 0x2f, 0x37, 0x70, 0xc, 0xb2, 0x69, 0xd6, 0x2f, 0x95, 0x5a, 0xff, 0xd5, 0xb2, 0x50, 0x77, 0x52, 0xec, 0x56, 0xc, 0xbb, 0xb2, 0xdf, 0x43, 0xed, 0x4a, 0x7f, 0x6b, 0xe9, 0xe2, 0x35, 0xe1, 0xe6, 0x4d, 0xb4, 0xf7, 0x7b, 0xe7, 0x4f, 0x1, 0x3c, 0x97, 0x51, 0x17, 0xe, 0x8d, 0x79, 0xf5, 0xc7, 0x9f, 0x12, 0xe1, 0x68, 0x7},
	},
		[]byte{0x68, 0x80, 0x73, 0x99, 0x76, 0x33, 0xe6, 0x16, 0x22, 0x59, 0x9e, 0x49, 0x53, 0x35, 0x5b, 0x50, 0xdc, 0x20, 0xda, 0x1d, 0x14, 0xb1, 0x75, 0xff, 0x2e, 0x73, 0x0, 0x6b, 0x82, 0xf3, 0x73, 0x66},
		[]byte{0x6b, 0x6f, 0x72, 0x74, 0x68, 0x6f})

	require.Error(t, err)

	//comparessed suffix
	err = sigs.Verify(&crypto.Signature{
		SigType: crypto.ED25519,
		Data:    []byte{0xdf, 0xdf, 0x5e, 0x54, 0xcf, 0xbc, 0x78, 0xc0, 0xe2, 0x3f, 0xca, 0x2f, 0x37, 0x70, 0xc, 0xb2, 0x69, 0xd6, 0x2f, 0x95, 0x5e, 0xff, 0xd5, 0xb2, 0x50, 0x77, 0x52, 0xec, 0x56, 0xc, 0xbb, 0xb2, 0xdf, 0x43, 0xed, 0x4a, 0x7f, 0x6b, 0xe9, 0xe2, 0x35, 0xe1, 0xe6, 0x4d, 0xb4, 0xf7, 0x7b, 0xe7, 0x4f, 0x1, 0x3c, 0x97, 0x51, 0x17, 0xe, 0x8d, 0x79, 0xf5, 0xc7, 0x9f, 0x12, 0xe1, 0x68, 0x7, 0x55},
	},
		[]byte{0x68, 0x80, 0x73, 0x99, 0x76, 0x33, 0xe6, 0x16, 0x22, 0x59, 0x9e, 0x49, 0x53, 0x35, 0x5b, 0x50, 0xdc, 0x20, 0xda, 0x1d, 0x14, 0xb1, 0x75, 0xff, 0x2e, 0x73, 0x0, 0x6b, 0x82, 0xf3, 0x73, 0x66},
		[]byte{0x6b, 0x6f, 0x72, 0x74, 0x68, 0x6f})
	require.Error(t, err)

}

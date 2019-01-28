package syscoinrpc_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	syscoinrpc "github.com/thebotguys/golang-syscoin-rpc-client"
)

func TestGenerateInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Generating.Generate(1, 1)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGenerateToAddressInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Generating.GenerateToAddress(1, "", 1)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGenerateOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	blockHashes, err := cl.Generating.Generate(11, 0)
	require.NoError(t, err, "Generate: Must not error")

	t.Log("Generate - Generated block hashes:", blockHashes)
}

func TestGenerateToAddressOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	testAddress := "SaaxXq67HhzPbsKNNJBQeQK5qf5Hpv8qq2"

	blockHashes, err := cl.Generating.GenerateToAddress(11, testAddress, 0)
	require.NoError(t, err, "GenerateToAddress: Must not error")

	t.Log("GenerateToAddress - Generated block hashes:", blockHashes)
}

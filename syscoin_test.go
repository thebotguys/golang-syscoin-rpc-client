package syscoinrpc_test

import (
	"github.com/stretchr/testify/require"
	"github.com/thebotguys/golang-syscoin-rpc-client"

	"os"
	"testing"
)

func TestGetGovernanceInfoInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Syscoin.GetGovernanceInfo()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetGovernanceInfoOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	info, err := cl.Syscoin.GetGovernanceInfo()
	require.NoError(t, err, "GetGovernanceInfo : must not error")

	t.Log("GetGovernanceInfo :", info)
}

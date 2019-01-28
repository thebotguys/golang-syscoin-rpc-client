package syscoinrpc_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	syscoinrpc "github.com/thebotguys/golang-syscoin-rpc-client"
)

func TestCreateAuxBlockInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Mining.CreateAuxBlock("")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetAuxBlockInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Mining.GetAuxBlock("invalid", "invalid")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetBlockTemplateInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Mining.GetBlockTemplate("", nil, nil)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestCreateAuxBlockOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	testCoinbase := "SaaxXq67HhzPbsKNNJBQeQK5qf5Hpv8qq2"

	info, err := cl.Mining.CreateAuxBlock(testCoinbase)
	require.NoError(t, err, "CreateAuxBlock: must not error")

	t.Log(info)
}

func TestGetAuxBlockOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	info, err := cl.Mining.GetAuxBlock("", "")
	require.NoError(t, err, "GetAuxBlock - no param: must not error")
	t.Log(info)

	_, err = cl.Mining.GetAuxBlock("", "invalid")
	require.EqualError(t, err, "Cannot pass hash or serializedAuxPow alone, must pass both or none", "GetAuxBlock - only one arg: must error")

	_, err = cl.Mining.GetAuxBlock("invalid", "invalid")
	require.Error(t, err, "GetAuxBlock - invalid args: must error")
}

func TestGetBlockTemplateOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	info, err := cl.Mining.GetBlockTemplate("", nil, nil)
	require.NoError(t, err, "GetBlockTemplate - no param: must not error")
	t.Log(info)

	_, err = cl.Mining.GetBlockTemplate("template", nil, nil)
	require.NoError(t, err, "GetBlockTemplate - mode only: must not error")

	_, err = cl.Mining.GetBlockTemplate("template", nil, []string{"segwit"})
	require.NoError(t, err, "GetBlockTemplate mode + rules: must not error")
}
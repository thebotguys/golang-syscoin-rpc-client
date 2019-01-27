package syscoinrpc_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thebotguys/golang-syscoin-rpc-client"
)

func TestDebugInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	err = cl.Control.Debug("0")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetHelpInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Control.GetHelp("")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetMemoryInfoInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Control.GetMemoryInfo()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestDebugOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	err = cl.Control.Debug("0")
	require.NoError(t, err, "Debug: Must not error")
}

func TestGetHelpOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	expectedText := `"help ( \"command\" )\n\nList all commands, or get help for a specified command.\n\nArguments:\n1. \"command\"     (string, optional) The command to get help on\n\nResult:\n\"text\"     (string) The help text\n"`

	text, err := cl.Control.GetHelp("help")
	require.NoError(t, err, "Debug: Must not error")
	require.Equal(t, expectedText, text, "GetHelp: `help` message must equal to the expected one")
}

func TestGetMemoryInfoOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	info, err := cl.Control.GetMemoryInfo()
	require.NoError(t, err, "GetMemoryInfo: must not error")

	t.Log("GetMemoryInfo :", info)
}

func TestStopServerOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	t.Skip("This test would terminate the node, remove this skip to test it anyway")

	err = cl.Control.StopServer()
	require.NoError(t, err, "StopServer: must not error")
}

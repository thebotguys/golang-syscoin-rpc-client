package syscoinrpc_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thebotguys/golang-syscoin-rpc-client"
)

func TestGetHelpInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Control.GetHelp("")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestLoggingInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Control.Logging([]string{}, []string{})
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetMemoryInfoInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Control.GetMemoryInfo()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestStopServerInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	err = cl.Control.StopServer()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetUptimeInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Control.GetUptime()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetHelpOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	expectedText := `"help ( \"command\" )\n\nList all commands, or get help for a specified command.\n\nArguments:\n1. \"command\"     (string, optional) The command to get help on\n\nResult:\n\"text\"     (string) The help text\n"`

	text, err := cl.Control.GetHelp("help")
	require.NoError(t, err, "Debug: Must not error")
	require.Equal(t, expectedText, text, "GetHelp: `help` message must equal to the expected one")
}

func TestLoggingOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	includes := []string{"syscoin"}
	excludes := []string{"tor"}

	loggings, err := cl.Control.Logging(nil, nil)
	require.NoError(t, err, "Loggings - no param: must not error")
	t.Log("Loggings - no param:", loggings)

	loggings, err = cl.Control.Logging(includes, nil)
	require.NoError(t, err, "Loggings - only include: must not error")
	t.Log("Loggings - only include:", loggings)

	loggings, err = cl.Control.Logging(includes, excludes)
	require.NoError(t, err, "Loggings - include + exclude: must not error")
	t.Log("Loggings - include + exclude:", loggings)

	loggings, err = cl.Control.Logging(nil, excludes)
	require.EqualError(t, err, syscoinrpc.ErrLoggingFilters.Error(), "Loggings - only exclude: must error")
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

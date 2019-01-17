package syscoinrpc_test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	syscoinrpc "github.com/thebotguys/golang-syscoin-rpc-client"
)

const invalidURL = `http://invalid.url`

func init() {
	http.DefaultClient.Timeout = time.Second // for quick tests, can be changed for more reliability.
}

func TestGetAddressBalanceInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.AddressIndex.GetAddressBalance(nil, false)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetAddressDeltasInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.AddressIndex.GetAddressDeltas(nil, 0, 0)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetAddressMemPoolInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.AddressIndex.GetAddressMemPool(nil)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetAddressTxIDsInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.AddressIndex.GetAddressTxIDs(nil, 0, 0)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetAddressUTXOsInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.AddressIndex.GetAddressUTXOs(nil)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetAddressBalanceOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	testAddresses := []string{"SU8UsT1LLMR8XvFFehbovp1L4P51xmnetr", "Saqi3gtjyVEndehH4PWc7bRR4ayzAZhrnj", "ShmVjaK4bW2LfhbMyx253QvyDbjD1h71yx"}
	testSeparatedOutput := false

	bal, err := cl.AddressIndex.GetAddressBalance(testAddresses, testSeparatedOutput)
	require.NoError(t, err, "Must not error on valid URL, check if the node is running")

	balJSON, _ := json.Marshal(bal)
	t.Log(string(balJSON))
}

func TestGetAddressDeltasOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	testAddresses := []string{"SU8UsT1LLMR8XvFFehbovp1L4P51xmnetr", "Saqi3gtjyVEndehH4PWc7bRR4ayzAZhrnj", "ShmVjaK4bW2LfhbMyx253QvyDbjD1h71yx"}
	testStart := 0
	testEnd := 0

	deltas, err := cl.AddressIndex.GetAddressDeltas(testAddresses, testStart, testEnd)
	require.NoError(t, err, "Must not error on valid URL, check if the node is running")

	deltasJSON, _ := json.Marshal(deltas)
	t.Log(string(deltasJSON))
}

func TestGetAddressMemPoolOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	testAddresses := []string{"SU8UsT1LLMR8XvFFehbovp1L4P51xmnetr", "Saqi3gtjyVEndehH4PWc7bRR4ayzAZhrnj", "ShmVjaK4bW2LfhbMyx253QvyDbjD1h71yx"}

	memPoolDeltas, err := cl.AddressIndex.GetAddressMemPool(testAddresses)
	require.NoError(t, err, "Must not error on valid URL, check if the node is running")

	memPoolDeltasJSON, _ := json.Marshal(memPoolDeltas)
	t.Log(string(memPoolDeltasJSON))
}

func TestGetAddressTxIDsOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	testAddresses := []string{"SU8UsT1LLMR8XvFFehbovp1L4P51xmnetr", "Saqi3gtjyVEndehH4PWc7bRR4ayzAZhrnj", "ShmVjaK4bW2LfhbMyx253QvyDbjD1h71yx"}
	testStart := 0
	testEnd := 0

	addressTxIDs, err := cl.AddressIndex.GetAddressTxIDs(testAddresses, testStart, testEnd)
	require.NoError(t, err, "Must not error on valid URL, check if the node is running")

	addressTxIDsJSON, _ := json.Marshal(addressTxIDs)
	t.Log(string(addressTxIDsJSON))
}

func TestGetAddressUTXOsOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	testAddresses := []string{"SU8UsT1LLMR8XvFFehbovp1L4P51xmnetr", "Saqi3gtjyVEndehH4PWc7bRR4ayzAZhrnj", "ShmVjaK4bW2LfhbMyx253QvyDbjD1h71yx"}

	addressUTXOs, err := cl.AddressIndex.GetAddressUTXOs(testAddresses)
	require.NoError(t, err, "Must not error on valid URL, check if the node is running")

	addressUTXOsJSON, _ := json.Marshal(addressUTXOs)
	t.Log(string(addressUTXOsJSON))
}

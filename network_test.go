package syscoinrpc_test

import (
	"github.com/stretchr/testify/require"
	"github.com/thebotguys/golang-syscoin-rpc-client"

	"os"
	"testing"
)

func TestAddNodeInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	err = cl.Network.AddNode("", "")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestClearBannedIPsInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	err = cl.Network.ClearBannedIPs()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestDisconnectNodeInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	err = cl.Network.DisconnectNode("", "")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetAddedNodeInfoInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Network.GetAddedNodeInfo("")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetConnectionCountInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Network.GetConnectionCount()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetNetworkTotalsInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Network.GetNetworkTotals()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetNetworkInfoInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Network.GetNetworkInfo()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetPeerInfoInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Network.GetPeerInfo()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestListBannedIPsInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Network.ListBannedIPs()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestPingInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	err = cl.Network.Ping()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestSetBanInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	err = cl.Network.SetBan("", "", 0, false)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestSetNetworkActiveInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	err = cl.Network.SetNetworkActive(false)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestAddNodeOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	t.Skip("This test is not ready yet, please see TODO")

	testNode := "To be filled" // TODO: Fill with proper test data when sys4 testnet goes alive.
	testCommand := "To be filled"
	err = cl.Network.AddNode(testNode, testCommand)
	require.NoError(t, err, "AddNode : must not error")
}

func TestClearBannedIPsOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	t.Skip("This test is not ready yet, please see TODO")

	err = cl.Network.ClearBannedIPs()
	require.NoError(t, err, "ClearBannedIPs : must not error")
}

func TestDisconnectNodeOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	t.Skip("This test is not ready yet, please see TODO")

	testNode := "To be filled" // TODO: Fill with proper test data when sys4 testnet goes alive.
	testNodeID := "To be filled"
	err = cl.Network.DisconnectNode(testNode, testNodeID)
	require.NoError(t, err, "DisconnectNode : must not error")
}

func TestGetAddedNodeInfoOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	t.Skip("This test is not ready yet, please see TODO")

	testNode := "To be filled" // TODO: Fill with proper test data when sys4 testnet goes alive.

	_, err = cl.Network.GetAddedNodeInfo("")
	require.NoError(t, err, "GetAddedNodeInfo : must not error")

	_, err = cl.Network.GetAddedNodeInfo(testNode)
	require.NoError(t, err, "GetAddedNodeInfo : must not error")
}

func TestGetConnectionCountOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	count, err := cl.Network.GetConnectionCount()
	require.NoError(t, err, "GetConnectionCount : must not error")

	t.Log("GetConnectionCount :", count)
}

func TestGetNetworkTotalsOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	totals, err := cl.Network.GetNetworkTotals()
	require.NoError(t, err, "GetNetworkTotals : must not error")

	t.Log("GetNetworkTotals :", totals)
}

func TestGetNetworkInfoOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	info, err := cl.Network.GetNetworkInfo()
	require.NoError(t, err, "GetNetworkInfo : must not error")

	t.Log("GetNetworkInfo :", info)
}

func TestGetPeerInfoOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	info, err := cl.Network.GetPeerInfo()
	require.NoError(t, err, "GetPeerInfo : must not error")

	t.Log("GetPeerInfo :", info)
}

func TestListBannedIPsOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	ips, err := cl.Network.ListBannedIPs()
	require.NoError(t, err, "ListBannedIPs : must not error")

	t.Log("ListBannedIPs :", ips)
}

func TestPingOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	err = cl.Network.Ping()
	require.NoError(t, err, "Ping : must not error")
}

func TestSetBanOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	t.Skip("This test is not ready yet, please see TODO")

	testIP := "To be filled" // TODO: Fill with proper test data when sys4 testnet goes alive.
	testCommand := "To be filled"
	testBanTime := uint64(0)
	testIsAbsoluteTimestamp := false

	err = cl.Network.SetBan(testIP, testCommand, testBanTime, testIsAbsoluteTimestamp)
	require.NoError(t, err, "SetBan : must not error")
}

func TestSetNetworkActiveOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	err = cl.Network.SetNetworkActive(true)
	require.NoError(t, err, "SetNetworkActive : must not error")
}

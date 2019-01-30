// Package syscoinrpc contains the implementation of a syscoin JSON-RPC client.
package syscoinrpc

import (
	"net/http"
)

const (
	// LocalNodeURL represents a valid testnet node URL.
	LocalNodeURL string = "http://127.0.0.1:8370"
)

// Client represents a syscoin JSON-RPC over HTTP client.
type Client struct {
	url        string            // The url of the node to connect to.
	user       string            // The RPC Username.
	pass       string            // The RPC Password.
	httpClient *http.Client      // The JSON-RPC over HTTP sub client.
	Blockchain *BlockchainClient // The client of `blockchain` calls.
	Control    *ControlClient    // The client of `control` calls.
	Generating *GeneratingClient // The client of `generating` calls.
	Mining     *MiningClient     // The client of `mining` calls.
	Network    *NetworkClient    // The client of `network` calls.
}

// NewClient creates a new client object.
func NewClient(url string, rpcUser string, rpcPassword string) (*Client, error) {
	cl := &Client{
		url:        url,
		user:       rpcUser,
		pass:       rpcPassword,
		httpClient: http.DefaultClient,
	}

	cl.Blockchain = &BlockchainClient{cl}
	cl.Control = &ControlClient{cl}
	cl.Generating = &GeneratingClient{cl}
	cl.Mining = &MiningClient{cl}
	cl.Network = &NetworkClient{cl}

	return cl, nil
}

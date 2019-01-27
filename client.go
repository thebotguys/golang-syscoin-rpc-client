// Package syscoinrpc contains the implementation of a syscoin JSON-RPC client.
package syscoinrpc

import (
	"net/http"
)

const (
	// LocalNodeURL represents a valid testnet node URL.
	LocalNodeURL string = "http://127.0.0.1:8368"
)

// Client represents a syscoin JSON-RPC over HTTP client.
type Client struct {
	url             string                 // The url of the node to connect to.
	user            string                 // The RPC Username.
	pass            string                 // The RPC Password.
	httpClient      *http.Client           // The JSON-RPC over HTTP sub client.
	AddressIndex    *addressIndexClient    // The client of `addressindex` calls.
	BlockchainIndex *blockchainIndexClient // The client of `blockchaindex` calls.
	Control         *controlClient         // The client of `control` calls.
	Generating      *generatingClient      // The client of `generating` calls.
}

// NewClient creates a new client object.
func NewClient(url string, rpcUser string, rpcPassword string) (*Client, error) {
	cl := &Client{
		url:        url,
		user:       rpcUser,
		pass:       rpcPassword,
		httpClient: http.DefaultClient,
	}

	cl.AddressIndex = &addressIndexClient{cl}
	cl.BlockchainIndex = &blockchainIndexClient{cl}
	cl.Control = &controlClient{cl}
	cl.Generating = &generatingClient{cl}

	return cl, nil
}

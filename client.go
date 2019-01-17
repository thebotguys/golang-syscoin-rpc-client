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
	url          string              // The url of the node to connect to.
	user         string              // The RPC Username.
	pass         string              // The RPC Password.
	httpClient   *http.Client        // The JSON-RPC over HTTP sub client.
	AddressIndex *addressIndexClient // The client of addressindex calls.
}

// NewClient creates a new client object.
func NewClient(url string, rpcUser string, rpcPassword string) (*Client, error) {
	// curl --user "d2838fb4615effa2:5c2e4ead196bd23114e3b84f16d021e7" --data-binary '{"jsonrpc": "1.0", "id":"curltest", "method": "getaddressbalance", "params": [{"addresses": ["SU8UsT1LLMR8XvFFehbovp1L4P51xmnetr", "Saqi3gtjyVEndehH4PWc7bRR4ayzAZhrnj", "ShmVjaK4bW2LfhbMyx253QvyDbjD1h71yx"], "separated_output" : true}] }' http://127.0.0.1:8368
	// {"result":{"balance":10.10270000,"received":10.10360000},"error":null,"id":"curltest"}

	cl := &Client{
		url:        url,
		user:       rpcUser,
		pass:       rpcPassword,
		httpClient: http.DefaultClient,
	}

	cl.AddressIndex = &addressIndexClient{cl}

	return cl, nil
}

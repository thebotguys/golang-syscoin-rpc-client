package syscoinrpc

import "encoding/json"

// GeneratingClient wraps all `generating` related functions.
type GeneratingClient struct {
	c *Client // The binded client, must not be nil.
}

func (gc *GeneratingClient) do(method string, params ...interface{}) (json.RawMessage, error) {
	return gc.c.do(method, params...)
}

// Generate mines instantly (before RPC call returns) a specified
// number of blocks to a wallet address in the node.
//
// Returns the hashes of the generated blocks.
//
//     nBlocks  : The number of blocks to generate.
//     maxTries : The number of iterations to try (default = 0 -> 1000000 iterations).
func (gc *GeneratingClient) Generate(nBlocks uint64, maxTries uint64) ([]string, error) {
	if maxTries == 0 {
		maxTries = 1000000
	}

	response, err := gc.do("generate", nBlocks, maxTries)
	if err != nil {
		return nil, err
	}

	var hashes []string
	err = json.Unmarshal(response, &hashes)
	if err != nil {
		return nil, err
	}

	return hashes, nil
}

// GenerateToAddress mines instantly (before RPC call returns) a specified
// number of blocks to a the specified wallet address.
//
// Returns the hashes of the generated blocks.
//
//     nBlocks  : The number of blocks to generate.
//     address  : The address to send the newly generated Syscoin to.
//     maxTries : The number of iterations to try (default = 0 -> 1000000 iterations).
func (gc *GeneratingClient) GenerateToAddress(nBlocks uint64, address string, maxTries uint64) ([]string, error) {
	if maxTries == 0 {
		maxTries = 1000000
	}

	response, err := gc.do("generatetoaddress", nBlocks, address, maxTries)
	if err != nil {
		return nil, err
	}

	var hashes []string
	err = json.Unmarshal(response, &hashes)
	if err != nil {
		return nil, err
	}

	return hashes, nil
}

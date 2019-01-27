package syscoinrpc

import "encoding/json"

// controlClient wraps all `control` related functions.
type controlClient struct {
	c *Client // The binded client, must not be nil.
}

func (cc *controlClient) do(method string, params ...interface{}) (json.RawMessage, error) {
	return cc.c.do(method, params...)
}

// Debug sets the debug flags
//
// Allowed flags are the following:
//     0 : no flag
//     1 : all flags
//     Selective flags: addrman|alert|bench|coindb|db|lock|rand|rpc|selectcoins|mempool|mempoolrej|net|proxy|prune|http|libevent|tor|zmq|syscoin|privatesend|instantsend|masternode|spork|keepass|mnpayments|gobject
//
// It is possible to set multiple flags by chaining them with a '+' character (e.g. "addrman+alert")
func (cc *controlClient) Debug(flags string) error {
	_, err := cc.do("debug", flags)
	return err
}

// GetHelp returns the help text for the specified command.
func (cc *controlClient) GetHelp(commandName string) (string, error) {
	response, err := cc.do("help", commandName)
	if err != nil {
		return "", err
	}
	return string(response), nil
}

// MemoryInfo represents general information about memory usage.
type MemoryInfo struct {
	// Locked is the data about locked memory usage.
	Locked LockedMemoryInfo `json:"locked,required"`
}

// LockedMemoryInfo represents general information about locked memory usage.
type LockedMemoryInfo struct {
	// Used is the number of bytes used.
	Used uint64 `json:"used,required"`
	// Free is the number of bytes available in the current arenas.
	Free uint64 `json:"free,required"`
	// Total is the total number of managed bytes.
	Total uint64 `json:"total,required"`
	// Locked is the amount of bytes that succeeded locking.
	// If this number is smaller than total, locking pages
	// failed at some point and key data could be swapped to disk.
	Locked uint64 `json:"locked,required"`
	// UsedChunks is the number of allocated chunks.
	UsedChunks uint64 `json:"chunks_used,required"`
	// FreeChunks is the number of unused chunks.
	FreeChunks uint64 `json:"chunks_free,required"`
}

// GetMemoryInfo return general information about memory usage.
func (cc *controlClient) GetMemoryInfo() (*MemoryInfo, error) {
	response, err := cc.do("getmemoryinfo")
	if err != nil {
		return nil, err
	}

	var info MemoryInfo
	err = json.Unmarshal(response, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// StopServer stops the running syscoin server node.
func (cc *controlClient) StopServer() error {
	_, err := cc.do("stop")
	return err
}

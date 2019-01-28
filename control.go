package syscoinrpc

import (
	"encoding/json"
	"errors"
	"strconv"
)

// ControlClient wraps all `control` related functions.
type ControlClient struct {
	c *Client // The binded client, must not be nil.
}

func (cc *ControlClient) do(method string, params ...interface{}) (json.RawMessage, error) {
	return cc.c.do(method, params...)
}

// GetHelp returns the help text for the specified command.
func (cc *ControlClient) GetHelp(commandName string) (string, error) {
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
func (cc *ControlClient) GetMemoryInfo() (*MemoryInfo, error) {
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

// ErrLoggingFilters is an error when putting invalid filters to `logging` call.
var ErrLoggingFilters = errors.New("Must define include AND exclude fields, or none of them, or include only")

// Logging gets and sets the logging configuration.
//
// When called without an argument, returns the list of categories with status that are currently being debug logged or not.
//
// When called with arguments, adds or removes categories from debug logging and return the lists above.
//
// The arguments are evaluated in order "include", "exclude".
//
// If an item is both included and excluded, it will thus end up being excluded.
//
// The valid logging categories are: net, tor, mempool, http, bench, zmq, db, rpc, estimatefee, addrman, selectcoins, reindex, cmpctblock, rand, prune, proxy, mempoolrej, libevent, coindb, qt, leveldb, threadpool, masternode, gobject, mnpayments, mnsync, spork, syscoin
//
// In addition, the following are available as category names with special meanings:
//
//     "all",  "1" : represent all logging categories.
//     "none", "0" : even if other logging categories are specified, ignore all of them.
func (cc *ControlClient) Logging(include []string, exclude []string) (map[string]bool, error) {
	params := make([]interface{}, 0, 2)
	if include != nil {
		params = append(params, include)
		if exclude != nil {
			params = append(params, exclude)
		}
	} else if include == nil && exclude != nil {
		return nil, ErrLoggingFilters
	}

	response, err := cc.do("logging", params...)
	if err != nil {
		return nil, err
	}

	var loggings map[string]bool
	err = json.Unmarshal(response, &loggings)
	if err != nil {
		return nil, err
	}

	return loggings, nil
}

// StopServer stops the running syscoin server node.
func (cc *ControlClient) StopServer() error {
	_, err := cc.do("stop")
	return err
}

// GetUptime returns the total uptime of the server.
func (cc *ControlClient) GetUptime() (uint64, error) {
	response, err := cc.do("uptime")
	if err != nil {
		return 0, err
	}

	uptime, err := strconv.ParseUint(string(response), 10, 64)
	if err != nil {
		return 0, err
	}

	return uptime, nil
}

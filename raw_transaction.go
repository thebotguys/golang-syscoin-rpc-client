package syscoinrpc

import (
	"encoding/json"
)

// RawTransactionClient wraps all `rawtransaction` related functions.
type RawTransactionClient struct {
	c *Client // The bound client, must not be nil.
}

func (rtc *RawTransactionClient) do(method string, params ...interface{}) (json.RawMessage, error) {
	return rtc.c.do(method, params...)
}

// TODO

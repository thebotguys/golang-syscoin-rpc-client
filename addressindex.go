package syscoinrpc

import (
	"encoding/json"
	"time"
)

// addressIndexClient wraps all addressindex related functions.
type addressIndexClient struct {
	c *Client // The binded client, must not be nil.
}

// getAddressBalancePayload represents the payload of a `getAddressBalance` call.
type getAddressBalancePayload struct {
	// Addresses is the list of the base58check encoded addresses.
	Addresses []string `json:"addresses,required"`
	// SeparatedOutput if set to true, will return balances of the addresses passed
	// in as an array instead of the summed balance. Default is false.
	SeparatedOutput bool `json:"separated_output,omitempty"`
}

// addressBalance represents a GetAddressBalance response.
type addressBalance struct {
	Balance  float64 `json:"balance,required"`
	Received float64 `json:"received,required"`
}

// GetAddressBalance returns the balance for an address(es) (requires addressindex to be enabled).
//     addresses       : The array of base58check encoded addresses
//     separatedOutput : If set to true, will return balances of the
//                       addresses passed in as an array instead of
//                       the summed balance [Optional]
func (aic *addressIndexClient) GetAddressBalance(addresses []string, separatedOutput bool) (*addressBalance, error) {
	payload := getAddressBalancePayload{
		Addresses:       addresses,
		SeparatedOutput: separatedOutput,
	}

	var response addressBalance
	res, err := aic.c.do(aic.c.url, "getaddressbalance", payload)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// getAddressDeltasPayload represents the payload of a `getAddressDeltas` call.
type getAddressDeltasPayload struct {
	// Addresses is the list of the base58check encoded addresses.
	Addresses []string `json:"addresses,required"`
	// Start is the start block height.
	Start int `json:"start,omitempty"`
	// End is the end block height.
	End int `json:"end,omitempty"`
}

// addressDeltas represents a GetAddressDeltas response.
type addressDeltas []addressDelta

// addressDelta represents a single address delta.
type addressDelta struct {
	// Satoshis is the difference of satoshis.
	Satoshis int `json:"satoshis,required"`
	// TxID is the related txid
	TxID string `json:"txid,required"`
	// Index is the related input or output index.
	Index int `json:"index,required"`
	// BlockIndex is the related block index.
	BlockIndex int `json:"blockindex,required"`
	// Height is the block height.
	Height int `json:"height,required"`
	// Address is the base58check encoded address.
	Address string `json:"address,required"`
}

// GetAddressDeltas returns all changes for an address (requires addressindex to be enabled).
//     addresses: The array of base58check encoded addresses
//     start    : The start block height [Optional]
//     end      : The end block height   [Optional]
func (aic *addressIndexClient) GetAddressDeltas(addresses []string, start int, end int) (addressDeltas, error) {
	payload := getAddressDeltasPayload{
		Addresses: addresses,
		Start:     start,
		End:       end,
	}

	var response addressDeltas
	res, err := aic.c.do(aic.c.url, "getaddressdeltas", payload)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// getAddressMemPoolPayload represents the payload of a `getAddressDeltas` call.
type getAddressMemPoolPayload struct {
	// Addresses is the list of the base58check encoded addresses.
	Addresses []string `json:"addresses,required"`
}

// addressMemPool represents a GetAddressMemPool response.
type addressMemPoolDeltas []addressMemPoolDelta

// addressMemPoolDelta represents a single address mempool delta.
type addressMemPoolDelta struct {
	// Address is the base58check encoded address.
	Address string `json:"address,required"`
	// TxID is the related txid
	TxID string `json:"txid,required"`
	// Index is the related input or output index.
	Index int `json:"index,required"`
	// Satoshis is the difference of satoshis.
	Satoshis int `json:"satoshis,required"`
	// Timestamp is the time the transaction entered the mempool as UNIX timestamp.
	Timestamp time.Time `json:"timestamp,required"`
	// PrevTxIn is the previous txid (if spending).
	PrevTxIn string `json:"prevtxin,required"`
	// PrevTxOut is the previous transaction output index (if spending).
	PrevTxOut string `json:"prevtxout,required"`
}

// GetAddressMemPool returns all mempool deltas for an address (requires addressindex to be enabled).
//     addresses: The array of base58check encoded addresses
func (aic *addressIndexClient) GetAddressMemPool(addresses []string) (addressMemPoolDeltas, error) {
	payload := getAddressMemPoolPayload{addresses}

	var response addressMemPoolDeltas
	res, err := aic.c.do(aic.c.url, "getaddressmempool", payload)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// getAddressTxIDsPayload represents the payload of a `GetAddressTxIDs` call.
type getAddressTxIDsPayload struct {
	// Addresses is the list of the base58check encoded addresses.
	Addresses []string `json:"addresses,required"`
	// Start is the start block height.
	Start int `json:"start,omitempty"`
	// End is the end block height.
	End int `json:"end,omitempty"`
}

// addressTxIDs represents the payload of a `GetAddressTxIDs` call.
type addressTxIDs []string

// GetAddressTxIDs returns the txids for an address(es) (requires addressindex to be enabled).
//     addresses: The array of base58check encoded addresses
//     start    : The start block height [Optional]
//     end      : The end block height   [Optional]
func (aic *addressIndexClient) GetAddressTxIDs(addresses []string, start int, end int) (addressTxIDs, error) {
	payload := getAddressTxIDsPayload{
		Addresses: addresses,
		Start:     start,
		End:       end,
	}

	var response addressTxIDs
	res, err := aic.c.do(aic.c.url, "getaddresstxids", payload)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// getAddressUTXOsPayload represents the payload of a `GetAddressUTXOs` call.
type getAddressUTXOsPayload struct {
	// Addresses is the list of the base58check encoded addresses.
	Addresses []string `json:"addresses,required"`
}

// addressUTXOs represents the payload of a `GetAddressUTXOs` call.
type addressUTXOs []addressUTXO

// addressTxID represents a single transaction ID.
type addressUTXO struct {
	// Address is the base58check encoded address.
	Address string `json:"address,required"`
	// TxID is the output txid
	TxID string `json:"txid,required"`
	// OutputIndex is the output index.
	OutputIndex int `json:"outputindex,required"`
	//Script is the script hex encoded.
	Script string `json:"script,required"`
	// Satoshis is the number of satoshis of the output.
	Satoshis int `json:"satoshis,required"`
	// Height is the block height.
	Height int `json:"height,required"`
}

// GetAddressUTXOs returns the txids for an address(es) (requires addressindex to be enabled).
//     addresses: The array of base58check encoded addresses
func (aic *addressIndexClient) GetAddressUTXOs(addresses []string) (addressUTXOs, error) {
	payload := getAddressUTXOsPayload{addresses}

	var response addressUTXOs
	res, err := aic.c.do(aic.c.url, "getaddressutxos", payload)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

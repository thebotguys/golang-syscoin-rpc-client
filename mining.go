package syscoinrpc

import (
	"encoding/json"
	"errors"
)

// MiningClient wraps all `mining` related functions.
type MiningClient struct {
	c *Client // The binded client, must not be nil.
}

func (mc *MiningClient) do(method string, params ...interface{}) (json.RawMessage, error) {
	return mc.c.do(method, params...)
}

// MergeMineInfo represents information required to merge-mine a block.
type MergeMineInfo struct {
	// Hash is the hash of the created block.
	BlockHash string `json:"hash,required"`
	// ChainID is the chain ID for this block.
	ChainID uint64 `json:"chainid,required"`
	// PreviousBlockHash is the hash of the previous block.
	PreviousBlockHash string `json:"previousblockhash,required"`
	// CoinbaseValue is the value of the block's coinbase.
	CoinbaseValue float64 `json:"coinbasevalue,required"`
	// Bits is the compressed target of the block.
	Bits string `json:"bits,required"`
	// Height is the height of the block.
	Height uint64 `json:"height,required"`
	// Target is the target in reversed byte order, deprecated.
	Target string `json:"_target,required"`
}

// CreateAuxBlock creates a new block and return information required to merge-mine it.
//
//     coinbase : The coinbase transaction payout address.
func (mc *MiningClient) CreateAuxBlock(coinbase string) (*MergeMineInfo, error) {
	response, err := mc.do("createauxblock", coinbase)
	if err != nil {
		return nil, err
	}

	var info MergeMineInfo
	err = json.Unmarshal(response, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetAuxBlock creates or submits a merge-mined block.
//
// Without arguments, create a new block and return information
// required to merge-mine it.  With arguments, submit a solved
// auxpow for a previously returned block.
//
//     hash   : The hash of the block to submit.
//     auxpow : The serialised auxpow found.
func (mc *MiningClient) GetAuxBlock(hash string, serializedAuxPow string) (*MergeMineInfo, error) {
	params := make([]interface{}, 0, 2)

	if hash != "" && serializedAuxPow != "" {
		params = append(params, hash, serializedAuxPow)
	} else if !(hash == "" && serializedAuxPow == "") {
		return nil, errors.New("Cannot pass hash or serializedAuxPow alone, must pass both or none")
	}

	response, err := mc.do("getauxblock", params...)
	if err != nil {
		return nil, err
	}

	var info MergeMineInfo
	err = json.Unmarshal(response, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// getBlockTemplatePayload represents the payload of a `getblocktemplate` request.
type getBlockTemplatePayload struct {
	// Mode must be set to "template", "proposal" (see BIP 23), or omitted.
	Mode string `json:"mode,omitempty"`
	// Capabilities is the list of client side supported features: 'longpoll',
	// 'coinbasetxn', 'coinbasevalue', 'proposal', 'serverlist', 'workid'.
	Capabilities []string `json:"capabilities,omitempty"`
	// Rules is the list of client side supported softfork deployments.
	Rules []string `json:"rules,omitempty"`
}

// TransactionTemplate represents a transaction template for miners and masternodes.
type TransactionTemplate struct {
	// Data is the transaction data encoded in hexadecimal (byte-per-byte)
	Data string `json:"data,required"`
	// TxID is the transaction id encoded in little-endian hexadecimal.
	TxID string `json:"txid,required"`
	// Hash is the hash encoded in little-endian hexadecimal (including witness data).
	Hash string `json:"hash,required"`
	// Depends is the array of transactions before this one (by 1-based index in
	// 'transactions' list) that must be present in the final block if this one is.
	Depends []uint64 `json:"depends,required"`
	// Fee is the difference in value between transaction inputs and outputs (in
	// satoshis); for coinbase transactions, this is a negative Number of the total
	// collected block fees (ie, not including the block subsidy); if key is not
	// present, fee is unknown and clients MUST NOT assume there isn't one.
	Fee uint64 `json:"fee,required"`
	// SigOps is the total SigOps cost, as counted for purposes of block limits;
	// if key is not present, sigop cost is unknown and clients MUST NOT assume
	// it is zero.
	SigOps uint64 `json:"sigOps,required"`
	// Weight is the total transaction weight, as counted for purposes of block
	// limits.
	Weight uint64 `json:"weight,required"`
}

// MasternodeTemplate represents a masternode template for miners and masternodes.
type MasternodeTemplate struct {
	// PayeeAddress is the payee address.
	PayeeAddress string
	// ScriptPubKey is the payee scriptPubKey.
	ScriptPubKey string
	// Amount is the required amount to pay.
	Amount uint64
}

// BlockTemplate represents a block template for miners and masternodes.
type BlockTemplate struct {
	// Version is the preferred block version.
	Version uint64 `json:"version,required"`
	// Rules is the set of specific block rules that are to be
	// enforced.
	Rules []string `json:"rules,required"`
	// VBAvailable is the set of pending, supported versionbit
	// (BIP 9) softfork deployments.
	//
	// Every single item identifies the bit number as
	// indicating acceptance and readiness for the named softfork rule.
	VBAvailable map[string]uint64 `json:"vbavailable,required"`
	// VBRequired is the bit mask of versionbits the server
	// requires set in submissions.
	VBRequired byte `json:"vbrequired,required"`
	// PreviousBlockHash is the hash of current highest block
	PreviousBlockHash string `json:"previousblockhash,required"`
	// Transactions is the array of contents of non-coinbase
	// transactions that should be included in the next block.
	Transactions []TransactionTemplate `json:"transactions,required"`
	// CoinbaseAux is the map of the data that should be included in
	// the coinbase's scriptSig content.
	//
	// Key name is to be ignored, and value included in scriptSig.
	CoinbaseAux map[string]string `json:"coinbaseaux,required"`
	// CoinbaseValue is the maximum allowable input to coinbase transaction,
	// including the generation award and transaction fees (in satoshis).
	CoinbaseValue uint64 `json:"coinbasevalue,required"`
	// CoinbaseTxn is the information for coinbase transaction.
	// TODO: find the correct struct for this.
	CoinbaseTxn json.RawMessage `json:"coinbasetxn,required"`
	// Target is the hash target
	Target string `json:"target,required"`
	// MinTime is the minimum timestamp appropriate for next
	// block time in seconds since epoch (Jan 1 1970 GMT).
	MinTime uint64 `json:"mintime,required"`
	// Mutable is the list of ways the block template may be changed.
	//
	// Every single item is a way the block template may be changed,
	// e.g. 'time', 'transactions', 'prevblock'.
	Mutable []string `json:"mutable,required"`
	// NonceRange is a range of valid nonces.
	NonceRange string `json:"noncerange,required"`
	// SigOpLimit is the limit of sigops in blocks.
	SigOpLimit uint64 `json:"sigoplimit,required"`
	// SizeLimit is the limit of block size.
	SizeLimit uint64 `json:"sizelimit,required"`
	// WeightLimit is the limit of block weight.
	WeightLimit uint64 `json:"weightlimit,required"`
	// CurrentTime is the current timestamp in seconds since epoch
	// (Jan 1 1970 GMT).
	CurrentTime uint64 `json:"curtime,required"`
	// Bits is the compressed target of next block.
	Bits string `json:"bits,required"`
	// Height is the height of the next block.
	Height uint64 `json:"uint64,required"`
	// Masternode is the required masternode payee that must
	// be included in the next block.
	Masternode MasternodeTemplate `json:"masternode,required"`
	// MasternodePaymentsEnforced is true, if masternode
	// payments are enforced.
	MasternodePaymentsEnforced bool `json:"masternode_payments_enforced,required"`
	// SuperBlock is the array of required superblock
	// payees that must be included in the next block.
	SuperBlock []MasternodeTemplate `json:"superblock,required"`
	// SuperBlocksStarted is true, if superblock payments started.
	SuperBlocksStarted bool `json:"superblocks_started,required"`
	// SuperBlockEnabled is true, if superblock payments are enabled.
	SuperBlocksEnabled bool `json:"superblocks_enabled,required"`
}

// GetBlockTemplate returns data needed to construct a block to work on.
//
// If the request parameters include a 'mode' key,
// that is used to explicitly select between the default 'template'
// request or a 'proposal'.
//
//     NOTE: all parameters are optional.
func (mc *MiningClient) GetBlockTemplate(Mode string, Capabilities []string, Rules []string) (*BlockTemplate, error) {
	payload := getBlockTemplatePayload{
		Mode:         Mode,
		Capabilities: Capabilities,
		Rules:        Rules,
	}

	response, err := mc.do("getblocktemplate", payload)
	if err != nil {
		return nil, err
	}

	var template BlockTemplate
	err = json.Unmarshal(response, &template)
	if err != nil {
		return nil, err
	}

	return &template, nil
}

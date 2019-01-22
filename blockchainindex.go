package syscoinrpc

import (
	"encoding/json"
	"errors"
	"strconv"
)

// blockchainIndexClient wraps all addressindex related functions.
type blockchainIndexClient struct {
	c *Client // The binded client, must not be nil.
}

func (bic *blockchainIndexClient) do(method string, params ...interface{}) (json.RawMessage, error) {
	return bic.c.do(method, params...)
}

// GetBestBlockHash returns the hash of the best (tip) block in
// the longest blockchain.
func (bic *blockchainIndexClient) GetBestBlockHash() (string, error) {
	res, err := bic.do("getbestblockhash")
	if err != nil {
		return "", err
	}
	var hash string
	err = json.Unmarshal(res, &hash)
	if err != nil {
		return "", err
	}

	return hash, nil
}

// AuxPow is the Auxiliary Proof of work data binded to the block.
// It contains data like coinbase block reward transaction.
// TODO: Document it better.
type AuxPow struct {
	// Tx is the block reward transaction.
	Tx AuxPowTx `json:"tx,required"`
	// Index is the Transaction index in the block.
	Index uint64 `json:"index,required"`
	// ChainIndex is the index of the transaction in the chain.
	ChainIndex uint64 `json:"chainindex,required"`
	// TODO: Discover and document this type.
	MerkleBranch []string `json:"merklebranch,required"`
	// TODO: Discover and document this type.
	ChainMerkleBranch []string `json:"chainmerklebranch,required"`
	// TODO: Document this variable.
	ParentBlock string `json:"parentblock,required"`
}

// AuxPowTx represents a block reward transaction.
type AuxPowTx struct {
	// Hex is the Hex representation of the Tx.
	Hex string `json:"hex,required"`
	// TxID is the transaction ID.
	TxID string `json:"txid,required"`
	// Size is the transaction size.
	Size uint64 `json:"size,required"`
	// Version is the transaction version.
	Version uint64 `json:"version,required"`
	// LockTime is the time (expressed as UNIX Timestamp)
	// when the transaction has been locked.
	LockTime uint64 `json:"locktime,required"`
	// Vin is the array of transaction vin objects.
	Vin []VinObject `json:"vin,required"`
	// Vout is the array of transaction vout objects.
	Vout []VoutObject `json:"vout,required"`
	// BlockHash is the block hash.
	BlockHash string `json:"blockhash,required"`
}

// VinObject represents a vin (value input) object of a transaction.
type VinObject struct {
	// Coinbase is the coinbase.
	Coinbase string `json:"coinbase,required"`
	// Sequence is the sequence number of the vin object.
	Sequence uint64 `json:"sequence,required"`
}

// VoutObject represents a vout (value output) object of a transaction.
//     TODO: Document it better.
type VoutObject struct {
}

// FullBlock represents full data of a block.
//
// It is a result from `getblock` verbose call.
type FullBlock struct {
	*FullBlockHeader
	// Size is the block size.
	Size uint64 `json:"size,required"`
	// Tx is the array of transaction IDs.
	Tx []string `json:"tx,required"`
	// AuxPow is the Auxiliary Proof of work data binded to the block.
	// It contains data like coinbase block reward transaction.
	// TODO: Document it better.
	AuxPow AuxPow `json:"auxpow,required"`
}

// GetBlock returns a string that is serialized, hex-encoded data for block 'hash'.
func (bic *blockchainIndexClient) GetBlock(blockHash string) (string, error) {
	response, err := bic.do("getblock", blockHash, false)
	if err != nil {
		return "", err
	}

	return string(response), nil
}

// GetFullBlock returns an Object with information about block <hash>.
func (bic *blockchainIndexClient) GetFullBlock(blockHash string) (*FullBlock, error) {
	response, err := bic.do("getblock", blockHash, true)
	if err != nil {
		return nil, err
	}

	var block FullBlock
	err = json.Unmarshal(response, &block)
	if err != nil {
		return nil, err
	}

	return &block, nil
}

// BlockchainInfo represents the response of a `getblockchaininfo` call.
type BlockchainInfo struct {
	// Chain is the chain name.
	Chain string `json:"chain,required"`
	// Blocks is the current number of blocks processed in the server.
	Blocks uint64 `json:"blocks,required"`
	// Headers is the current number of headers that have been validated by the node.
	Headers uint64 `json:"headers,required"`
	// BestBlockHash is the hash of the currently best block.
	BestBlockHash string `json:"bestblockhash,required"`
	// CurrentDifficulty is the current difficulty
	CurrentDifficulty float64 `json:"difficulty,required"`
	// MedianTime is the median time for the current best block.
	MedianTime uint64 `json:"mediantime,required"`
	// VerificationProgress is the estimate of verification progress completion (0..1).
	VerificationProgress float32 `json:"verificationprogress,required"`
	// ChainWork is the total amount of work in active chain, in hexadecimal.
	ChainWork string `json:"chainwork,required"`
	// Pruned is true if the blocks are subject to pruning.
	Pruned bool `json:"pruned,required"`
	//PruneHeight is the lowest-height complete block stored.
	PruneHeight uint64 `json:"pruneheight,required"`
	// Softforks is the status of softforks in progress
	Softforks []Softfork `json:"softforks,required"`
	//BIP9Softforks is the status of BIP9 softforks in progress.
	BIP9Softforks map[string]BIP9Softfork `json:"bip9_softforks,required"`
}

// Softfork represents a completed soft fork.
type Softfork struct {
	// ID is the name of the soft fork.
	ID string `json:"id,required"`
	// Version is the new block version.
	Version uint64 `json:"version,required"`
	// Reject is the progress toward rejecting pre-softfork blocks.
	Reject SoftforkRejectProgress `json:"reject,required"`
}

// SoftforkRejectProgress represents the progress toward rejecting pre-softfork blocks.
type SoftforkRejectProgress struct {
	// Status is true if threshold reached.
	Status bool `json:"status,required"`
}

// BIP9Softfork is the status of a BIP9 softfork in progress
type BIP9Softfork struct {
	// Status is the status of the pending soft fork.
	//
	// value is one of "defined", "started", "locked_in",
	// "active", "failed".
	Status string `json:"status,required"`
	// Bit is the bit (0-28) in the block version field
	// used to signal this softfork (only for "started" status).
	Bit uint8 `json:"bit,required"`
	// StartTime is the minimum median time past of a block at
	// which the bit gains its meaning.
	StartTime uint64 `json:"startTime,required"`
	// Timeout is the median time past of a block at which the
	// deployment is considered failed if not yet locked in.
	Timeout uint64 `json:"timeout,required"`
	// Since is the height of the first block to which the status
	// applies.
	Since uint64 `json:"since,required"`
}

func (bic *blockchainIndexClient) GetBlockchainInfo() (*BlockchainInfo, error) {
	response, err := bic.do("getblockchaininfo")
	if err != nil {
		return nil, err
	}

	var info BlockchainInfo
	err = json.Unmarshal(response, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetBlockCount returns the number of blocks in the longest blockchain.
//     Returns 0 on error, with the error.
func (bic *blockchainIndexClient) GetBlockCount() (uint64, error) {
	response, err := bic.do("getblockcount")
	if err != nil {
		return 0, err
	}

	blockCount, err := strconv.ParseUint(string(response), 10, 64)
	if err != nil {
		return 0, err
	}

	return blockCount, nil
}

// GetBlockHash returns the hash of the block at the given height.
func (bic *blockchainIndexClient) GetBlockHash(height uint64) (string, error) {
	response, err := bic.do("getblockhash", height)
	if err != nil {
		return "", err
	}

	return string(response), nil
}

// GetBlockHashes returns array of hashes of blocks within the timestamp range provided.
//     high: the newer block timestamp.
//     low : the older block timestamp.
func (bic *blockchainIndexClient) GetBlockHashes(high uint64, low uint64) ([]string, error) {
	response, err := bic.do("getblockhashes", high, low)
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

// FullBlockHeader represents a full block header,
// response of a verbose `getblockheader` or
// `getblockheaders` call.
type FullBlockHeader struct {
	// Hash is the block hash (same as provided)
	Hash string
	// Confirmations is the number of confirmations, or -1 if the block is not on the chain.
	Confirmations int `json:"confirmations,required"`
	// Height is the block height or index.
	Height uint64 `json:"height,required"`
	// Version is the block version.
	Version uint64 `json:"version,required"`
	// VersionHex is the block version formatted in hexadecimal.
	VersionHex string `json:"versionHex,required"`
	// MerkleRoot is the merkle root.
	MerkleRoot string `json:"merkleroot,required"`
	// Time is the block time in seconds since epoch (Jan 1 1970 GMT).
	Time uint64 `json:"time,required"`
	// MedianTime is the median block time in seconds since epoch (Jan 1 1970 GMT).
	MedianTime uint64 `json:"mediantime,required"`
	// Nonce is the nonce of the block.
	Nonce uint64 `json:"nonce,required"`
	// Bits are the bits of the block.
	// TODO: find a better explanation of this parameter.
	Bits string `json:"bits,required"`
	// Difficulty is the difficulty when the block was mined.
	Difficulty float64 `json:"difficulty,required"`
	// ChainWork is the expected number of hashes required to
	// produce the chain up to this block (in hex).
	ChainWork string `json:"chainwork,required"`
	// PreviousBlockHash is the hash of the previous block.
	PreviousBlockHash string `json:"previousblockhash,required"`
	// NextBlockHash is the hash of the next block.
	NextBlockHash string `json:"nextblockhash,required"`
}

// GetBlockHeader returns a string that is serialized, hex-encoded data for block header 'hash'.
func (bic *blockchainIndexClient) GetBlockHeader(hash string) (string, error) {
	response, err := bic.do("getblockheader", hash, false)
	if err != nil {
		return "", err
	}

	return string(response), nil
}

// GetFullBlockHeader returns an Object with information about block header <hash>.
func (bic *blockchainIndexClient) GetFullBlockHeader(hash string) (*FullBlockHeader, error) {
	response, err := bic.do("getblockheader", hash, true)
	if err != nil {
		return nil, err
	}

	var fullHeader FullBlockHeader
	err = json.Unmarshal(response, &fullHeader)
	if err != nil {
		return nil, err
	}

	return &fullHeader, nil
}

// GetBlockHeaders returns an array of items where
// each item is a string that is serialized,
// hex-encoded data for a single blockheader.
//     Hash  : The starting point hash.
//     Count : The number of headers to return (default/max = 2000).
func (bic *blockchainIndexClient) GetBlockHeaders(hash string, count uint) ([]string, error) {
	if count == 0 {
		count = 2000
	} else if count > 2000 {
		return nil, errors.New("Cannot ask more than 2000 headers")
	}

	response, err := bic.do("getblockheaders", hash, count, false)
	if err != nil {
		return nil, err
	}

	var headers []string
	err = json.Unmarshal(response, &headers)
	if err != nil {
		return nil, err
	}

	return headers, nil
}

// GetFullBlockHeaders returns an array of items with information
// about <count> blockheaders starting from <hash>.
//     Hash  : The starting point hash.
//     Count : The number of headers to return (default/max = 2000).
func (bic *blockchainIndexClient) GetFullBlockHeaders(hash string, count uint) ([]*FullBlockHeader, error) {
	if count == 0 {
		count = 2000
	} else if count > 2000 {
		return nil, errors.New("Cannot ask more than 2000 headers")
	}

	response, err := bic.do("getblockheaders", hash, count, true)
	if err != nil {
		return nil, err
	}

	var fullHeaders []*FullBlockHeader
	err = json.Unmarshal(response, &fullHeaders)
	if err != nil {
		return nil, err
	}

	return fullHeaders, nil
}

// ChainTip represents a response from the `getchaintips` call.
type ChainTip struct {
	// Height is the height of the chain tip.
	Height uint64 `json:"height,required"`
	// Hash is the block hash of the tip.
	Hash string `json:"hash,required"`
	// Difficulty is the difficulty of the mined block of the tip.
	Difficulty float64 `json:"difficulty,required"`
	// ChainWork is the expected number of hashes required to produce
	// the current chain (in hex)
	ChainWork string `json:"chainwork,required"`
	// BranchLen is the length of the branch of the chain (0 for main chain).
	BranchLen uint64 `json:"branchlen,required"`
	// ForkPoint is the fork point of the tip (same as Hash for the main chain).
	ForkPoint string `json:"forkpoint,required"`
	// Status is the status of the chain of the tip ("active" for the main chain).
	//
	//  Possible values for status:
	// "invalid"               This branch contains at least one invalid block
	// "headers-only"          Not all blocks for this branch are available, but the headers are valid
	// "valid-headers"         All blocks are available for this branch, but they were never fully validated
	// "valid-fork"            This branch is not part of the active chain, but is fully validated
	// "active"                This is the tip of the active main chain, which is certainly valid
	Status string `json:"status,required"`
}

// GetChainTips returns information about all known tips in the block tree,
// including the main chain as well as orphaned branches.
//     count     : Only show this much of latest tips
//     branchlen : Only show tips that have equal or greater length of branch
func (bic *blockchainIndexClient) GetChainTips(count uint64, branchlen uint64) ([]*ChainTip, error) {
	if count == 0 {
		count = 1
	}
	response, err := bic.do("getchaintips", count, branchlen)
	if err != nil {
		return nil, err
	}

	var tips []*ChainTip
	err = json.Unmarshal(response, &tips)
	if err != nil {
		return nil, err
	}

	return tips, nil
}

// GetDifficulty returns the current difficulty.
func (bic *blockchainIndexClient) GetDifficulty() (float64, error) {
	response, err := bic.do("getdifficulty")
	if err != nil {
		return -1, err
	}

	difficulty, err := strconv.ParseFloat(string(response), 64)
	if err != nil {
		return -1, err
	}

	return difficulty, nil
}

// MempoolEntry represents an entry of the mempool of the node.
type MempoolEntry struct {
	// Size is the transaction size in bytes.
	Size uint64 `json:"size,required"`
	// Fee is the transaction fee in SYS.
	Fee float64 `json:"fee,required"`
	// ModifiedFee is the transaction fee with fee deltas used for mining priority.
	ModifiedFee float64 `json:"modifiedfee,required"`
	// Time is the local time transaction entered pool in seconds since 1 Jan 1970 GMT.
	Time uint64 `json:"time,required"`
	// Height is the block height when the transaction entered the pool.
	Height uint64 `json:"height,required"`
	// StartingPriority is DEPRECATED. Priority when transaction entered pool.
	StartingPriority float64 `json:"startingpriority,required"`
	// CurrentPriority is DEPRECATED. Transaction priority now
	CurrentPriority float64 `json:"currentpriority,required"`
	// DescendantCount is the number of in-mempool descendant transactions (including this one).
	DescendantCount uint64 `json:"descendant_count,required"`
	// DescendantSize is the size of in-mempool descendants (including this one).
	DescendantSize uint64 `json:"descendantsize,required"`
	// DescendantFees is the modified fees (see above) of in-mempool descendants (including this one).
	DescendantFees float64 `json:"descendantfees,required"`
	// DescendantCount is the number of in-mempool descendant transactions (including this one).
	AncestorCount uint64 `json:"ancestorcount,required"`
	// AncestorSize is the size of in-mempool ancestors (including this one).
	AncestorSize uint64 `json:"ancestorsize,required"`
	// AncestorFees is the modified fees (see above) of in-mempool ancestors (including this one).
	AncestorFees float64 `json:"ancestorfees,required"`
	// DependingTransactions is the array of unconfirmed transactions used as inputs for this transaction
	DependingTransactions []string `json:"depends,required"`
	// InstantSend is true if this transaction was sent as an InstantSend one.
	InstantSend bool `json:"instantsend,required"`
	// InstantLock is true if this transaction was locked via InstantSend.
	InstantLock bool `json:"instantlock,required"`
}

// GetMempoolAncestors If txid is in the mempool, returns all in-mempool ancestors summarized data.
//     txID : The transaction id (must be in mempool)
func (bic *blockchainIndexClient) GetMempoolAncestors(txID string) ([]string, error) {
	response, err := bic.do("getmempoolancestors", txID, false)
	if err != nil {
		return nil, err
	}

	var ancestors []string
	err = json.Unmarshal(response, &ancestors)
	if err != nil {
		return nil, err
	}

	return ancestors, nil
}

// GetMempoolAncestorsFull If txid is in the mempool, returns all in-mempool ancestors full data.
//     txID : The transaction id (must be in mempool)
func (bic *blockchainIndexClient) GetMempoolAncestorsFull(txID string) ([]*MempoolEntry, error) {
	response, err := bic.do("getmempoolancestors", txID, true)
	if err != nil {
		return nil, err
	}

	var ancestors []*MempoolEntry
	err = json.Unmarshal(response, &ancestors)
	if err != nil {
		return nil, err
	}

	return ancestors, nil
}

// GetMempoolDescendants If txid is in the mempool, returns all in-mempool Descendants summarized data.
//     txID : The transaction id (must be in mempool)
func (bic *blockchainIndexClient) GetMempoolDescendants(txID string) ([]string, error) {
	response, err := bic.do("getmempooldescendants", txID, false)
	if err != nil {
		return nil, err
	}

	var descendants []string
	err = json.Unmarshal(response, &descendants)
	if err != nil {
		return nil, err
	}

	return descendants, nil
}

// GetMempoolDescendantsFull If txid is in the mempool, returns all in-mempool Descendants full data.
//     txID : The transaction id (must be in mempool)
func (bic *blockchainIndexClient) GetMempoolDescendantsFull(txID string) ([]*MempoolEntry, error) {
	response, err := bic.do("getmempooldescendants", txID, true)
	if err != nil {
		return nil, err
	}

	var descendants []*MempoolEntry
	err = json.Unmarshal(response, &descendants)
	if err != nil {
		return nil, err
	}

	return descendants, nil
}

// GetMempoolEntry returns full mempool data for given transaction.
//     txID : The transaction id (must be in mempool).
func (bic *blockchainIndexClient) GetMempoolEntry(txID string) (*MempoolEntry, error) {
	response, err := bic.do("getmempoolentry", txID)
	if err != nil {
		return nil, err
	}

	var entry MempoolEntry
	err = json.Unmarshal(response, &entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

// MempoolInfo represents general information about the mempool.
type MempoolInfo struct {
	//Size is the current tx count.
	Size uint64 `json:"size,required"`
	// Byes is the sum of all tx sizes.
	Bytes uint64 `json:"bytes,required"`
	// Usage is the total memory usage for the mempool.
	Usage uint64 `json:"usage,required"`
	// MaxMempool is the maximum memory usage for the mempool.
	MaxMempool uint64 `json:"maxmempool,required"`
	// MempoolMinFee is the minimum fee for tx to be accepted
	MempoolMinFee float64 `json:"mempoolminfee,required"`
}

func (bic *blockchainIndexClient) GetMempoolInfo() (*MempoolInfo, error) {
	response, err := bic.do("getmempoolinfo")
	if err != nil {
		return nil, err
	}

	var info MempoolInfo
	err = json.Unmarshal(response, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetRawMempool returns all transaction ids in memory pool as array of string transaction ids.
//
//     HINT: use `getmempoolentry` to fetch a specific transaction from the mempool.
func (bic *blockchainIndexClient) GetRawMempool() ([]string, error) {
	response, err := bic.do("getrawmempool", false)
	if err != nil {
		return nil, err
	}

	var rawpool []string
	err = json.Unmarshal(response, &rawpool)
	if err != nil {
		return nil, err
	}

	return rawpool, nil
}

// GetRawMempool returns all transaction ids in memory pool as array of objects.
//
//     HINT: use `getmempoolentry` to fetch a specific transaction from the mempool.
//
// Response type is a map [transactionID]MempoolEntry object.
func (bic *blockchainIndexClient) GetRawMempoolFull() (map[string]*MempoolEntry, error) {
	response, err := bic.do("getrawmempool", true)
	if err != nil {
		return nil, err
	}

	var rawpool map[string]*MempoolEntry
	err = json.Unmarshal(response, &rawpool)
	if err != nil {
		return nil, err
	}

	return rawpool, nil
}

// SpentInfo represent where a transaction is spent.
type SpentInfo struct {
	TxID  string `json:"txid,required"`
	Index uint64 `json:"index,required"`
}

// spentInfoPayload is the payload of a `getspentinfo` request.
type spentInfoPayload SpentInfo

// GetSpentInfo returns the txid and index where an output is spent.
func (bic *blockchainIndexClient) GetSpentInfo(txID string, index uint64) (*SpentInfo, error) {
	payload := spentInfoPayload{
		TxID:  txID,
		Index: index,
	}
	response, err := bic.do("getspentinfo", payload)
	if err != nil {
		return nil, err
	}

	var spentinfo SpentInfo
	err = json.Unmarshal(response, &spentinfo)
	if err != nil {
		return nil, err
	}

	return &spentinfo, nil
}

// TxOut represents a Transaction Output.
type TxOut struct {
	// BestBlock is the best block hash.
	BestBlock string `json:"bestblock,required"`
	// Confirmations is the number of confirmations.
	Confirmations uint64 `json:"confirmations,required"`
	// Value is the transaction value in SYS.
	Value float64 `json:"value,required"`
	// ScriptPubKey is the PubKey script in the output.
	ScriptPubKey ScriptPubKey `json:"scriptPubKey,required"`
	// Version is the Output version.
	Version uint64 `json:"version,required"`
	// Coinbase is true if the output comes from the coinbase.
	Coinbase bool `json:"coinbase,required"`
}

// ScriptPubKey represents a pub key script.
type ScriptPubKey struct {
	// Asm is the ASM code of the PubKey script.
	Asm string `json:"asm,required"`
	// Hex is the Hex of the PubKey script.
	Hex string `json:"code,required"`
	// RequiredSignatures is the number of required signatures.
	RequiredSignatures uint64 `json:"reqSigs,required"`
	// Type is the type of the PubKey script (e.g. pubkeyhash).
	Type string `json:"type,required"`
	// Addresses is the array of syscoin addresses involved in the PubKey script.
	Addresses []string `json:"addresses,required"`
}

func (bic *blockchainIndexClient) GetTxOut(txID string, n uint64, includeMempool bool) (*TxOut, error) {
	response, err := bic.do("gettxout", txID, n, includeMempool)
	if err != nil {
		return nil, err
	}

	var out TxOut
	err = json.Unmarshal(response, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

// GetTxOutProof returns a hex-encoded proof that "txid" was included in a block.
//
//     txIDs : An array of transaction hashes to filter.
//
//     NOTE : By default this function only works sometimes. This is when there is an
//            unspent output in the utxo for this transaction. To make it always work,
//            you need to maintain a transaction index, using the -txindex and -spentindex
//            command line option or specify the block in which the transaction is included
//            manually (by blockhash).
func (bic *blockchainIndexClient) GetTxOutProof(txIDs []string) (string, error) {
	proof, err := bic.do("gettxoutproof", txIDs)
	if err != nil {
		return "", err
	}

	return string(proof), nil
}

// GetTxOutProofInBlock returns a hex-encoded proof that "txid" was included in the block
// with the specified hash.
//
//     txIDs : An array of transaction hashes to filter.
//     blockHash : The block hash to search txIDs into.
//
//     NOTE : By default this function only works sometimes. This is when there is an
//            unspent output in the utxo for this transaction. To make it always work,
//            you need to maintain a transaction index, using the -txindex command line
//            option or specify the block in which the transaction is included manually
//            (by blockhash).
func (bic *blockchainIndexClient) GetTxOutProofInBlock(txIDs []string, blockHash string) (string, error) {
	proof, err := bic.do("gettxoutproof", txIDs, blockHash)
	if err != nil {
		return "", err
	}

	return string(proof), nil
}

// TxOutSetInfo represents statistics about the unspent transaction output set.
type TxOutSetInfo struct {
	// Height is the current block height (index).
	Height uint64 `json:"height,required,required"`
	// BestBlockHash is the hash of the best block.
	BestBlockHash string `json:"bestblock,required"`
	// TransactionCount is the number of unspent transactions.
	TransactionCount uint64 `json:"transactions,required"`
	// TxOutCount is the number of unspent transaction outputs.
	TxOutCount uint64 `json:"txouts,required"`
	// HashSerialized is the serialized hash of the unspent transaction set.
	HashSerialized string `json:"hash_serialized,required"`
	// DiskSize is the estimated size of the chainstate on disk.
	DiskSize uint64 `json:"disk_size,required"`
	// TotalAmount is the total unspent amount, in SYS.
	TotalAmount float64 `json:"total_amount,required"`
}

// GetTxOutSetInfo returns statistics about the unspent transaction output set.
//     NOTE : This call may take some time.
func (bic *blockchainIndexClient) GetTxOutSetInfo() (*TxOutSetInfo, error) {
	response, err := bic.do("gettxoutsetinfo")
	if err != nil {
		return nil, err
	}

	var info TxOutSetInfo
	err = json.Unmarshal(response, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// PreciousBlock treats a block as if it were received before others with the same work.
//
// A later preciousblock call can override the effect of an earlier one.
//
// The effects of preciousblock are not retained across restarts.
func (bic *blockchainIndexClient) PreciousBlock(blockHash string) error {
	_, err := bic.do("preciousblock", blockHash)
	return err
}

// PruneBlockchains prunes the blockchain up to the specified block height.
//
// Returns the height of the last block pruned.
//
//     heightOrTimestamp: May be set to a discrete height, or a unix timestamp
//                        to prune blocks whose block time is at least 2 hours
//                        older than the provided timestamp.
//
func (bic *blockchainIndexClient) PruneBlockchain(heightOrTimestamp uint64) (uint64, error) {
	response, err := bic.do("pruneblockchain", heightOrTimestamp)
	if err != nil {
		return 0, err
	}

	lastBlockPruned, err := strconv.ParseUint(string(response), 10, 64)
	if err != nil {
		return 0, err
	}

	return lastBlockPruned, nil
}

// VerifyChain verifies blockchain database, based on two parameters.
//
//     checkLevel : optional, 0-4, default=4 - How thorough the block verification is.
//     nBlocks    : optional, default=6, 0=all - The number of blocks to check.
func (bic *blockchainIndexClient) VerifyChain(checkLevel uint64, nBlocks uint64) (bool, error) {
	response, err := bic.do("verifychain", checkLevel, nBlocks)
	if err != nil {
		return false, err
	}

	val, err := strconv.ParseBool(string(response))
	if err != nil {
		return false, err
	}

	return val, nil
}

// VerifyTxOutProof verifies that a proof points to a transaction in a block,
// returning the transaction it commits to and throwing an RPC error if the
// block is not in our best chain.
//
//     proof : The hex-encoded proof generated by `gettxoutproof`.
func (bic *blockchainIndexClient) VerifyTxOutProof(proof string) ([]string, error) {
	response, err := bic.do("verifytxoutproof", proof)
	if err != nil {
		return nil, err
	}

	var proofTxIDs []string
	err = json.Unmarshal(response, &proofTxIDs)
	if err != nil {
		return nil, err
	}

	return proofTxIDs, nil
}
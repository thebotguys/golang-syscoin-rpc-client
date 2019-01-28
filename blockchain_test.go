package syscoinrpc_test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	syscoinrpc "github.com/thebotguys/golang-syscoin-rpc-client"
)

const invalidURL = "http://invalid.url"

func init() {
	http.DefaultClient.Timeout = time.Second * 1 // for quick tests, can be changed for more reliability.
}

var (
	testBlockHeader = syscoinrpc.FullBlockHeader{
		Hash: "9f362bce7390fb38dfa0f98c11fb9a5158aeb280f29c8f6cb5ef43d916173bf1",
		//Confirmations: ignored in this test
		Height:            1,
		Version:           805306624,
		VersionHex:        "30000100",
		MerkleRoot:        "ebc03853a2a7d1de194374a5729910e0df02b826ced4bf9d37fd4beb7df92f26",
		Time:              1525175468,
		MedianTime:        1525175468,
		Nonce:             0,
		Bits:              "207fffff",
		Difficulty:        4.656542373906925e-010,
		ChainWork:         "0000000000000000000000000000000000000000000000000000000000100012",
		PreviousBlockHash: "000006e5c08d6d2414435b294210266753b05a75f90e926dd5e6082306812622",
		NextBlockHash:     "742d1aa459648259a5464df30654c2d4203d4a8c77f895cc31188745a2c41cc7",
	}

	testBlock = syscoinrpc.FullBlock{
		FullBlockHeader: &testBlockHeader,
		Size:            393,
		Tx:              []string{"ebc03853a2a7d1de194374a5729910e0df02b826ced4bf9d37fd4beb7df92f26"},
		AuxPow: syscoinrpc.AuxPow{
			Tx: syscoinrpc.AuxPowTx{
				Hex:      "02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff29289f362bce7390fb38dfa0f98c11fb9a5158aeb280f29c8f6cb5ef43d916173bf10100000000000000ffffffff0000000000",
				TxID:     "d3d562dd548c71d2db1b7e6392bd958989b174181ff51f5d6e70b487f394d463",
				Size:     92,
				Version:  2,
				LockTime: 0,
				Vin: []syscoinrpc.VinObject{
					{
						Coinbase: "289f362bce7390fb38dfa0f98c11fb9a5158aeb280f29c8f6cb5ef43d916173bf10100000000000000",
						Sequence: 4294967295,
					},
				},
				Vout:      []syscoinrpc.VoutObject{},
				BlockHash: "bae49789e089f764a52fde5064c3257f3f07ed340dc6a7ed0748a62c29cd42d5",
			},
			MerkleBranch:      []string{},
			ChainMerkleBranch: []string{},
			ParentBlock:       "01000000000000000000000000000000000000000000000000000000000000000000000063d494f387b4706e5d1ff51f1874b1898995bd92637e1bdbd2718c54dd62d5d3000000000000000000000000",
		},
	}
)

func TestGetBestBlockHashInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetBestBlockHash()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetBlockInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	testBlockHash := "9f362bce7390fb38dfa0f98c11fb9a5158aeb280f29c8f6cb5ef43d916173bf1"

	_, err = cl.Blockchain.GetBlock(testBlockHash)
	require.Error(t, err, "Must error on any method with invalid URL")

	_, err = cl.Blockchain.GetFullBlock(testBlockHash)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetBlockchainInfoInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetBlockchainInfo()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetBlockCountInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetBlockCount()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetBlockHashInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	testHeight := uint64(1)

	_, err = cl.Blockchain.GetBlockHash(testHeight)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetBlockHeaderInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	testBlockHash := "9f362bce7390fb38dfa0f98c11fb9a5158aeb280f29c8f6cb5ef43d916173bf1"

	_, err = cl.Blockchain.GetBlockHeader(testBlockHash)
	require.Error(t, err, "Must error on any method with invalid URL")

	_, err = cl.Blockchain.GetFullBlockHeader(testBlockHash)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetAllBlockStatsInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetAllBlockStats("")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetChainTipsInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetChainTips()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetChainTxStatsInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetChainTxStats(0, "")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetDifficultyInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetDifficulty()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetMempoolAncestorsInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetMempoolAncestors("")
	require.Error(t, err, "Must error on any method with invalid URL")

	_, err = cl.Blockchain.GetMempoolAncestorsFull("")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetMempoolDescendantsInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetMempoolDescendants("")
	require.Error(t, err, "Must error on any method with invalid URL")

	_, err = cl.Blockchain.GetMempoolDescendantsFull("")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetMempoolEntryInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetMempoolEntry("")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetMempoolInfoInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetMempoolInfo()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetRawMempoolInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetRawMempool()
	require.Error(t, err, "Must error on any method with invalid URL")

	_, err = cl.Blockchain.GetRawMempoolFull()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetTxOutInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetTxOut("", 0, false)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetTxOutProofInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetTxOutProof([]string{})
	require.Error(t, err, "Must error on any method with invalid URL")

	_, err = cl.Blockchain.GetTxOutProofInBlock([]string{}, "")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetTxOutSetInfoInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetTxOutSetInfo()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestPreciousBlockInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	err = cl.Blockchain.PreciousBlock("")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestPruneBlockchainInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.PruneBlockchain(uint64(1))
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestSaveMempoolInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	err = cl.Blockchain.SaveMempool()
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestVerifyChainInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.VerifyChain(0, 0)
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestVerifyTxOutProofInvalid(t *testing.T) {
	cl, err := syscoinrpc.NewClient(invalidURL, "", "")
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.VerifyTxOutProof("")
	require.Error(t, err, "Must error on any method with invalid URL")
}

func TestGetBestBlockHashOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	hash, err := cl.Blockchain.GetBestBlockHash()
	require.NoError(t, err, "Must not error on valid URL, check if the node is running")

	t.Log("Best Block Hash:", hash)
}

func TestGetBlockOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	testBlockHash := "9f362bce7390fb38dfa0f98c11fb9a5158aeb280f29c8f6cb5ef43d916173bf1"

	block, err := cl.Blockchain.GetBlock(testBlockHash)
	require.NoError(t, err, "GetBlock: Must not error on valid URL, check if the node is running")

	t.Log("Test Block Hash:", block)

	expectedBlock := testBlock

	fullBlock, err := cl.Blockchain.GetFullBlock(testBlockHash)
	require.NoError(t, err, "GetFullBlock: Must not error on valid URL, check if the node is running")

	fullBlock.Confirmations = 0 // Removing confirmation because they change when we mine.
	require.Equal(t, expectedBlock, *fullBlock, "Must be equal to test block")
}

func TestGetBlockchainInfoOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	info, err := cl.Blockchain.GetBlockchainInfo()
	require.NoError(t, err, "Must not error on valid URL, check if the node is running")

	infoJSON, _ := json.Marshal(info)
	t.Log(string(infoJSON))
}

func TestGetBlockCountOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	count, err := cl.Blockchain.GetBlockCount()
	require.NoError(t, err, "Must not error on valid URL, check if the node is running")

	t.Log("Block Count:", count)
}

func TestGetBlockHashOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	testHeight := uint64(1)

	hash, err := cl.Blockchain.GetBlockHash(testHeight)
	require.NoError(t, err, "Must not error on valid URL, check if the node is running")

	t.Log("Block Hash:", hash)
}

func TestGetBlockHeaderOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	testBlockHash := "9f362bce7390fb38dfa0f98c11fb9a5158aeb280f29c8f6cb5ef43d916173bf1"

	block, err := cl.Blockchain.GetBlockHeader(testBlockHash)
	require.NoError(t, err, "GetBlockHeader: Must not error on valid URL, check if the node is running")

	t.Log("Test Block Header:", block)

	expectedBlockHeader := testBlockHeader

	fullBlockHeader, err := cl.Blockchain.GetFullBlockHeader(testBlockHash)
	require.NoError(t, err, "GetFullBlockHeader: Must not error on valid URL, check if the node is running")

	fullBlockHeader.Confirmations = 0
	require.Equal(t, expectedBlockHeader, *fullBlockHeader, "Must be equal to test block header")
}

func TestGetAllBlockStatsOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	hash := "dead1f156254f1d723b043a0b7ff2cbc4735c87b1852b5d829f85e8f8796f773"

	stats, err := cl.Blockchain.GetAllBlockStats(hash)
	require.NoError(t, err, "GetAllBlockStats: must not error")

	t.Log("GetAllBlockStats:", stats)
}

func TestGetChainTipsOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	tips, err := cl.Blockchain.GetChainTips()
	require.NoError(t, err, "GetChainTips: Must not error on valid URL, check if the node is running")

	t.Log("Test Chain Tips:", tips)
}

func TestGetChainTxStatsOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	stats, err := cl.Blockchain.GetChainTxStats(0, "")
	require.NoError(t, err, "GetChainTxStats: Must not error on valid URL, check if the node is running")

	t.Log("ChainTxStats:", stats)
}

func TestGetDifficultyOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	difficulty, err := cl.Blockchain.GetDifficulty()
	require.NoError(t, err, "GetBlockHeader: Must not error on valid URL, check if the node is running")

	t.Log("Test Difficulty:", difficulty)
}

func TestGetMempoolAncestorsOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	txID := "d3d562dd548c71d2db1b7e6392bd958989b174181ff51f5d6e70b487f394d463"

	_, err = cl.Blockchain.GetMempoolAncestors(txID)
	require.Error(t, err, "GetMempoolAncestors : must error with \"Transaction not in mempool\"")

	_, err = cl.Blockchain.GetMempoolAncestorsFull(txID)
	require.Error(t, err, "GetMempoolAncestorsFull : must error with \"Transaction not in mempool\"")
}

func TestGetMempoolDescendantsOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	txID := "d3d562dd548c71d2db1b7e6392bd958989b174181ff51f5d6e70b487f394d463"

	_, err = cl.Blockchain.GetMempoolDescendants(txID)
	require.Error(t, err, "GetMempoolDescendants : must error with \"Transaction not in mempool\"")

	_, err = cl.Blockchain.GetMempoolDescendantsFull(txID)
	require.Error(t, err, "GetMempoolDescendantsFull : must error with \"Transaction not in mempool\"")
}

func TestGetMempoolEntryOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	txID := "d3d562dd548c71d2db1b7e6392bd958989b174181ff51f5d6e70b487f394d463"

	_, err = cl.Blockchain.GetMempoolEntry(txID)
	require.Error(t, err, "GetMempoolEntry : must error with \"Transaction not in mempool\"")
}

func TestGetMempoolInfoOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	_, err = cl.Blockchain.GetMempoolInfo()
	require.NoError(t, err, "GetMempoolEntry : must not error")
}

func TestGetRawMempoolOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	rawPool, err := cl.Blockchain.GetRawMempool()
	require.NoError(t, err, "GetRawMempool : must not error")

	t.Log("GetRawMempool :", rawPool)

	rawPoolFull, err := cl.Blockchain.GetRawMempoolFull()
	require.NoError(t, err, "GetRawMempoolFull : must not error")

	t.Log("GetRawMempoolFull :", rawPoolFull)
}

func TestGetTxOutOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	txID := "0437cd7f8525ceed2324359c2d0ba26006d92d856a9c20fa0241106ee5a597c9"
	n := uint64(0)
	includeMempool := true

	out, err := cl.Blockchain.GetTxOut(txID, n, includeMempool)
	require.NoError(t, err, "GetTxOut : must not error")

	t.Log("GetTxOut :", out)
}

func TestGetTxOutProofOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	txIDs := []string{"0f6f667c7b31395a5c0e25b5b0f10cb94a1fa1d4be0f0c9b83e5d99e30ffeed3"}
	blockHash := "9299180cfe417ee3c6bcc96bd706de154a7cc92eaff662b56178c2b5ff0fdfe2"

	proofs, err := cl.Blockchain.GetTxOutProof(txIDs)
	require.NoError(t, err, "GetTxOutProof: must not error")

	t.Log("GetTxOutProof :", proofs)

	proofs, err = cl.Blockchain.GetTxOutProofInBlock(txIDs, blockHash)
	require.NoError(t, err, "GetTxOutProofInBlock: must not error")

	t.Log("GetTxOutProofInBlock :", proofs)
}

func TestGetTxOutSetInfoOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	info, err := cl.Blockchain.GetTxOutSetInfo()
	require.NoError(t, err, "GetTxOutSetInfo: must not error")

	t.Log("GetTxOutSetInfo :", info)
}

func TestPreciousBlockOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	t.Skip("This call would alter the node, so for this tests is skipped, remove the skip instruction to do it anyway")

	testBlockHash := "9f362bce7390fb38dfa0f98c11fb9a5158aeb280f29c8f6cb5ef43d916173bf1"

	err = cl.Blockchain.PreciousBlock(testBlockHash)
	require.NoError(t, err, "PreciousBlock: must not error")
}

func TestPruneBlockchainOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	t.Skip("This call would alter the node, so for this tests is skipped, remove the skip instruction to do it anyway")

	testHeight := uint64(10000)
	testTimestamp := uint64(time.Now().Add(-3 * time.Hour).Unix())

	_, err = cl.Blockchain.PruneBlockchain(testHeight)
	require.NoError(t, err, "PruneBlockchain: must not error")

	_, err = cl.Blockchain.PruneBlockchain(testTimestamp)
	require.NoError(t, err, "PruneBlockchain: must not error")
}

func TestSaveMempoolOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	t.Skip("This call would alter the node, so for this tests is skipped, remove the skip instruction to do it anyway")

	err = cl.Blockchain.SaveMempool()
	require.NoError(t, err, "SaveMempool: must not error")
}

func TestVerifyChainOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	verified, err := cl.Blockchain.VerifyChain(4, 6)
	require.NoError(t, err, "VerifyChain: must not error")

	t.Log("VerifyChain :", verified)
}

func TestVerifyTxOutProofOK(t *testing.T) {
	cl, err := syscoinrpc.NewClient(syscoinrpc.LocalNodeURL, os.Getenv("RPC_USER"), os.Getenv("RPC_PASSWORD"))
	require.NoError(t, err, "Must have no error on creation, even with invalid URL")

	testProof := "00010030592936b277849fde57969d468b07b70cb4e90a3642dd7de22558d6188fbbf911d3eeff309ed9e5839b0c0fbed4a11f4ab90cf1b0b5250e5c5a39317b7c666f0f02ef4d5cd15f06180000000001000000010000000000000000000000000000000000000000000000000000000000000000ffffffff6003ed8c081d2f5669614254432f4d696e656420627920666163696c6974793936332f2cfabe6d6db6e3813c648a3ec91b98691651f4459b5500a128e953ebfeb90416784eae644b040000000000000010902734086d644fc5b7e2da3e68be0000ffffffff021e142a4b000000001976a914536ffa992491508dca0354e52f32a3a7a679a53a88ac0000000000000000266a24aa21a9eda64137e63f795c26dbaba0895c2d4c39feac11d9cf0f823cc58004cd95d65e5c000000007b67455430797ac1b40e6a6af12421809a28d3c9c3a7870500000000000000000c4d9f4d83b526cb802e0600a684ef32ad46b31696b5d31320e4e85f7a3c9f99e25299c8fa44125c8746001b35a088593bbac70c593347dddb0ed1484b438c523cc7ae88bdd96bcf06cd99f387bf48722ed2379d993071d33ad8d17404b6fc7a92486be467df89169f14eb58d25867bdbda839e97c1100883f8184b77e52b00a76c934e183f37f3ea8660c9b8cdda80e0ee8249258e935412dd71783b2254da8457f4ce60c04f5365331a57cead102b7bb75342da15849305a061e45ace1ff1bdcb4c6bd86e9f22ea3163e2c515e470b8c89f3dd8b0605b58ec70bdbae5a8783aca6c7f55d2bef7dd37b60cdbb8fe7a8a438a91cb6ef9b39443f41ee68346f846372abf6f503aa77f7493b69b50b6f9b91113e59131840620b7ca8374535f66daeda0bbb54caf97046809f66d5ef5359d6a68b91bdee7c4679e80f4ea5a5cbd05a49c7fd842c774ef86d725f93f01d0cfb437c11bff1ae296dec4708eb6b654121776b065871c8a0a8f09877b9b91493d190a14589144544a800ecf16628e2338400000000027659ab8a2e71540487b9c2b83f04421c18af6cfe7de65e870b7b20e1b8b2243fa7500873ee5ded9ff99deeeb7514386b22e57901588db57a9528f3ba8d1cbfad0200000000000020e194e148a5720e845039884a5ecb7a5678a5f905093a2000000000000000000045cb026e55f692d58ff7700b797736713006379f1cdc4f7c515e7c279ef7c51634ef4d5c33d62f17aef92c6d0100000001d3eeff309ed9e5839b0c0fbed4a11f4ab90cf1b0b5250e5c5a39317b7c666f0f0101"

	proofs, err := cl.Blockchain.VerifyTxOutProof(testProof)
	require.NoError(t, err, "VerifyTxOutProof : must not error")

	t.Log("VerifyTxOutProof :", proofs)
}

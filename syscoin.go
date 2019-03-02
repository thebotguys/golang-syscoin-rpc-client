package syscoinrpc

import (
	"encoding/json"
)

// SyscoinClient wraps all `syscoin` related functions.
type SyscoinClient struct {
	c *Client // The bound client, must not be nil.
}

func (sc *SyscoinClient) do(method string, params ...interface{}) (json.RawMessage, error) {
	return sc.c.do(method, params...)
}

// GovernanceInfo represents an object containing governance parameters.
type GovernanceInfo struct {
	// GovernanceMinimumQuorum is the absolute
	// minimum number of votes needed to trigger
	// a governance action.
	GovernanceMinimumQuorum int64 `json:"governanceminquorum,required"`
	// MasternodeWatchDogMaxSeconds is the sentinel
	// watchdog expiration time in seconds. (DEPRECATED)
	MasternodeWatchDogMaxSeconds int64 `json:"masternodewatchdogmaxseconds,required"`
	// SentinelPingMaxSeconds is the sentinel
	// ping expiration time in seconds.
	SentinelPingMaxSeconds int64 `json:"sentinelpingmaxseconds,required"`
	// ProposalFee is the collateral transaction
	// fee which must be paid to create a proposal
	// in SYS.
	ProposalFee float64 `json:"proposalfee,required"`
	// SuperblockCycle is the the number of blocks
	// between superblocks.
	SuperblockCycle int64 `json:"superblockcycle,required"`
	// LastSuperblock is the block number of the last superblock.
	LastSuperblock int64 `json:"lastsuperblock,required"`
	// NextSuperblock is the block number of the next superblock.
	NextSuperblock int64 `json:"nextsuperblock,required"`
	// MaxGovernanceObjectDataSize is the maximum governance
	// object data size in bytes.
	MaxGovernanceObjectDataSize int64 `json:"maxgovobjdatasize,required"`
}

// GetGovernanceInfo returns an object containing governance parameters.
func (sc *SyscoinClient) GetGovernanceInfo() (*GovernanceInfo, error) {
	response, err := sc.do("getgovernanceinfo")
	if err != nil {
		return nil, err
	}

	var info GovernanceInfo
	err = json.Unmarshal(response, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

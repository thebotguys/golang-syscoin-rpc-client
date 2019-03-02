# Golang Syscoin JSON RPC Client

[![Build Status](https://travis-ci.org/thebotguys/golang-syscoin-rpc-client.svg?branch=master)](https://travis-ci.org/thebotguys/golang-syscoin-rpc-client)
[![Godoc](https://godoc.org/thebotguys/golang-syscoin-rpc-client?status.svg)](https://godoc.org/thebotguys/golang-syscoin-rpc-client)
[![Goreportcard](https://goreportcard.com/badge/github.com/thebotguys/golang-syscoin-rpc-client)](https://goreportcard.com/report/github.com/thebotguys/golang-syscoin-rpc-client)
[![codecov](https://codecov.io/gh/thebotguys/golang-syscoin-rpc-client/branch/master/graph/badge.svg)](https://codecov.io/gh/thebotguys/golang-syscoin-rpc-client)

**IMPORTANT: THIS IS A WORK IN PROGRESS AND NOT READY FOR PRODUCTION!**

Syscoin Golang JSON-RPC Client is an indipendent project trying to implement a JSON RPC client for the Syscoin JSON-RPC API.

## Install

To install the package simply run

``` bash
go get github.com/thebotguys/golang-syscoin-rpc-client
```

## Usage

Here is an example of usage, other examples can be found in ***test_files.go***.

``` go
import (
    "github.com/thebotguys/golang-syscoin-rpc-client"
)

func main() {
    // Example loading from environment variables.
    rpcHost := os.Getenv("SYSCOIN_RPC_HOST")
    rpcPort := os.Getenv("SYSCOIN_RPC_PORT")

    rpcEndpoint := fmt.Sprintf("%s:%s", rpcHost, rpcPort)
    rpcUsername := os.Getenv("SYSCOIN_RPC_USERNAME")
    rpcPassword := os.Getenv("SYSCOIN_RPC_PASSWORD")

    // First we need to instantiate a client.
    client, err := syscoinrpc.NewClient(rpcEndpoint, rpcUser, rpcPassword)
    if err != nil {
        // Handle the error
    }

    // Then you call the RPC Methods (for example, getbestblockhash)
    // from the specific sub-client (in this case `blockchain`)
    // You can check the godoc for details about all functions.

    hash, err := client.Blockchain.GetBestBlockHash()
    if err != nil {
        // Handle the error
    }

    fmt.Println(hash)
}
```

## Currently Implemented commands

### Blockchain

- [x] `getbestblockhash`
- [x] `getblock`
- [x] `getblockchaininfo`
- [x] `getblockcount`
- [x] `getblockhash`
- [x] `getblockheader`
- [x] `getblockstats`
- [x] `getchaintips`
- [x] `getchaintxstats`
- [x] `getdifficulty`
- [x] `getmempoolancestors`
- [x] `getmempooldescendants`
- [x] `getmempoolentry`
- [x] `getmempoolinfo`
- [x] `getrawmempool`
- [x] `gettxout`
- [x] `gettxoutproof`
- [x] `gettxoutsetinfo`
- [x] `preciousblock`
- [x] `pruneblockchain`
- [x] `savemempool`
- [ ] `scantxoutset` NOT IMPLEMENTED : EXPERIMENTAL warning: this call may be removed or changed in future releases.
- [x] `verifychain`
- [x] `verifytxoutproof`

### Control

- [x] `getmemoryinfo`
- [x] `help`
- [x] `logging`
- [x] `stop`
- [x] `uptime`

### Generating

- [x] `generate`
- [x] `generatetoaddress`

### Mining

- [x] `createauxblock`
- [x] `getauxblock`
- [x] `getblocktemplate`
- [x] `getmininginfo`
- [x] `getnetworkhashps`
- [x] `prioritisetransaction`
- [x] `submitauxblock`
- [x] `submitblock`

### Network commands

- [x] `addnode`
- [x] `clearbanned`
- [x] `disconnectnode`
- [x] `getaddednodeinfo`
- [x] `getconnectioncount`
- [x] `getnettotals`
- [x] `getnetworkinfo`
- [x] `getpeerinfo`
- [x] `listbanned`
- [x] `ping`
- [x] `setban`
- [x] `setnetworkactive`

### RawTransaction commands

- [ ] `combinepsbt`
- [ ] `combinerawtransaction`
- [ ] `converttopsbt`
- [ ] `createpsbt`
- [ ] `createrawtransaction`
- [ ] `decodepsbt`
- [ ] `decoderawtransaction`
- [ ] `decodescript`
- [ ] `finalizepsbt`
- [ ] `fundrawtransaction`
- [ ] `getrawtransaction`
- [ ] `sendrawtransaction`
- [ ] `signrawtransaction`
- [ ] `signrawtransactionwithkey`
- [ ] `testmempoolaccept`

### Syscoin commands

- [ ] `addressbalance`
- [ ] `assetallocationburn`
- [ ] `assetallocationinfo`
- [ ] `assetallocationmint`
- [ ] `assetallocationsend`
- [ ] `assetallocationsenderstatus`
- [ ] `assetinfo`
- [ ] `assetnew`
- [ ] `assetsend`
- [ ] `assettransfer`
- [ ] `assetupdate`
- [x] `getgovernanceinfo`
- [ ] `getsuperblockbudget`
- [ ] `gobject`
- [ ] `listassetallocationmempoolbalances`
- [ ] `listassetallocations`
- [ ] `listassetallocationtransactions`
- [ ] `listassets`
- [ ] `mnsync`
- [ ] `spork`
- [ ] `syscoinaddscript`
- [ ] `syscoinburn`
- [ ] `syscoindecoderawtransaction`
- [ ] `syscoinlistreceivedbyaddress`
- [ ] `syscoinmint`
- [ ] `syscoinsetethheaders`
- [ ] `syscoinsetethstatus`
- [ ] `syscointxfund`
- [ ] `tpstestadd`
- [ ] `tpstestinfo`
- [ ] `tpstestsetenabled`
- [ ] `voteraw`
  
### Util commands

- [ ] `createmultisig`
- [ ] `estimatefee`
- [ ] `estimatesmartfee`
- [ ] `signmessagewithprivkey`
- [ ] `validateaddress`
- [ ] `verifymessage`

### Wallet commands

***TBD***

### ZMQ commands

- [ ] `getzmqnotifications`

# Golang Syscoin JSON RPC Client

[![Build Status](https://travis-ci.org/thebotguys/golang-syscoin-rpc-client.svg?branch=master)](https://travis-ci.org/thebotguys/golang-syscoin-rpc-client)

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

    // Then you call the RPC Methods (for example, getaddressbalance)
    // from the specific sub-client (in this case `addressindex`)
    // You can check the godoc for details about all functions.
    addressesToCheck := []string{ /* Your addresses */}
    separatedOutput := false
    addrBal, err := client.AddressIndex.GetAddressBalance(addressesToCheck, separatedOutput)
    if err != nil {
        // Handle the error
    }

    fmt.Println(addrBal)
}
```

## Additional Notes

Full Reference is available at [https://syscoin.readme.io/v3.2.0/reference](https://syscoin.readme.io/v3.2.0/reference).

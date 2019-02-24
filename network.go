package syscoinrpc

import (
	"encoding/json"
	"strconv"
)

// NetworkClient wraps all `blockchain` related functions.
type NetworkClient struct {
	c *Client // The binded client, must not be nil.
}

func (nc *NetworkClient) do(method string, params ...interface{}) (json.RawMessage, error) {
	return nc.c.do(method, params...)
}

// AddNode attempts to add or remove a node from the addnode
// list or tries a connection to a node once.
//
// Nodes added using addnode (or -connect) are protected from
// DoS disconnection and are not required to be full nodes or
// support SegWit as other outbound peers are (though such
// peers will not be synced from).
//     node    : The node (see `getpeerinfo` for nodes).
//     command : 'add' to add a node to the list,
//               'remove' to remove a node from the list,
//               'onetry' to try a connection to the node once.
func (nc *NetworkClient) AddNode(node string, command string) error {
	_, err := nc.do("addnode", node, command)
	return err
}

// ClearBannedIPs clears all banned ip list.
func (nc *NetworkClient) ClearBannedIPs() error {
	_, err := nc.do("clearbanned")
	return err
}

// DisconnectNode immediately disconnects from the specified peer node.
//
// Strictly one out of 'address' and 'nodeid' can be provided to identify
// the node.
//
// To disconnect by nodeid, set 'address' to the empty string.
//
//     address : The IP address/port of the node.
//     nodeid  : The node ID (see getpeerinfo for node IDs).
func (nc *NetworkClient) DisconnectNode(ip string, nodeid string) error {
	_, err := nc.do("disconnectnode", ip, nodeid)
	return err
}

// ConnectionAddress represents the connection
// details of a node.
type ConnectionAddress struct {
	// Address is the syscoin server IP and port
	// we're connected to.
	Address string `json:"address,required"`
	// Connection is `connection`, `inbound` or
	// `outbound`.
	Connection string `json:"connection,required"`
}

// AddedNodeInfo represents network information about
// an added node.
type AddedNodeInfo struct {
	// AddedNode is the node IP address or name
	// (as provided to addnode).
	AddedNode string `json:"addednode,required"`
	// Connected is true if the node is connected.
	Connected bool `json:"connected,required"`
	// Addresses is the connection details array
	// Valid only if connected == true.
	Addresses []ConnectionAddress `json:"addresses,omitempty"`
}

// GetAddedNodeInfo returns information about the given added node, or all added nodes if ip is empty.
//
// (note that onetry addnodes are not listed here).
func (nc *NetworkClient) GetAddedNodeInfo(ip string) ([]*AddedNodeInfo, error) {
	var params []interface{}

	if ip != "" {
		params = append(params, ip)
	}

	response, err := nc.do("getaddednodeinfo", params)
	if err != nil {
		return nil, err
	}

	var nodes []*AddedNodeInfo
	err = json.Unmarshal(response, &nodes)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

// GetConnectionCount returns the number of connections to other nodes.
func (nc *NetworkClient) GetConnectionCount() (uint64, error) {
	response, err := nc.do("getconnectioncount")
	if err != nil {
		return 0, err
	}

	count, err := strconv.ParseUint(string(response), 10, 64)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// UploadTarget represents the upload target of the network of the node.
type UploadTarget struct {
	// TimeframeSeconds is the length of the measuring timeframe in seconds.
	TimeframeSeconds uint64 `json:"timeframe,required"`
	// Target is the target in bytes.
	Target uint64 `json:"target,required"`
	// TargetReached is true if the target has been reached.
	TargetReached bool `json:"target_reached,required"`
	// ServingHistoricalBlocks is true if serving historical blocks.
	ServingHistoricalBlocks bool `json:"serve_historical_blocks,required"`
	// BytesLeftInCycle is the number of bytes left in current time cycle.
	BytesLeftInCycle uint64 `json:"bytes_left_in_cycle,required"`
	// SecondsLeftInCycle is the number of seconds left in current time cycle.
	SecondsLeftInCycle uint64 `json:"time_left_in_cycle,required"`
}

// NetworkTotals represents information about network traffic, including bytes in,
// bytes out and current time.
type NetworkTotals struct {
	// TotalBytesReceived is the number of total bytes in
	TotalBytesReceived uint64 `json:"totalbytesrecv,required"`
	// TotalBytesSent is the number of total bytes out
	TotalBytesSent uint64 `json:"totalbytessent,required"`
	// TimestampMillis is the current UNIX time, in milliseconds.
	TimestampMillis uint64 `json:"timemillis,required"`
	// UploadTarget contains upload target data of the network.
	UploadTarget UploadTarget `json:"uploadtarget,required"`
}

// GetNetworkTotals returns the number of connections to other nodes.
func (nc *NetworkClient) GetNetworkTotals() (*NetworkTotals, error) {
	response, err := nc.do("getnettotals")
	if err != nil {
		return nil, err
	}

	var totals NetworkTotals
	err = json.Unmarshal(response, &totals)
	if err != nil {
		return nil, err
	}

	return &totals, nil
}

// Network represents a single network information.
type Network struct {
	// Name is the network (ipv4, ipv6 or onion) name.
	Name string `json:"name,required"`
	// Limited is true if the network is limited (e.g. using -onlynet flag).
	Limited bool `json:"limited,required"`
	// Reachable is true if the network is reachable.
	Reachable bool `json:"reachable,required"`
	// Proxy is the proxy that is used for this network, or empty if none.
	Proxy string `json:"proxy,required"`
	// RandomizedProxyCredentials is true if randomized credentials are used.
	RandomizedProxyCredentials bool `json:"proxy_randomize_credentials,required"`
}

// LocalAddress represents a local address in a network.
type LocalAddress struct {
	Address string // Network Address.
	Port    int    // Network Port.
	Score   uint64 // Address relative score.
}

// NetworkInfo represents generic network information.
type NetworkInfo struct {
	// Version is the server version.
	Version uint64 `json:"version,required"`
	// Subversion is the server sub-version.
	Subversion string `json:"subversion,required"`
	// ProtocolVersion is the protocol version.
	ProtocolVersion uint64 `json:"protocolversion,required"`
	// LocalServices represents the services the node offers to the network.
	LocalServices string `json:"localservices,required"`
	// LocalRelay is true if transaction relay is requested from peers,
	// false otherwise.
	LocalRelay bool `json:"localrelay,required"`
	// TimeOffset represents the time offset.
	TimeOffset int64 `json:"timeoffset,required"`
	// ConnectionCount represents the number of connections.
	ConnectionCount uint64 `json:"connections,required"`
	// NetworkIsActive is true if p2p networking is enabled.
	NetworkIsActive bool `json:"networkactive,required"`
	// Networks is the array of networks the node is connected to.
	Networks []Network
	// RelayFee is the minimum relay fee for transactions in SYS/kB.
	RelayFee float64 `json:"relayfee,required"`
	// IncrementalFee is the minimum fee increment for mempool limiting
	// or BIP 125 replacement in SYS/kB.
	IncrementalFee float64 `json:"incrementalfee,required"`
	// LocalAddresses is the list of local addresses.
	LocalAddresses []LocalAddress `json:"localaddresses,required"`
	// Warnings contains any network and blockchain warning.
	Warnings string `json:"warnings,required"`
}

// GetNetworkInfo returns an object containing various state info
// regarding P2P networking.
func (nc *NetworkClient) GetNetworkInfo() (*NetworkInfo, error) {
	response, err := nc.do("getnetworkinfo")
	if err != nil {
		return nil, err
	}

	var info NetworkInfo
	err = json.Unmarshal(response, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// PeerInfo represents general info about a peer in the network.
type PeerInfo struct {
	// ID is the peer index.
	ID uint64 `json:"id,required"`
	// Address is the IP address and port of the peer.
	Address string `json:"addr,required"`
	// BindAddress is the bind address of the connection to the peer.
	BindAddress string `json:"addrbind,required"`
	// LocalAddress is the local address as reported by the peer.
	LocalAddress string `json:"addrlocal,required"`
	// Services represents the services offered by the peer.
	Services string `json:"services,required"`
	// RelaysTx is true if peer has asked us to relay transactions to it.
	RelaysTx bool `json:"relaytxes,required"`
	// LastSend is the time in seconds since epoch (Jan 1 1970 GMT)
	// of the last send.
	LastSend uint64 `json:"lastsend,required"`
	// LastReceive is the time in seconds since epoch (Jan 1 1970 GMT)
	// of the last receive.
	LastReceive uint64 `json:"lastrecv,required"`
	// BytesSend is the total number of bytes sent to the peer.
	BytesSent uint64 `json:"bytessent,required"`
	// BytesReceived is the total number of bytes received by the peer.
	BytesReceived uint64 `json:"bytesrecv,required"`
	// ConnectionTime is the connection time in seconds since epoch (Jan 1 1970 GMT)
	ConnectionTime uint64 `json:"conntime,required"`
	// TimeOffset is the time offset in seconds.
	TimeOffset uint64 `json:"timeoffset,required"`
	// PingTime is the ping time (if available).
	PingTime uint64 `json:"pingtime,required"`
	// MinimumPingTime is the minimum observed ping time (if any at all).
	MinimumPingTime uint64 `json:"minping,required"`
	// PingWaitTime is the ping wait time (if non-zero).
	PingWaitTime uint64 `json:"pingwait,required"`
	// Version is the peer version, such as 70001.
	Version uint64 `json:"version,required"`
	// Subversion is the sub-version of the peer, such as "Satoshi:0.8.5"
	Subversion string `json:"subver,required"`
	// IsInbound is true if the peer is inbound or false if outbound.
	IsInbound bool `json:"inbound,required"`
	// IsAddedManually is true if connection was due to addnode/-connect,
	// false if it was an automatic/inbound connection.
	IsAddedManually bool `json:"addnode,required"`
	// StartingHeight is the starting height (block) of the peer.
	Startingheight uint64 `json:"startingheight,required"`
	// BanScore contains the ban score of the peer.
	BanScore uint64 `json:"banscore,required"`
	// LastSyncedHeader is the last header we have in common with this peer.
	LastSyncedHeader uint64 `json:"synced_headers,required"`
	// LastSyncedBlock is the last block we have in common with this peer.
	LastSyncedBlock uint64 `json:"synced_blocks,required"`
	// Inflight is the array of heights of blocks we're currently asking
	// from this peer.
	Inflight []uint64 `json:"inflight,required"`
	// IsWhitelisted is true if the peer is whitelisted.
	IsWhitelisted bool `json:"whitelisted,required"`
	// BytesSentPerMessage is the total bytes sent aggregated by message type.
	//
	// { "address" : total }
	BytesSentPerMessage map[string]uint64 `json:"bytessent_per_msg,required"`
	// BytesReceivedPerMessage is the total bytes received aggregated by message type.
	//
	// { "address" : total }
	BytesReceivedPerMessage map[string]uint64 `json:"bytesrecv_per_msg,required"`
}

// GetPeerInfo returns an object containing various state info
// regarding connected peers.
func (nc *NetworkClient) GetPeerInfo() ([]*PeerInfo, error) {
	response, err := nc.do("getpeerinfo")
	if err != nil {
		return nil, err
	}

	var info []*PeerInfo
	err = json.Unmarshal(response, &info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

// ListBannedIPs lists all banned IPs/Subnets.
func (nc *NetworkClient) ListBannedIPs() ([]string, error) {
	response, err := nc.do("listbanned")
	if err != nil {
		return nil, err
	}

	var ips []string
	err = json.Unmarshal(response, &ips)
	if err != nil {
		return nil, err
	}

	return ips, nil
}

// Ping requests that a ping be sent to all other nodes, to measure ping time.
//
// Results provided in getpeerinfo, pingtime and pingwait fields are decimal seconds.
//
// The ping command is handled in queue with all other commands, so it measures
// processing backlog, not just network ping.
func (nc *NetworkClient) Ping() error {
	_, err := nc.do("ping")
	return err
}

// SetBan attempts to add or remove an IP/Subnet from the banned list.
//
// A banTime of 0 implies using the default banTime defined in the server
// (usually 24h).
//
// If isAbsoluteTime is true, bantime is interpreted as absolute endtime of
// the ban.
func (nc *NetworkClient) SetBan(ip string, command string, banTime uint64, isAbsoluteTime bool) error {
	_, err := nc.do("setban", ip, command, banTime, isAbsoluteTime)
	return err
}

// SetNetworkActive disables/enables all p2p network activity.
func (nc *NetworkClient) SetNetworkActive(active bool) error {
	_, err := nc.do("setnetworkactive", active)
	return err
}

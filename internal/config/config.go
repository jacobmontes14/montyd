// Package config is the base configuration for all nodes
package config

import (
	"flag"
	"strings"
	"time"
)

type Config struct {
	ID                 string
	ClientAddr         string
	PeerAddr           string
	Peers              map[string]string
	DataDir            string
	HeartbeatInterval  time.Duration
	ElectionTimeoutMin time.Duration
	ElectionTimeoutMax time.Duration
}

const (
	peerIndex   = 0
	peerAddr    = 1
	defaultTime = 50 * time.Millisecond
)

func Load() (config *Config, err error) {
	id := flag.String("nodeID", "", "the nodes unique ID")
	clientAddr := flag.String("clientID", "", "client address for callers")
	peerAddr := flag.String("peerID", "", "peer address for other nodes")
	peers := flag.String("peers", "", "list of peer nodes")
	datadir := flag.String("datadir", "", "directory for data")
	heartbeatInterval := flag.Duration("heartbeatInterval", defaultTime, "node heartbeat interval")
	electionTimeoutMin := flag.Duration("electionTimeoutMin", defaultTime, "minimum election timeout")
	electionTimeoutMax := flag.Duration("electionTimeoutMax", defaultTime, "maximum election timeout")

	flag.Parse()

	return &Config{
		ID:                 *id,
		ClientAddr:         *clientAddr,
		PeerAddr:           *peerAddr,
		Peers:              parsePeerList(peers),
		DataDir:            *datadir,
		HeartbeatInterval:  *heartbeatInterval,
		ElectionTimeoutMin: *electionTimeoutMin,
		ElectionTimeoutMax: *electionTimeoutMax,
	}, nil
}

func parsePeerList(peers *string) map[string]string {
	peersList := strings.Split(*peers, ",")
	completePeerMap := make(map[string]string)

	for _, peer := range peersList {
		peerAndAddr := strings.Split(peer, "=")
		completePeerMap[peerAndAddr[peerIndex]] = peerAndAddr[peerAddr]
	}

	return completePeerMap
}

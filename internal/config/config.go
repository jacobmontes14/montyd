// Package config is the base configuration for all nodes
package config

import (
	"errors"
	"flag"
	"fmt"
	"net"
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

// defaults from the Raft paper
const (
	defaultHeartbeatInterval  = 50 * time.Millisecond
	defaultElectionTimeoutMin = 150 * time.Millisecond
	defaultElectionTimeoutMax = 300 * time.Millisecond
)

func Load() (*Config, error) {
	id := flag.String("nodeID", "", "the nodes unique ID")
	clientAddr := flag.String("clientAddr", "", "client address for callers")
	peerAddr := flag.String("peerAddr", "", "peer address for other nodes")
	peers := flag.String("peers", "", "list of peer nodes")
	datadir := flag.String("datadir", "", "directory for data")
	heartbeatInterval := flag.Duration("heartbeatInterval", defaultHeartbeatInterval, "node heartbeat interval")
	electionTimeoutMin := flag.Duration("electionTimeoutMin", defaultElectionTimeoutMin, "minimum election timeout")
	electionTimeoutMax := flag.Duration("electionTimeoutMax", defaultElectionTimeoutMax, "maximum election timeout")

	flag.Parse()
	verifiedPeers, err := parsePeerList(*peers)
	if err != nil {
		return nil, fmt.Errorf("invalid peers: %w", err)
	}
	newConfig := &Config{
		ID:                 *id,
		ClientAddr:         *clientAddr,
		PeerAddr:           *peerAddr,
		Peers:              verifiedPeers,
		DataDir:            *datadir,
		HeartbeatInterval:  *heartbeatInterval,
		ElectionTimeoutMin: *electionTimeoutMin,
		ElectionTimeoutMax: *electionTimeoutMax,
	}

	if err := validateConfig(newConfig); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}

	return newConfig, nil
}

// parsePeerList parses a comma-separated list of id=host:port pairs.
func parsePeerList(peers string) (map[string]string, error) {
	if strings.TrimSpace(peers) == "" {
		return nil, errors.New("no peers provided")
	}
	peersList := strings.Split(peers, ",")
	completePeerMap := make(map[string]string)

	for _, peer := range peersList {
		peer = strings.TrimSpace(peer)
		if peer == "" {
			return nil, errors.New("empty entry in peer list")
		}

		id, addr, ok := strings.Cut(peer, "=")
		if !ok {
			return nil, fmt.Errorf("pair %q is missing '='", peer)
		}
		id = strings.TrimSpace(id)
		addr = strings.TrimSpace(addr)

		if id == "" {
			return nil, fmt.Errorf("pair %q has no id", peer)
		}
		if addr == "" {
			return nil, fmt.Errorf("pair %q has no address", peer)
		}
		if _, _, err := net.SplitHostPort(addr); err != nil {
			return nil, fmt.Errorf("pair %q has invalid address: %w", peer, err)
		}
		if _, exists := completePeerMap[id]; exists {
			return nil, fmt.Errorf("duplicate peer id %q", id)
		}

		completePeerMap[id] = addr
	}

	return completePeerMap, nil
}

func validateConfig(config *Config) error {
	var newErrors []error
	if config.ID == "" {
		newErrors = append(newErrors, errors.New("config id not provided"))
	}
	if config.ClientAddr == "" {
		newErrors = append(newErrors, errors.New("client address not provided"))
	}
	if config.PeerAddr == "" {
		newErrors = append(newErrors, errors.New("peer address not provided"))
	}
	if config.DataDir == "" {
		newErrors = append(newErrors, errors.New("data directory not provided"))
	}

	if config.HeartbeatInterval <= 0 {
		newErrors = append(newErrors, fmt.Errorf("heartbeat interval must be positive, got %v", config.HeartbeatInterval))
	}
	if config.ElectionTimeoutMin <= 0 {
		newErrors = append(newErrors, fmt.Errorf("election timeout min must be positive, got %v", config.ElectionTimeoutMin))
	}
	if config.ElectionTimeoutMax <= 0 {
		newErrors = append(newErrors, fmt.Errorf("election timeout max must be positive, got %v", config.ElectionTimeoutMax))
	}
	// raft picks a random timeout in [min, max), so the range can't be empty
	if config.ElectionTimeoutMin >= config.ElectionTimeoutMax {
		newErrors = append(newErrors, fmt.Errorf("election timeout min (%v) must be less than max (%v)", config.ElectionTimeoutMin, config.ElectionTimeoutMax))
	}
	// followers would start elections while the leader is still healthy
	if config.HeartbeatInterval >= config.ElectionTimeoutMin {
		newErrors = append(newErrors, fmt.Errorf("heartbeat interval (%v) must be less than election timeout min (%v)", config.HeartbeatInterval, config.ElectionTimeoutMin))
	}

	// the node may list itself in peers, but only at its own peer address
	if selfAddr, ok := config.Peers[config.ID]; ok && selfAddr != config.PeerAddr {
		newErrors = append(newErrors, fmt.Errorf("peer list has %q at %q, but its peer address is %q", config.ID, selfAddr, config.PeerAddr))
	}

	seenAddrs := make(map[string]string)
	for id, addr := range config.Peers {
		if otherID, exists := seenAddrs[addr]; exists {
			newErrors = append(newErrors, fmt.Errorf("peers %q and %q share address %q", otherID, id, addr))
		}
		seenAddrs[addr] = id
	}

	return errors.Join(newErrors...)
}

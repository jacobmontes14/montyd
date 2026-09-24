// Package kvstore provides an in-memory key/value store that is modified
// only by applying encoded commands, so it can be replicated with Raft.
package kvstore

import (
	"fmt"
	"sync"

	pb "github.com/jacobmontes14/montyd/proto/generated"
	"google.golang.org/protobuf/proto"
)

type KeyStore struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewKeyStore() *KeyStore {
	return &KeyStore{data: make(map[string]string)}
}

func (kv *KeyStore) Apply(data []byte) error {
	var cmd pb.Command

	if err := proto.Unmarshal(data, &cmd); err != nil {
		return fmt.Errorf("decode error: %w", err)
	}

	switch cmd.GetOp() {
	case pb.Op_OP_SET:
		kv.set(cmd.GetKey(), cmd.GetValue())
	case pb.Op_OP_DELETE:
		kv.delete(cmd.GetKey())
	case pb.Op_OP_NOOP:
	default:
		return fmt.Errorf("unknown operation: %v", cmd.GetOp())
	}

	return nil
}

func (kv *KeyStore) set(key string, value string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	kv.data[key] = value
}

func (kv *KeyStore) Get(key string) (string, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

	value, ok := kv.data[key]
	return value, ok
}

func (kv *KeyStore) delete(key string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	delete(kv.data, key)
}

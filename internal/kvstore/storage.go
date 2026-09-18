package storage

import (
	"context"
	"sync"

	pb "github.com/jacobmontes14/montyd/proto/generated"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KeyStore struct {
	mu   sync.RWMutex
	data map[string]string
}

func (kv *KeyStore) Set(key string, value string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	kv.data[key] = value
}

func (kv *KeyStore) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

	value, ok := kv.data[req.GetKey()]

	if !ok {
		return nil, status.Errorf(codes.NotFound, "key %q not found", req.GetKey())
	}
	return &pb.GetResponse{Value: value}, nil
}

func (kv *KeyStore) Delete(key string) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	delete(kv.data, key)
}

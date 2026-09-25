package server

import (
	"testing"

	"github.com/jacobmontes14/montyd/internal/kvstore"
	pb "github.com/jacobmontes14/montyd/proto/generated"
	"github.com/stretchr/testify/assert"
)

const (
	value = "2"
	key   = "1"
)

func newTestKV(t *testing.T) *KV {
	t.Helper()

	store := kvstore.NewKeyStore()
	newKVStore := NewKV(store)

	return newKVStore
}

// happy path
func TestSetThenGet(t *testing.T) {
	kv := newTestKV(t)

	setRequest := &pb.SetRequest{
		Key:   key,
		Value: value,
	}

	_, err := kv.Set(t.Context(), setRequest)
	if err != nil {
		t.Fatalf("failed to set key: %v and value: %v. Err: %v", key, value, err)
	}

	getRequest := &pb.GetRequest{
		Key: key,
	}

	resp, err := kv.Get(t.Context(), getRequest)
	if err != nil {
		t.Fatalf("failed to get value: %v. Err: %v", value, err)
	}

	assert.Equal(t, value, resp.GetValue())
}

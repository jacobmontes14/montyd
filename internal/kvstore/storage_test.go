package kvstore

import (
	"testing"

	"github.com/stretchr/testify/assert"

	pb "github.com/jacobmontes14/montyd/proto/generated"
	"google.golang.org/protobuf/proto"
)

const (
	Key   = "1"
	Value = "2"
)

func encode(t *testing.T, cmd *pb.Command) []byte {
	t.Helper()

	enc, err := proto.Marshal(cmd)
	if err != nil {
		t.Fatalf("failed to encode: %v", err)
	}

	return enc
}

func TestSetThenGet(t *testing.T) {
	store := NewKeyStore()

	command := &pb.Command{
		Op:    pb.Op_OP_SET,
		Key:   Key,
		Value: Value,
	}

	encodedCommand := encode(t, command)
	if err := store.Apply(encodedCommand); err != nil {
		t.Fatalf("failed to apply command: %v", err)
	}

	val, ok := store.Get(Key)
	if !ok {
		t.Fatalf("failed to get value with given key: %v", Key)
	}
	assert.Equal(t, Value, val)
}

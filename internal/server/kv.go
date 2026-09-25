// Package server where where the network layer for KVstore lives
package server

import (
	"context"

	"github.com/jacobmontes14/montyd/internal/kvstore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"

	pb "github.com/jacobmontes14/montyd/proto/generated"
)

type KV struct {
	pb.UnimplementedKVStoreServer
	store *kvstore.KeyStore
}

func NewKV(store *kvstore.KeyStore) *KV {
	return &KV{store: store}
}

func (s *KV) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	key := req.GetKey()
	if key == "" {
		return nil, status.Error(codes.InvalidArgument, "key not provided")
	}
	value, ok := s.store.Get(key)

	if !ok {
		return nil, status.Errorf(codes.NotFound, "key %v not found", key)
	}

	return &pb.GetResponse{Value: value}, nil
}

func (s *KV) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	key := req.GetKey()
	value := req.GetValue()

	if key == "" {
		return nil, status.Error(codes.InvalidArgument, "key is required")
	}
	if value == "" {
		return nil, status.Error(codes.InvalidArgument, "value is required")
	}

	data, err := proto.Marshal(&pb.Command{
		Op:    pb.Op_OP_SET,
		Key:   key,
		Value: value,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "encode error: %v", err)
	}

	if err := s.store.Apply(data); err != nil {
		return nil, status.Errorf(codes.Internal, "store error: %v", err)
	}

	return &pb.SetResponse{}, nil
}

func (s *KV) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	key := req.GetKey()

	if key == "" {
		return nil, status.Error(codes.InvalidArgument, "key is required")
	}

	data, err := proto.Marshal(&pb.Command{
		Op:  pb.Op_OP_DELETE,
		Key: key,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "encode error: %v", err)
	}

	if err := s.store.Apply(data); err != nil {
		return nil, status.Errorf(codes.Internal, "store error: %v", err)
	}

	return &pb.DeleteResponse{}, nil
}

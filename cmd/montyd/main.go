package main

import (
	"log"
	"net"

	"github.com/jacobmontes14/montyd/internal/kvstore"
	"github.com/jacobmontes14/montyd/internal/server"
	pb "github.com/jacobmontes14/montyd/proto/generated"
	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to create listener: %v", err)
	}

	srv := grpc.NewServer()
	store := kvstore.NewKeyStore()
	pb.RegisterKVStoreServer(srv, server.NewKV(store))

	if err := srv.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

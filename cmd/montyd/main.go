package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/jacobmontes14/montyd/internal/config"
	"github.com/jacobmontes14/montyd/internal/kvstore"
	"github.com/jacobmontes14/montyd/internal/server"
	pb "github.com/jacobmontes14/montyd/proto/generated"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("failed to read in configurations: %v", err)
	}

	list, err := net.Listen("tcp", conf.ClientAddr)
	if err != nil {
		log.Fatalf("failed to create listener: %v", err)
	}

	log.Printf("starting up montyd")
	srv := grpc.NewServer()
	store := kvstore.NewKeyStore()
	pb.RegisterKVStoreServer(srv, server.NewKV(store))

	if err := srv.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

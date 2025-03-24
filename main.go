package main

import (
	server "github.com/jacobmontes14/montyd/internal/api"
)

func main() {
	s := server.NewServer(":8080")
	s.Start()
}

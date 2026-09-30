#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"

GEN=proto/generated

mkdir -p "$GEN/kvstorepb"
protoc -I proto \
  --go_out="$GEN/kvstorepb" --go_opt=paths=source_relative \
  --go-grpc_out="$GEN/kvstorepb" --go-grpc_opt=paths=source_relative \
  kvstore.proto command.proto

mkdir -p "$GEN/raftpb"
protoc -I proto \
--go_out="$GEN/raftpb" --go_opt=paths=source_relative \
--go-grpc_out="$GEN/raftpb" --go-grpc_opt=paths=source_relative \
raft.proto

#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")"

GEN=proto/generated

protoc -I proto \
  --go_out="$GEN" --go_opt=paths=source_relative \
  --go-grpc_out="$GEN" --go-grpc_opt=paths=source_relative \
  kvstore.proto command.proto

if [[ -f proto/raft.proto ]]; then
  mkdir -p "$GEN/raftpb"
  protoc -I proto \
    --go_out="$GEN/raftpb" --go_opt=paths=source_relative \
    --go-grpc_out="$GEN/raftpb" --go-grpc_opt=paths=source_relative \
    raft.proto
fi

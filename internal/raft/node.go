// package raft is where the raft implementation lives
package raft

type Node struct {
	currentTerm int32
	votedFor    int32
	log         []string
}

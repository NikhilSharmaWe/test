package main

import (
	pb "github.com/NikhilSharmaWe/test/proto"
	"google.golang.org/protobuf/proto"
)

func ComposeNamePB(n pb.Name) ([]byte, error) {
	b, err := proto.Marshal(&n)
	return b, err
}

func ParseNamePB(b []byte) (pb.Name, error) {
	n := pb.Name{}
	err := proto.Unmarshal(b, &n)
	return n, err
}

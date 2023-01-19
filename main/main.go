package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"strangerDB/pr"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	s := grpc.NewServer()
	srv := pr.Server{}
	pr.RegisterDBServer(s, &srv)
	l, err := net.Listen("tcp", ":8080")
	check(err)
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}

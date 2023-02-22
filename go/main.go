package main

import (
	"fmt"
	"github.com/blorente/standalone-repros/go/lib"
	"github.com/blorente/standalone-repros/proto"
	"google.golang.org/grpc"
	"os"
)

func main() {
	lib.AddToTwo()
	fmt.Println("FOO:", os.Getenv("FOO"))
}

func serv(grpcSrv *grpc.Server, protoSrv *proto.FooServiceServer) {
	stuff := "stuff"
	req := proto.HelloRequest{Name: &stuff}
	req.GetName()
}

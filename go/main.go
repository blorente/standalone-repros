package main

import (
	"fmt"
	"github.com/blorente/standalone-repros/go/lib"
	"github.com/blorente/standalone-repros/proto"
	"google.golang.org/grpc"
	"os"
)

func main() {
	lib.MyFun()
	fmt.Println("FOO:", os.Getenv("FOO"))
}

func serv(grpcSrv *grpc.Server, protoSrv *proto.FooServiceServer) {}

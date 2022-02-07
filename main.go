package main

import (
	"github.com/blorente/standalone-repros/lib"
	"github.com/blorente/standalone-repros/proto"
	"google.golang.org/grpc"
)

func main() {
	lib.MyFun()
}

func serv(grpcSrv *grpc.Server, protoSrv *proto.TestServiceServer) {}

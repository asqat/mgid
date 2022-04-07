package main

import (
	"github.com/asqat/mgid/employee/pkg/db"
	"github.com/asqat/mgid/employee/pkg/server"
)

//go:generate protoc --go_out=pkg --go-grpc_out=pkg --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/employee.proto
func main() {
	err := db.InitDummyDatabase()
	if err != nil {
		panic(err)
	}

	var port uint16 = 5600
	if err = server.StartGrpcServer(port); err != nil {
		panic(err)
	}
}

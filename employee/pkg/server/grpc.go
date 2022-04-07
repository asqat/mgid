package server

import (
	"fmt"
	"github.com/asqat/mgid/employee/pkg/db"
	employee "github.com/asqat/mgid/employee/pkg/proto"
	"google.golang.org/grpc"
	"net"
)

func StartGrpcServer(port uint16) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	grpcServer := &server{
		data: db.GetDatabase(), // apply a dummy database from dataset.json
	}

	server := grpc.NewServer()

	employee.RegisterEmployeeServiceServer(server, grpcServer)

	fmt.Printf("gRPC server have start on %d port\n", port)
	if err := server.Serve(listen); err != nil {
		return err
	}
	return nil
}

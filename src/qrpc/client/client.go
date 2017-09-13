package client

import (
	pb "qrpc/utils"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	Gconn *grpc.ClientConn
)

func InitGConn() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	Gconn = conn
	if err != nil {
		fmt.Println("did not connect: %v", err)
	}
}

func Log(content string, serviceType, serviceFlag, level string) {
	c := pb.NewLogClient(Gconn)

	r, err := c.Record(context.Background(), &pb.LogRequest{ServiceType: serviceType, ServiceFlag: serviceFlag, Level: level, Content: content})
	if err != nil {
		fmt.Println("could not greet: %v", err)
		//Close()
		//InitGConn()
		//Log(content)

	}
	fmt.Println("Greeting: %s", r.Status)
}

func Close() {
	defer Gconn.Close()
}

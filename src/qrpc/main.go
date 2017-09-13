package main

import (
	"fmt"
	"qrpc/client"
)

func main() {
	fmt.Println("hello word")
	client.InitGConn()
	defer client.Close()
}

package main

import (
	"fmt"
	"tcp_test/app"
	"time"
)

func main() {
	fmt.Println("tcp-test start1")

	tcpInst := app.MakeTcpHandler("tcptest", 6000, "127.0.0.1")

	fmt.Println("tcp-test start2")
	go app.RunTcp(tcpInst)

	for {
		time.Sleep(time.Millisecond * 100)
	}

	tcpInst.Close()

	fmt.Println("tcp-test close")
}

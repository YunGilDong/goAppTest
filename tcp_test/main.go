package main

import (
	"fmt"
	"net/http"
	"tcp_test/app"
	"tcp_test/model"
	"time"

	"github.com/gorilla/handlers"
)

func main() {
	m := app.MakeHandler()
	app.InitObjectDB()

	app.ObjDB.LocDB[1] = &model.Local{LOC_ID: 1}

	go func() {
		err := http.ListenAndServe(":5000", handlers.CORS()(m))
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("tcp-test start1")

	tcpInst := app.MakeTcpHandler("tcptest", 6000, "127.0.0.1")
	chTcpMng := app.TcpMngChannel{
		Terminate_manage:   make(chan bool),
		Terminate_manageRx: make(chan bool),
		Terminate_manageTx: make(chan bool),
	}

	fmt.Println("tcp-test start2")
	go app.RunTcp(tcpInst, chTcpMng)

	for {
		time.Sleep(time.Millisecond * 1000)
		fmt.Println("run...")
	}

	defer func() {
		chTcpMng.Terminate_manage <- true
		chTcpMng.Terminate_manageRx <- true
		chTcpMng.Terminate_manageTx <- true

		tcpInst.Close()

		fmt.Println("tcp-test close")
	}()
}

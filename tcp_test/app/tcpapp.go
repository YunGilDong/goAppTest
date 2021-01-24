package app

import (
	"fmt"
	"tcp_test/comm"
	"time"
)

type AppTcpHandler struct {
	tcp comm.CommHandler

	m_message []byte
	m_length  int
	// general commbase
}

// tcp init
//

// func (t *TcpHandler) ManageRX() bool {
// 	return false
// }
// func (t *TcpHandler) ManageTX() bool {

func MakeTcpHandler(name string, port int, addr string) *AppTcpHandler {
	//var tcpInst *TcpHandler = comm.NewTcpHandler(name, port, addr)
	a := &AppTcpHandler{
		tcp:       comm.NewTcpHandler(name, port, addr),
		m_message: []byte{},
		m_length:  0,
	}

	return a
}

func (a *AppTcpHandler) ManageComm() {

	// rx & tx loop
	// manage RX
	// manage TX

	// manage connect
	if !a.tcp.IsConnected() {
		a.Close()
		a.tcp.Connect()
	}

	a.ManageRx()
	a.ManageTx()
}

func (a *AppTcpHandler) ManageRx() bool {
	var data []byte = []byte{}
	fmt.Println("manageRX")
	len, err := a.tcp.Read(data)
	fmt.Println("manageRX", len)
	if err != nil {
		fmt.Println("read err: ", err)
		a.Close()
		return false
	}
	fmt.Println("read len", len, data)
	return true
}

func (a *AppTcpHandler) ManageTx() bool {
	var data []byte = []byte("hello world tcp app")
	len, err := a.tcp.Send(data)
	fmt.Println("manageTX", len)
	if err != nil {
		fmt.Println("send err:", err)
		a.Close()
		return false
	}
	fmt.Println("send len:", len, data)
	return true
}

func (a *AppTcpHandler) Close() {
	a.tcp.Close()
}

func RunTcp(a *AppTcpHandler) {

	// ManageComm

	for {
		a.ManageComm()

		time.Sleep(time.Millisecond * 2000)
	}

}

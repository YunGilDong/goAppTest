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
}

func (a *AppTcpHandler) Manage(ch_connected chan bool) {

	select {
	case connected := <-ch_connected:
		if !connected {
			a.tcp.Close()
			a.tcp.Connect()
		}

	default:
		if !a.tcp.IsConnected() {
			fmt.Println("isConnect..", a.tcp.IsConnected())
			a.tcp.Connect()
		}
	}
}

func (a *AppTcpHandler) ManageRx(ch_connected chan bool) bool {
	data := make([]byte, 10240, 10240)
	fmt.Println("manageRX")
	len, err := a.tcp.Read(data)
	fmt.Println("manageRX", len)
	if err != nil {
		fmt.Println("read err: ", err)
		a.Close()
		ch_connected <- false
		return false
	}
	fmt.Println("read len", len, data[0:len])
	return true
}

func (a *AppTcpHandler) ManageTx(ch_connected chan bool) bool {
	var data []byte = []byte("hello world tcp app")
	len, err := a.tcp.Send(data)
	fmt.Println("manageTX", len)
	if err != nil {
		fmt.Println("send err:", err)
		a.Close()
		ch_connected <- false
		return false
	}
	fmt.Println("send len:", len, string(data))
	return true
}

func (a *AppTcpHandler) Close() {
	a.tcp.Close()
}

func RunTcp(a *AppTcpHandler, chTcpMng TcpMngChannel) {
	ch_connected := make(chan bool)

	// ManageComm
	a.tcp.Connect()

	go a.Manage(ch_connected)

	go func() {
		for {
			select {
			case terminate := <-chTcpMng.Terminate_manage:
				if terminate {
					break
				}

			default:
				a.Manage(ch_connected)
				time.Sleep(time.Millisecond * 1000)
			}
		}

	}()

	go func() {
		for {
			select {
			case terminate := <-chTcpMng.Terminate_manageRx:
				if terminate {
					break
				}
			default:
				if a.tcp.IsConnected() {
					a.ManageRx(ch_connected)
				}
				time.Sleep(time.Millisecond * 1000)
			}
		}
	}()

	go func() {
		for {
			select {
			case terminate := <-chTcpMng.Terminate_manageTx:
				if terminate {
					break
				}
			default:
				if a.tcp.IsConnected() {
					a.ManageTx(ch_connected)
				}
				time.Sleep(time.Millisecond * 1000)
			}

		}
	}()
}

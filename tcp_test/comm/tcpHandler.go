package comm

import (
	"fmt"
	"log"
	"net"
)

//---------------------------------------------------------------------------
// Implement CommHandler
//---------------------------------------------------------------------------
type TcpHandler struct {
	name      string
	port      int
	addr      string
	Connected bool
	tcp       *net.TCPConn
}

func newTcpHandler(name string, port int, addr string) CommHandler {

	return &TcpHandler{
		name: name,
		port: port,
		addr: addr,

		Connected: false,
	}
}

//---------------------------------------------------------------------------
// InitComm
//---------------------------------------------------------------------------
func (t *TcpHandler) InitComm(name string, port int, addr string) {
	// 필요없나?

}

//---------------------------------------------------------------------------
// Connect
//---------------------------------------------------------------------------
func (t *TcpHandler) Connect() (bool, error) {
	fmt.Println("connete to server ...")
	target := fmt.Sprintf("%s:%d", t.addr, t.port)
	raddr, err := net.ResolveTCPAddr("tcp", target)
	if err != nil {
		log.Fatal(err)
		return false, err
	}

	// connect to server
	conn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	t.tcp = conn
	t.tcp.SetNoDelay(false)

	fmt.Println("conneted server ok")
	return true, nil
}

//---------------------------------------------------------------------------
// SendMessage
//---------------------------------------------------------------------------
func (t *TcpHandler) Send(b []byte) (int, error) {
	var cnt int = 0
	cnt, err := t.tcp.Write(b)
	if err != nil {
		return cnt, err
	}

	return cnt, nil
}

//---------------------------------------------------------------------------
// Read
//---------------------------------------------------------------------------
func (t *TcpHandler) Read(data []byte) (int, error) {
	var cnt int = 0
	fmt.Println("Read(1)")
	cnt, err := t.tcp.Read(data)
	fmt.Println("Read(2)", cnt)

	if err != nil {
		return cnt, err
	}

	return cnt, nil
}

func (t *TcpHandler) IsConnected() bool {
	return t.Connected
}

//---------------------------------------------------------------------------
// ClearEnv
//---------------------------------------------------------------------------
func (t *TcpHandler) Close() {
	// clear tcp env
	// close tcp connect
	if !t.Connected {
		return
	}

	t.Connected = false
	err := t.tcp.Close()
	if err != nil {
		fmt.Println("Close err..", err)
	}
}

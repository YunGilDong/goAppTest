package comm

type CommHandler interface {
	InitComm(name string, port int, addr string)
	Connect() (bool, error)
	Send(b []byte) (int, error)
	Read(data []byte) (int, error)
	IsConnected() bool
	Close()
}

func NewTcpHandler(name string, port int, addr string) CommHandler {
	return newTcpHandler(name, port, addr)
}

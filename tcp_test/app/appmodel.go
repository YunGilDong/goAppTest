package app

type TcpMngChannel struct {
	Terminate_manage chan bool
	Terminate_manageRx chan bool
	Terminate_manageTx chan bool
}
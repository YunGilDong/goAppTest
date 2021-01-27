package app

import "tcp_test/model"

//---------------------------------------------------------------------------
//
//---------------------------------------------------------------------------
type ObjectDB struct {
	LocDB map[int]*model.Local
}

var ObjDB ObjectDB

//---------------------------------------------------------------------------
// Event Data Format
//---------------------------------------------------------------------------
type EventData struct {
	Msgtype string
	Data    string
}

//---------------------------------------------------------------------------
// Event Channel
//---------------------------------------------------------------------------
var ChEvent chan EventData

type TcpMngChannel struct {
	Terminate_manage   chan bool
	Terminate_manageRx chan bool
	Terminate_manageTx chan bool
}

type DBHandler interface {
	Ping() // 연결확인
}

func NewDBHandler() DBHandler {
	return newMariadbHandler()
}

func InitObjectDB() {
	ObjDB.LocDB = make(map[int]*model.Local)
}

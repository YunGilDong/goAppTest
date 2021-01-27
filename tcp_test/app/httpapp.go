package app

import (
	"fmt"
	"net/http"

	"github.com/antage/eventsource"
	"github.com/gorilla/mux"
)

//------------------------------------------------------------------------------
// HttpAppHandler
//------------------------------------------------------------------------------
type HttpAppHandler struct {
	http.Handler
}

//------------------------------------------------------------------------------
// testHandler
//------------------------------------------------------------------------------
func (a *HttpAppHandler) testHandler(w http.ResponseWriter, r *http.Request) {
	pLoc := ObjDB.LocDB[1]
	pLoc.TestSetId(10)
	fmt.Fprint(w, "hello.")
}

//------------------------------------------------------------------------------
// testHandler2
//------------------------------------------------------------------------------
func (a *HttpAppHandler) testHandler2(w http.ResponseWriter, r *http.Request) {
	pLoc := ObjDB.LocDB[1]

	fmt.Fprint(w, "hello.", pLoc.LOC_ID)
}

//------------------------------------------------------------------------------
// AppClose
//------------------------------------------------------------------------------
func (a *HttpAppHandler) AppClose() {
	// app close
}

//------------------------------------------------------------------------------
// initEventSource
//------------------------------------------------------------------------------
func initEventSource() eventsource.EventSource {
	es := eventsource.New(
		eventsource.DefaultSettings(),
		func(req *http.Request) [][]byte {
			return [][]byte{
				[]byte("X-Accel-Buffering: no"),
				[]byte("Access-Control-Allow-Origin: *"),
				[]byte("UTF-8"),
			}
		},
	)

	return es
}

func processEventMsg(es eventsource.EventSource) {
	for msg := range ChEvent {
		fmt.Println(msg)
	}
}

//------------------------------------------------------------------------------
// MakeHandler
//------------------------------------------------------------------------------
func MakeHandler() *HttpAppHandler {

	r := mux.NewRouter()
	a := &HttpAppHandler{
		Handler: r,
	}

	// Server send event
	ChEvent = make(chan EventData)
	es := initEventSource()
	go processEventMsg(es)

	// API
	r.HandleFunc("/test", a.testHandler).Methods("GET")
	r.HandleFunc("/test2", a.testHandler2).Methods("GET")
	r.Handle("/event", es)

	return a
}

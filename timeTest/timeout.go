package main

import (
	"log"
	"time"
)

func main() {
	log.Println("start")

L:
	for {
		select {
		case <-time.After(time.Second * 5):
			log.Println("time out!")
			break L
			// default:
			// 	time.Sleep(time.Second * 1)
		}
	}

	log.Println("end")
}

package main

import (
	"bytes"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var buf = bytes.NewBuffer(nil)

func main() {
	go http.ListenAndServe(":6060", nil)

	t1 := time.NewTicker(1 * time.Second)
	t2 := time.NewTicker(10 * time.Second)

	for true {
		select {
		case <-t1.C:
			mem()
			cpu()
		case <-t2.C:
			buf.Reset()
		}
	}
}

func mem() {
	buf.Write(make([]byte, 1*1024*1024))
}

func cpu() {
	for i := 0; i < 1000000000; i++ {

	}
}

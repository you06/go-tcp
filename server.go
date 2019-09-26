package main

import (
	"bufio"
	"log"
	"math"
	"net"
	"strconv"
	"strings"
)

func initServer(host string) {
	var count uint64
	log.Println("init server on address", host)
	ln, err := net.Listen("tcp", host)
	if err != nil {
		log.Panic(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go checkConnection(conn)
		count++

		if count%uint64(math.Pow10(len(strconv.Itoa(int(count)))-1)) == 0 {
			log.Printf("receive %d packet \n", count)
		}
	}
}

func checkConnection(conn net.Conn) {
	for {
		buf := make([]byte, 20479)
		r := bufio.NewReader(conn)
		w := bufio.NewWriter(conn)
		len, err := r.Read(buf)
		if err != nil {
			// fmt.Println("Error reading", err)
			return
		}
		if checkData(buf[:len]) {
			w.Write([]byte{0})
		} else {
			w.Write([]byte{1})
			log.Println("error body")
			log.Println(buf[:len])
		}
		w.Flush()
	}
}

func checkData(buf []byte) bool {
	if len(buf) <= 4 {
		return true
	}
	for i := 0; i < len(buf)-1; i++ {
		if i+4 < len(buf) {
			if buf[i+1] == 13 && buf[i+2] == 10 && buf[i+3] == 13 && buf[i+4] == 10 {
				break
			}
		}
		after := (int(buf[i])*mul)%dev + 1
		if !(after == int(buf[i+1])) {
			log.Printf("should be %b after %b, got %b %b %b\n", after, buf[i], buf[i+1], mul, dev)
			return false
		}
	}
	return true
}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, "\r\n\r\n")
	return
}

func logErrorBody(buf []byte) {

}

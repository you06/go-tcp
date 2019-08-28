package main

import (
	"fmt"
	"log"
	"net"
)

func initServer(host string) {
	log.Println("init server on address", host)
	ln, err := net.Listen("tcp", host)
	if err != nil {
		log.Panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err.Error())
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	log.Println(conn)
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
	}
}

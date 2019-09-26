package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

func initClient(host string, interval int) {
	for range time.Tick(time.Duration(interval) * time.Millisecond) {
		go sendTCPPacket(host)
	}
}

func sendTCPPacket(host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer conn.Close()

	msg := makeMessage()
	conn.Write(msg)
	conn.Write([]byte(stopCharacter))
	buff := make([]byte, 1024)
	conn.Read(buff)
	if buff[0] != 0 {
		log.Println(msg)
	}
}

func makeMessage() []byte {
	start := rand.Intn(255)
	msgLen := rand.Intn(5000) + 1000
	var msg []byte
	for i := 0; i < msgLen; i++ {
		msg = append(msg, byte(start))
		start = (start*mul)%dev + 1
	}
	return msg
}

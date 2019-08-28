package main

import (
	"log"
	"time"
)

func initClient(host string, interval int) {
	for _ = range time.Tick(time.Duration(interval) * time.Millisecond) {
		go sendTCPPacket(host)
	}
}

func sendTCPPacket(host string) {
	log.Println(host)
}

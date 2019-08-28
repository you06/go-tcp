package main


import (
	"flag"
	"log"
)


var (
	mode string
	host string
	interval int
)

func init() {
	flag.StringVar(&mode, "mode", "client", "running mode, client or server")
	flag.StringVar(&host, "host", "127.0.0.1:9999",
		"server listening address or client connecting address")
	flag.IntVar(&interval, "int", 100, "client send interval")
}

func main() {
	flag.Parse()

	switch mode {
	case "server": {
		initServer(host)
	}
	case "client": {
		initClient(host, interval)
	}
	default: {
		log.Panic("Invalid running mode.")
	}
	}
}

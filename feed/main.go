package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "ingestor:9000")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		// Simulate microburst
		for i := 0; i < 10000; i++ {
			msg := fmt.Sprintf("TICK,%d,%d", time.Now().UnixNano(), i)
			conn.Write([]byte(msg))
		}
		time.Sleep(1 * time.Second)
	}
}

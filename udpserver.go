package main

import (
	"log"
	"net"
)

func main() {
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 10053,
	}

	listener, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Now listen")

	buf := make([]byte, 512)

	for {
		l, addr, err := listener.ReadFromUDP(buf)
		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("From: %s, Data: %x", addr, buf[:l])
	}
}

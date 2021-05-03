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

		if l <= 12 {
			log.Printf("From: %s, Data: %x", addr, buf[:l])
			continue
		}
		// buffer for NAME
		tm := make([]byte, 256)
		b1 := buf[:11]
		b2 := buf[12:l]

		// current position in NAME
		p := 0
		for p < len(b2) && b2[p] != 0 {
			// label length
			ll := int(b2[p])
			p++
			tm = append(tm, b2[p:ll+p]...)
			tm = append(tm, '.')
			log.Printf("p=%3d, label length=%2d, name=%s", p, ll, tm)
			p += ll
		}
		p++
		log.Printf("QTYPE=%02X, QCLASS=%02X", b2[p:p+2], b2[p+2:p+4])

		log.Printf("From: %s, Data: %012X , %x , %s", addr, b1, b2, tm)
	}
}

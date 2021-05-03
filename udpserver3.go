package main

import (
	"encoding/binary"
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

		go func() {
			if l <= 12 {
				log.Printf("From: %s, Data: %x", addr, buf[:l])
				return

			}
			log.Printf("l=%d ln=%d lc=%d", l, len(buf), cap(buf))
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
			log.Printf("From: %s, Data: %012X , %x , %s", addr, b1, b2, tm)
			if len(b2) > p+4 {
				log.Printf("QTYPE=%02X, QCLASS=%02X", b2[p:p+2], b2[p+2:p+4])
			} else {
				return
			}

			// only for IN A
			if binary.BigEndian.Uint16(b2[p:p+2]) == 1 && binary.BigEndian.Uint16(b2[p+2:p+4]) == 1 {
				log.Printf("%s\t\t\tIN A", tm)
				// Question Data
				qd := b2[0 : p+4]
				// header
				buf := []byte{0x00, 0x00, 0x81, 0x80, 0x00, 0x01, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00}
				// copy question data
				buf = append(buf, qd...)
				// copy question data
				buf = append(buf, qd...)
				// TTL 300
				buf = append(buf, []byte{0x00, 0x00, 0x01, 0x2c}...)
				// Length = 4
				buf = append(buf, []byte{0x00, 0x04}...)
				// IP 1.1.1.1
				buf = append(buf, []byte{0x01, 0x01, 0x01, 01}...)
				// replace ID
				buf[0] = b1[0]
				buf[1] = b1[1]
				log.Printf("%x", buf)
				wl, err := listener.WriteTo(buf, addr)
				if err != nil {
					log.Fatalln("WriteTo() error: %s\n", err)
				} else {
					log.Printf("Wrote %d bytes to socket\n", wl)
				}

			}
		}()
	}
}

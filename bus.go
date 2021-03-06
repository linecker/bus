package main
 
import (
	"log"
	"net"
	"fmt"
	"syscall"
	"time"
)

const bufferSize = 512

var (
	address *net.UDPAddr
	connection *net.UDPConn
)

func Send(message string) {
	if connection != nil {
		fmt.Printf("Send: %s\n", message)
		connection.Write([]byte(message))
	}
}

type Receive func(*net.UDPAddr, int, []byte)

func Setup(callback Receive) {
	address, err := net.ResolveUDPAddr("udp", "239.0.0.1:9876")
	if err != nil {
		log.Fatal(err)
	}
	c, err := net.DialUDP("udp", nil, address)
	if err != nil {
		log.Fatal(err)
	}
	connection = c
	f, err := connection.File()
	err = syscall.SetsockoptInt(int(f.Fd()), syscall.IPPROTO_IP, syscall.IP_MULTICAST_LOOP, 1)	
	if err != nil {
		log.Fatal(err)
	}
	l, err := net.ListenMulticastUDP("udp", nil, address)
	if err != nil {
		log.Fatal(err)
	}
	l.SetReadBuffer(bufferSize)
	go func() {
		for {
			fmt.Println("read loop")
			b := make([]byte, bufferSize)
			n, src, err := l.ReadFromUDP(b)
			if err != nil {
				log.Fatal("ReadFromUDP failed", err)
			}
			callback(src, n, b)
		}
	}()
}

func recv_impl(source *net.UDPAddr, n int, b []byte) {
	fmt.Printf("Recv: %s\n" + string(b))
}

func main() {
	Setup(recv_impl)
	for {
		Send("send loop")
		time.Sleep(1 * time.Second)
	}
}


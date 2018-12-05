package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	port := ":8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	addr, e := net.ResolveTCPAddr("tcp", port)
	if e != nil {
		fmt.Println(`resolve error :`, e)
		return
	}
	l, e := net.ListenTCP("tcp", addr)
	if e != nil {
		fmt.Println(`listen error :`, e)
		return
	}
	for {
		c, e := l.Accept()
		if e != nil {
			fmt.Println(`accept error :`, e)
			return
		}
		go handleCon(c)
	}
}
func handleCon(c net.Conn) {
	defer c.Close()
	fmt.Println("new connection:", c.RemoteAddr().String())
	b := make([]byte, 2048000)
	n, e := c.Read(b)
	if e != nil {
		fmt.Println(`read error :`, e)
		return
	}
	fmt.Println("\n\n")
	fmt.Println(string(b[:n]))
	fmt.Println("------------------------------------------------")
	backData := "HTTP/1.1 200 OK\r\nContent-Length: " + fmt.Sprint(n) + "\r\n\r\n" + string(b[:n]) + "\r\n\r\n"
	fmt.Fprint(c, backData)
}

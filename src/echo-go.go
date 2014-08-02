package main

import (
	"net"
	"strconv"
	"fmt"
	"bufio"
)

const PORT = 20000

func main() {
	server,_ := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if server == nil {
		panic("could not be start");
	}
	conns := listen(server)
	for {
		fmt.Println("Handle connection")
		go handleConn(<-conns)
	}
}

func listen(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0
	go func() {
		for {
			client,_ := listener.Accept()
			if client == nil {
				fmt.Println("couldn't accept")
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line,err := b.ReadBytes('\n')
		if err != nil {
			break
		}
		client.Write(line)
	}

}

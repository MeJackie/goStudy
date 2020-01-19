package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	var tcpAddr *net.TCPAddr
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()

	fmt.Printf("Server ready to read ...")
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}

func tcpPipe(conn net.Conn)  {
	defer func() {
		fmt.Println(conn.RemoteAddr().String() + " Disconnected")
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	i := 0
	for {
		message, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(string(message))

		time.Sleep(time.Second * 3)

		msg := time.Now().String() + conn.RemoteAddr().String() + " Server say hello! \n"
		b := []byte(msg)
		conn.Write(b)
		i++

		if i > 10 {
			break
		}

	}
}

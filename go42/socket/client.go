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
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:999")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	if err != nil {
		fmt.Println("Client connect server error !" + err.Error())
		return
	}

	defer conn.Close()

	fmt.Println(conn.LocalAddr().String() + ": client connected !")
	onMessageRecived(conn)
}

func onMessageRecived(conn *net.TCPConn)  {
	b := []byte(conn.LocalAddr().String() + " Say hello to server ... \n")
	conn.Write(b)

	io.WriteString()

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		fmt.Println(msg)

		time.Sleep(time.Second * 3)
		fmt.Println("writing ...")
		b := []byte(conn.LocalAddr().String() + "write data to server ... \n")
		_, err = conn.Write(b)
		if err != nil {
			fmt.Println(err)
			break
		}

	}

}

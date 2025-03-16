package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// Use nc localhost 8080 in terminal to test
// basic coneection with server
func main() {
	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer ls.Close()

	for {
		con, err := ls.Accept()
		if err != nil {
			panic(err)
		}

		go func(connection net.Conn) {
			for {
				data, _ := bufio.NewReader(connection).ReadString('\n')
				if strings.Contains(data, "exit") {
					break
				}

				fmt.Printf("Data Recebida: %s", data)
				connection.Write([]byte("Mensagem recebida"))
			}

			connection.Close()
		}(con)
	}
}

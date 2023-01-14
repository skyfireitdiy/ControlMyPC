package server

import (
	"io"
	"log"
	"net"
	"os"
)

func ServerMain() {
	clientManageAddr := os.Args[2]
	clientIOAddr := os.Args[3]
	controlAddr := os.Args[4]
	controlIOAddr := os.Args[5]
	ioChan := make(chan []byte, 4096)

	var clientManageConn net.Conn
	var clientIOConn net.Conn

	clientManageListener, err := net.Listen("tcp", clientManageAddr)
	if err != nil {
		log.Fatalln(err)
	}

	clientIOListener, err := net.Listen("tcp", clientIOAddr)
	if err != nil {
		log.Fatalln(err)
	}

	controlListener, err := net.Listen("tcp", controlAddr)
	if err != nil {
		log.Fatalln(err)
	}

	controlIOListener, err := net.Listen("tcp", controlIOAddr)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		for {
			conn, err := controlListener.Accept()
			if err != nil {
				log.Fatalln(err)
			}

			io.Copy(clientManageConn, conn)
		}
	}()

	go func() {
		for {
			conn, err := controlIOListener.Accept()
			if err != nil {
				log.Fatalln(err)
			}

			go io.Copy(clientIOConn, conn)
			for c := range ioChan {
				total := len(c)
				sendLen := 0
				for sendLen != total {
					n, err := conn.Write(c[sendLen:])
					if err != nil {
						log.Println(err)
						return
					}
					sendLen += n
				}
			}
		}
	}()

	go func() {
		for {
			conn, err := clientManageListener.Accept()
			if err != nil {
				log.Fatalln(err)
			}
			clientManageConn = conn
		}
	}()

	func() {
		for {
			conn, err := clientIOListener.Accept()
			if err != nil {
				log.Fatalln(err)
			}

			clientIOConn = conn

			for {
				buf := make([]byte, 4096)
				n, err := conn.Read(buf)
				if err != nil {
					log.Println(err)
					break
				}

				select {
				case ioChan <- buf[:n]:
				default:
					log.Println("channel full")
				}
			}
		}
	}()

}

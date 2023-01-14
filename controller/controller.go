package controller

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func ControllerMain() {
    manageAddr := os.Args[2]
    ioAddr := os.Args[3]

    manageConn ,err:= net.Dial("tcp", manageAddr)
    if err != nil {
        log.Fatalln(err)
    }

    ioConn,err:= net.Dial("tcp", ioAddr)
    if err != nil {
        log.Fatalln(err)
    }

    go io.Copy(ioConn, os.Stdin)
    go io.Copy(os.Stdout, ioConn)

    ch := make(chan os.Signal, 1)
    signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
    for {
        sig :=<- ch
        msg := fmt.Sprintf("signal %d\n", sig)
        manageConn.Write([]byte(msg))
    }
}

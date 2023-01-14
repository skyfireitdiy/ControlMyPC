package client

import (
	"cmp/common"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"syscall"
)


func ClientMain(){
    serverManageAddr := os.Args[2]
    serverIOAddr := os.Args[3]

    serverManageConn, err := net.Dial("tcp", serverManageAddr)
    if err!= nil {
        log.Fatalln(err)
    }

    serverIOConn, err := net.Dial("tcp", serverIOAddr)
    if err != nil {
        log.Fatalln(err)
    }

    cmd, err := common.StartCommand("fish")
    if err != nil {
        log.Fatalln(err)
    }

    go io.Copy(serverIOConn, cmd.File)
    go io.Copy(cmd.File, serverIOConn)
    go common.HandleMsg(serverManageConn, func(msg string) error {
        log.Println(msg)
        args := strings.Split(msg, " ")
        switch args[0]{
        case "signal":{
            s, _ := strconv.Atoi(args[1])
            cmd.Cmd.Process.Signal(syscall.Signal(s))
        }
        }
        return nil
    })
    cmd.Wait()
}

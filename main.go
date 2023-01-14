package main

import (
	"cmp/client"
	"cmp/controller"
	"cmp/server"
	"os"
)

func main() {
    module := os.Args[1]
    switch (module){
    case "client":
        client.ClientMain()
    case "server":
        server.ServerMain()
    case "controller":
        controller.ControllerMain()
    }
}

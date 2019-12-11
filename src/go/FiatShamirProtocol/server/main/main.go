package main

import "cryptocrouse/src/go/FiatShamirProtocol/server/srvr"

func main() {
	server := srvr.ServerInit()
	server.Run()
}



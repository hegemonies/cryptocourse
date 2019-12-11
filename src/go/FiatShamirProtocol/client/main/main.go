package main

import "cryptocrouse/src/go/FiatShamirProtocol/client/clnt"

func main() {
	client := clnt.Client{}
	client.ConnectToServer()
	client.StartProof()
}

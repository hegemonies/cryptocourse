package main

import "cryptocrouse/src/go/FiatShamirProtocol/rogue/rg"

func main() {
	rogue := rg.InitRogue()
	rogue.ConnectToServer()
	rogue.TryToAcceptSecret()
}

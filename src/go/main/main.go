package main

import (
	"cryptocrouse/src/go/ElGamalCode"
	"cryptocrouse/src/go/FileWrapper"
)

func main() {
	producerName := "Alice"
	consumerName := "Bob"

	cryptosystem := ElGamalCode.CryptoSystem{}
	cryptosystem.Init()
	_ = cryptosystem.AddUser(producerName)
	_ = cryptosystem.AddUser(consumerName)

	filenameTestData := "test_data.png"

	cryptosystem.SendMessageFromFile(producerName, consumerName, filenameTestData)

	consumerInSystem := cryptosystem.Users[consumerName]
	FileWrapper.WriteToFile("1" + filenameTestData, consumerInSystem.GetMessage())
}

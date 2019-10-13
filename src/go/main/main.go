package main

import (
	"cryptocrouse/src/go/FileWrapper"
	"cryptocrouse/src/go/ShamirCode"
)

func main() {
	producerName := "Alice"
	consumerName := "Bob"

	cryptosystem := ShamirCode.CryptoSystem{}
	cryptosystem.Init()
	_ = cryptosystem.AddUser(producerName)
	_ = cryptosystem.AddUser(consumerName)

	filename_test_data := "test_data.png"
	//filename_test_data := "test.txt"

	cryptosystem.SendMessageFromFile(producerName, consumerName, filename_test_data)
	consumer := cryptosystem.Users[consumerName]

	cryptosystem.PrintUsers()

	FileWrapper.WriteToFile("1" + filename_test_data, consumer.GetMessage())
}

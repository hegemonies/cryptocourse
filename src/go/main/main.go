package main

import (
	"cryptocrouse/src/go/RSA"
	"fmt"
)

func main() {
	//filenameTestData := "test_data.png"
	//
	//file, err := os.Open(filenameTestData)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}
	//
	//defer func() {
	//	if err = file.Close(); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//
	//reader := bufio.NewReader(file)
	//fileInfo, _ := file.Stat()
	//bufferX := make([]byte, fileInfo.Size())
	//_, err = reader.Read(bufferX)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//bufferY := make([]byte, fileInfo.Size())
	//rand.Read(bufferY)
	//
	//bufferZ := make([]byte, fileInfo.Size())
	//
	//for i := 0; i < len(bufferX); i++ {
	//	bufferZ[i] = bufferX[i] ^ bufferY[i]
	//}
	//
	//for i := 0; i < len(bufferX); i++ {
	//	bufferZ[i] = bufferZ[i] ^ bufferY[i]
	//}
	//
	//
	//fileout, fileErr := os.Create("1" + filenameTestData)
	//if fileErr != nil {
	//	log.Fatal(fileErr)
	//	return
	//}
	//
	//defer func() {
	//	if err := fileout.Close(); err != nil {
	//		log.Fatal(err)
	//	}
	//}()
	//
	//_ = ioutil.WriteFile("1" + filenameTestData, bufferZ, 0644)

	producerName := "Alice"
	consumerName := "Bob"

	cryptosystem := RSA.CryptoSystem{}
	cryptosystem.Init()
	_ = cryptosystem.AddUser(producerName)
	_ = cryptosystem.AddUser(consumerName)

	filenameTestData := "test_data.png"

	cryptosystem.SendMessageFromFile(producerName, consumerName, filenameTestData)

	//producer := cryptosystem.Users[producerName]
	//consumer := cryptosystem.Users[consumerName]

	//FileWrapper.WriteToFile("encrypt-" + filenameTestData, producer.GetMessage())
	//FileWrapper.WriteToFile("decrypt-" + filenameTestData, consumer.GetMessage())
	cryptosystem.PrintUsers()

	if cryptosystem.CheckMessage(producerName, consumerName) != false {
		fmt.Println("OK")
	}
}

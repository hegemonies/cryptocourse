package FileWrapper

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
	"os"
)

// if won't work this
// then try this: i = binary.LittleEndian.Uint64(bytes)
func GetMessageFromFileByP(filename string, p uint64) []uint64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	reader := bufio.NewReader(file)
	countBytes := getMaxIndexOfBit(p) / 8
	if countBytes == 0 {
		countBytes++
	}
	buffer := make([]byte, countBytes)

	fileInfo, _ := file.Stat()
	countMessages := fileInfo.Size() / int64(countBytes)
	message := make([]uint64, 0 , countMessages)


	for i := 0; ; i++ {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		var tmpChunk uint64 = 0
		var j uint64
		for j = 0; j < uint64(n); j++ {
			tmpChunk = (tmpChunk << (j * 8)) | uint64(buffer[j])
		}
		message = append(message, tmpChunk)
	}

	return message
}

func GetMessageFromFile(filename string) []uint64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	reader := bufio.NewReader(file)
	var countBytes int64 = 1
	buffer := make([]byte, countBytes)

	fileInfo, _ := file.Stat()
	countMessages := fileInfo.Size() / countBytes
	message := make([]uint64, 0, countMessages)


	for i := 0; ; i++ {
		_, err := reader.Read(buffer)
		if err != nil {
			break
		}

		var chunk uint64 = 0
		chunk = (chunk << 8) | uint64(buffer[0])
		message = append(message, chunk)
	}

	return message
}

func getMaxIndexOfBit(number uint64) (index int) {
	var i uint64
	for i = 0; i < 64; i++ {
		if (number >> i) & 1 == 1 {
			index = int(i)
		}
	}
	return
}

func WriteToFileByP(filename string, message []uint64, p uint64) {
	file, fileErr := os.Create(filename)
	if fileErr != nil {
		log.Fatal(fileErr)
		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	buf := new(bytes.Buffer)

	countBytes := getMaxIndexOfBit(p) / 8
	if countBytes == 0 {
		countBytes++
	}

	for i := 0; i < len(message); i++ {
		chunk := convertFromUint64ToByte(message[i], countBytes)
		err := binary.Write(buf, binary.LittleEndian, chunk)
		if err != nil {
			log.Fatal(err)
		}
	}

	err1 := ioutil.WriteFile(filename, buf.Bytes(), 0644)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func WriteToFile(filename string, message []uint64) {
	file, fileErr := os.Create(filename)
	if fileErr != nil {
		log.Fatal(fileErr)
		return
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	buf := new(bytes.Buffer)

	countBytes := 1

	for i := 0; i < len(message); i++ {
		chunk := convertFromUint64ToByte(message[i], countBytes)
		err := binary.Write(buf, binary.LittleEndian, chunk)
		if err != nil {
			log.Fatal(err)
		}
	}

	err1 := ioutil.WriteFile(filename, buf.Bytes(), 0644)
	if err1 != nil {
		log.Fatal(err1)
	}
}

func convertFromUint64ToByte(from uint64, countBytes int) (to []byte) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, from)
	to = make([]byte, countBytes)
	copy(to, buf)
	return
}

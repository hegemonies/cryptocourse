package FileWrapper

import (
	"bufio"
	"log"
	"os"
)

func GetMessageFromFileByP(filename string, p uint64) []uint64 {
	message := make([]uint64, 0 , 0)

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
	countBytes := (getMaxIndexOfBit(p) / 8) + 1
	buffer := make([]byte, 0, countBytes)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
	}

	return  message
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

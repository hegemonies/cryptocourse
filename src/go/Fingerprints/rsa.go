package Fingerprints

import (
	"bufio"
	"crypto/md5"
	diffiehellman "cryptocrouse/src/go/Diffie-Hellman"
	"cryptocrouse/src/go/EuclideanAlgorithm"
	"cryptocrouse/src/go/FastExp"
	"cryptocrouse/src/go/FileWrapper"
	"encoding/binary"
	"fmt"
	"hash"
	"io"
	"log"
	"math/big"
	"math/rand"
	"os"
)

type User struct {
	Name       string
	c          uint64
	D          uint64
	p          uint64
	q          uint64
	N          uint64
	phi        uint64
	hash       hash.Hash
	y          []byte
	s          []byte
}

const (
	MaxQ = 50000
	MinP = 257
	MaxP = 100000
	MaxBound = 268435456
)

func (user *User) generateP() {
	for {
		user.generateQ()
		user.p = 2 * user.q + 1

		if diffiehellman.IsPrimeGoogle(user.p) && diffiehellman.IsPrimeGoogle(user.q) {
			if user.p >= MinP && user.p <= MaxP {
				break
			}
		}
	}
}

func (user *User) generateQ() {
	user.q = diffiehellman.GeneratePrimeNumberWithBounds(MaxQ)
}

func (user *User) generatePhi() {
	user.phi = (user.p - 1) * (user.q - 1)
}

func (user *User) generateN() {
	user.N = user.p * user.q
}

func (user *User) generateC() {
	_, _, y := EuclideanAlgorithm.GCD(int64(user.D), int64(user.phi))
	if y < 0 {
		user.c = uint64(int64(user.p) + y)
	} else {
		user.c = uint64(y)
	}
}

func (user *User) generateD() {
	for {
		user.D = rand.Uint64() % MaxBound

		if user.D >= user.phi {
			continue
		}

		GCD, _, _ := EuclideanAlgorithm.GCD(int64(user.phi), int64(user.D))
		if GCD == 1 {
			break
		}
	}
}

func (user *User) GeneratePrivateVariables() {
	user.generateP()
	user.generateN()
	user.generatePhi()

	for {
		user.generateD()
		user.generateC()

		if ((user.c * user.D) % (user.phi)) == 1 {
			break
		}
	}
}

func (user *User) PrintUserInfo(format string) {
	fmt.Printf(format, user.Name, user.c, user.D, user.p, user.q, user.N, user.phi, user.y, user.s, user.hash)
}

func (user *User) PrintOpenKeysInFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		_ = fmt.Errorf("%v\n", err)
	}
	defer file.Close()

	_, _ = fmt.Fprintf(file, "N=%d\nD=%d", user.N, user.D)
}

func (user *User) ComputeHash(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		_ = fmt.Errorf("%v\n", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	user.hash = md5.New()

	if _, err := io.Copy(user.hash, reader); err != nil {
		log.Fatal(err)
	}

	//fileBytes := FileWrapper.GetMessageFromFileInBytes(filename)

	user.y = user.hash.Sum(nil)
}

func (user *User) ComputeSignature() {
	user.s = make([]byte, 0, len(user.y))
	for i := 0; i < len(user.y); i++ {
		//buf := FileWrapper.ConvertFromUint64ToByte(FastExp.SmallFastExp(uint64(user.y[i]), user.c, user.N), 1)[0]
		buf := big.NewInt(0).Exp(big.NewInt(int64(user.y[i])), big.NewInt(int64(user.c)), big.NewInt(int64(user.N)))
		user.s = append(user.s, buf.Bytes()[0])
	}
}

func (user *User) WriteHahSumToFile(filename string) {
	FileWrapper.WriteByteArrayToFile(filename, user.s)
}

func (user *User) CheckSignature() bool {
	for i := 0; i < len(user.s); i++ {
		w := FastExp.SmallFastExp(uint64(user.s[i]), user.D, user.N)
		//w := big.NewInt(0).Exp(big.NewInt(int64(user.s[i])), big.NewInt(int64(user.D)), big.NewInt(int64(user.N)))

		//if w.Bytes()[0] != user.y[i] {
		if byte(w) == user.y[i] {
			log.Println("Invalid Signature")
			return false
		}
	}

	log.Println("Signature is valid")
	return true
}

func (user *User) ConvertUint64ToByte(from uint64) (to byte) {
	bufFrom := make([]byte, 8)
	binary.LittleEndian.PutUint64(bufFrom, from)
	bufTo := make([]byte, 1)
	copy(bufTo, bufFrom)
	return bufTo[0]
}

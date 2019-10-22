package Fingerprints

import (
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
	s          []uint64
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
	for {
		user.generateP()
		user.generateN()
		user.generatePhi()
		user.generateD()
		user.generateC()

		if ((user.c * user.D) % (user.phi)) == 1 {
			break
		}
	}
}

func (user *User) PrintUserInfo(format string) {
	fmt.Printf(format, user.Name, user.c, user.D, user.p, user.q, user.N, user.phi)
}

func (user *User) PrintOpenKeysInFile(filename string) {
	file, err := os.Open(filename)
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

	user.hash = md5.New()

	if _, err := io.Copy(user.hash, file); err != nil {
		log.Fatal(err)
	}

	//fileBytes := FileWrapper.GetMessageFromFileInBytes(filename)

	user.hash.Sum(nil)
	user.convertHashToBytes()
}

func (user *User) convertHashToBytes() {
	user.y = make([]byte, user.hash.Size())
	user.hash.Write(user.y)
}

func (user *User) ComputeSignature() {
	user.s = make([]uint64, len(user.y))
	for i := 0; i < len(user.y); i++ {
		user.s[i] = FastExp.SmallFastExp(uint64(user.y[i]), user.c, user.N)
	}
}

func (user *User) WriteHahSumToFile(filename string) {
	FileWrapper.WriteToFile(filename, user.s)
}

func (user *User) CheckSignature() bool {
	for i := 0; i < len(user.s); i++ {
		w := FastExp.SmallFastExp(user.s[i], user.D, user.N)

		if w != user.s[i] {
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

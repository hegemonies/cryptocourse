package Fingerprints

import (
	"bufio"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
)

type ElGamalUser struct {
	P       big.Int
	G       big.Int
	x       big.Int
	Y       big.Int
	Message []byte
	h       big.Int
	k       big.Int
	R       big.Int
	u       big.Int
	S       big.Int
	q       big.Int
}

var (
	MaxRandom = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(130), nil)
	MinBound = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(128), nil)
)

func (user *ElGamalUser) ComputeHash(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		_ = fmt.Errorf("%v\n", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	hash := md5.New()

	if _, err := io.Copy(hash, reader); err != nil {
		log.Fatal(err)
	}

	checksum := hash.Sum(nil)
	hexstr := hex.EncodeToString(checksum)
	user.h.SetString(hexstr, 16)
}

func (user *ElGamalUser) ComputeSignature() {
	user.generateK()
	user.generateR()
	user.generateU()
	user.generateS()
}

func (user *ElGamalUser) generateS() {
	user.S = *big.NewInt(0).Mod(
		big.NewInt(0).Mul(
			Inversion(&user.k, big.NewInt(0).Sub(
				&user.P,
				big.NewInt(1))),
			&user.u),
		big.NewInt(0).Sub(
			&user.P,
			big.NewInt(1)))
}

func Inversion(a, b *big.Int) *big.Int {
	x := big.NewInt(0)
	big.NewInt(0).GCD(x, nil, a, b)
	return x
}

func (user *ElGamalUser) generateU() {
	user.u = *big.NewInt(0).Mod(
		big.NewInt(0).Sub(&user.h, big.NewInt(0).Mul(&user.x, &user.R)),
		big.NewInt(0).Sub(&user.P, big.NewInt(1)))
}

func (user *ElGamalUser) generateR() {
	user.R = BigFastExp(user.G, user.k, user.P)
}

func (user *ElGamalUser) generateK() {
	for {
		user.k = *GetBigRandom()
		GCD := big.NewInt(0).GCD(
			nil,
			nil,
			&user.k,
			big.NewInt(0).Sub(
				&user.P,
				big.NewInt(1)))
		if GCD.Cmp(big.NewInt(1)) == 0 {
			break
		}
		//if GCD.Cmp(big.NewInt(1)) == 0 {
		//	inversionK := Inversion(&user.k, PSumOne(&user.P))
		//	tmp := big.NewInt(0).Mod(
		//		big.NewInt(0).Mul(
		//			inversionK,
		//			&user.k),
		//		PSumOne(&user.P))
		//	if tmp.Cmp(big.NewInt(1)) == 0 {
		//		break
		//	}
		//}
		//inversionK := Inversion(&user.k, PSumOne(&user.P))
		//tmp := big.NewInt(0).Mod(
		//	big.NewInt(0).Mul(
		//		inversionK,
		//		&user.k),
		//	PSumOne(&user.P))
		//if tmp.Cmp(big.NewInt(1)) == 0 {
		//	break
		//}
	}
}

func PSumOne(p *big.Int) *big.Int {
	return big.NewInt(0).Sub(
		p,
		big.NewInt(1))
}

func (user *ElGamalUser) GeneratePrivateNumbers() {
	user.generateP()
	user.generateG()
	user.generateX()
	user.generateY()
}

func (user *ElGamalUser) generateY() {
	user.Y = BigFastExp(user.G, user.x, user.P)
}

func (user *ElGamalUser) generateX() {
	user.x = *big.NewInt(0).Exp(
		big.NewInt(2),
		big.NewInt(128),
		&user.P)
}

func (user *ElGamalUser) generateQ() {
	user.q = GenerateBigPrimeNumber()
}

func (user *ElGamalUser) generateP() {
	user.P = *big.NewInt(0)

	for {
		user.generateQ()
		user.P.Add(
			big.NewInt(0).Mul(
				big.NewInt(2),
				&user.q),
			big.NewInt(1))
		if IsPrime(user.P) {
			break
		}
	}
}

func (user *ElGamalUser) generateG() {
	for {
		user.G = *GetBigRandom()
		tmp := BigFastExp(user.G, user.q, user.P)
		if tmp.Cmp(big.NewInt(1)) == 0 {
			break
		}
	}
}

func GenerateBigPrimeNumber() big.Int {
	for {
		tmp := GetBigRandom()
		if IsPrime(*tmp) {
			return *tmp
		}
	}
}

func IsPrime(number big.Int) bool {
	return number.ProbablyPrime(20)
}

func GetBigRandom() *big.Int {
	n, err := rand.Int(rand.Reader, MaxRandom)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func BigFastExp(a, x, p big.Int) big.Int {
	return *big.NewInt(0).Exp(&a, &x, &p)
}

func (user *ElGamalUser) CheckMessage(filename string, y, r, s, g, p *big.Int) bool {
	user.ComputeHash(filename)
	fmt.Printf("hash = %s\n", user.h.Text(10))

	t1 := big.NewInt(0).Exp(y, r, p)
	fmt.Printf("t1 = %s\n", t1.Text(10))
	t2 := big.NewInt(0).Exp(r, s, p)
	fmt.Printf("t2 = %s\n", t2.Text(10))
	lValue := big.NewInt(0).Mod(
		big.NewInt(0).Mul(
			t1,
			t2),
		p)
	fmt.Printf("lvalue = %s\n", lValue.Text(10))
	rvalue := big.NewInt(0).Exp(g, &user.h, p)
	fmt.Printf("rvalue = %s\n", rvalue.Text(10))

	if lValue.Cmp(rvalue) != 0 {
		return false
	}

	return true
}

func (user *ElGamalUser) CheckHash() {
	for {
		if user.h.Cmp(&user.P) < 0 {
			break
		}
		user.GeneratePrivateNumbers()
	}
}

func (user *ElGamalUser) PrintInfo() {
	fmt.Printf("p = %s\ng = %s\nx = %s\ny = %s\nh = %s\nk = %s\nr = %s\nu = %s\ns = %s\nq = %s\n",
		user.P.Text(10),
		user.G.Text(10),
		user.x.Text(10),
		user.Y.Text(10),
		user.h.Text(10),
		user.k.Text(10),
		user.R.Text(10),
		user.u.Text(10),
		user.S.Text(10),
		user.q.Text(10))
}

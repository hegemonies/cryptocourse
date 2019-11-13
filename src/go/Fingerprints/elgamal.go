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
	"strings"
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
	MinBound = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(130), nil)
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

func GenerateBigPrimeNumberRef() *big.Int {
	for {
		tmp := GetBigRandom()
		if IsPrime(*tmp) {
			return tmp
		}
	}
}

func GenerateBigPrimeNumberWithLimit(limit *big.Int) *big.Int {
	for {
		tmp := GetBigRandomWithLimit(limit)
		if IsBigPrime(tmp) {
			return tmp
		}
	}
}

func IsPrime(number big.Int) bool {
	return number.ProbablyPrime(20)
}

func IsPrimeRef(number *big.Int) bool {
	return number.ProbablyPrime(20)
}

func IsBigPrime(number *big.Int) bool {
	return number.ProbablyPrime(20)
}

func GetBigRandom() *big.Int {
	n, err := rand.Int(rand.Reader, MaxRandom)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

func GetBigRandomWithLimit(limit *big.Int) *big.Int {
	n, err := rand.Int(rand.Reader, limit)
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

	lValue := big.NewInt(0).Mod(
		big.NewInt(0).Mul(
			big.NewInt(0).Exp(y, r, p),
			big.NewInt(0).Exp(r, s, p)),
		p)
	rvalue := big.NewInt(0).Exp(g, &user.h, p)

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

func (user *ElGamalUser) PrintOpenKeysToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		_ = fmt.Errorf("%v\n", err)
	}
	defer file.Close()

	_, _ = fmt.Fprintf(file,
		"P=%s\nG=%s\nY=%s",
		user.P.Text(10),
		user.G.Text(10),
		user.Y.Text(10))
}

func (user *ElGamalUser) PrintSignatureToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		_ = fmt.Errorf("%v\n", err)
	}
	defer file.Close()

	_, _ = fmt.Fprintf(file, "R=%s\nS=%s", user.R.Text(10), user.S.Text(10))
}

func ElGamalGetSignatureFromFile(filename string) (R, S *big.Int) {
	file, err := os.Open(filename)
	if err != nil {
		_ = fmt.Errorf("%v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if lines != nil {
		R, _ = big.NewInt(0).SetString(strings.Split(lines[0], "=")[1], 10)
		S, _ = big.NewInt(0).SetString(strings.Split(lines[1], "=")[1], 10)
	}

	return
}

func ElGamalGetOpenKeysFromFile(filename string) (P, G, Y *big.Int) {
	file, err := os.Open(filename)
	if err != nil {
		_ = fmt.Errorf("%v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if lines != nil {
		P, _ = big.NewInt(0).SetString(strings.Split(lines[0], "=")[1], 10)
		G, _ = big.NewInt(0).SetString(strings.Split(lines[1], "=")[1], 10)
		Y, _ = big.NewInt(0).SetString(strings.Split(lines[2], "=")[1], 10)
	}

	return
}

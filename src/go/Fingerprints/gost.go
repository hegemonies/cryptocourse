package Fingerprints

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"strings"
)

type GostUser struct {
	P *big.Int
	Q *big.Int
	b *big.Int
	A *big.Int
	x *big.Int
	Y *big.Int
	h *big.Int
	k *big.Int
	R *big.Int
	S *big.Int
}

var (
	MaxLimitP = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(1024), nil)
	MaxLimitB = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(768), nil)
	MaxLimitQ = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(256), nil)
)

func (user *GostUser) GenerateKeys() {
	user.generateP()
	user.generateA()
	user.generateX()
	user.generateY()
}

func (user *GostUser) GenerateSignature(filename string) {
	user.ComputeHash(filename)
	user.generateS()
}

func (user *GostUser) CheckSignature(filename string, r, s, q, a, p, y *big.Int) bool {
	user.ComputeHash(filename)

	if (r.Cmp(q) >= 0 && r.Cmp(big.NewInt(0)) <= 0) ||
		(s.Cmp(q) >= 0 && s.Cmp(big.NewInt(0)) <= 0) {
		return false
	}

	u1 := user.computeU1(s, q)
	u2 := user.computeU2(r, q)
	v := user.computeV(a, u1, u2, p, q, y)

	if v.Cmp(r) != 0 {
		return false
	}

	return true
}

func (user *GostUser) computeU1(s, q *big.Int) *big.Int {
	return big.NewInt(0).Mod(
			big.NewInt(0).Mul(
				s,
				Inversion(user.h, q)),
			q)
}

func (user *GostUser) computeU2(r, q *big.Int) *big.Int {
	return big.NewInt(0).Mod( // todo: mb without sub
			big.NewInt(0).Mul(
				big.NewInt(0).Sub(
					q,
					r),
				Inversion(user.h, q)),
			q)
}

func (user *GostUser) computeV(a, u1, u2, p, q, y *big.Int) *big.Int {
	return big.NewInt(0).Mod(
			big.NewInt(0).Mod(
				big.NewInt(0).Mul(
					big.NewInt(0).Exp(
						a,
						u1,
						p),
					big.NewInt(0).Exp(
						y,
						u2,
						p)),
				p),
			q)
}

func (user *GostUser) generateP() {
	user.P = big.NewInt(0)

	for {
		user.generateQ()
		user.generateB()

		user.P.Add(
			big.NewInt(0).Mul(
				user.b,
				user.Q),
			big.NewInt(1))

		if IsBigPrime(user.P) {
			break
		}
	}
}

func (user *GostUser) generateQ() {
	user.Q = GenerateBigPrimeNumberWithLimit(MaxLimitQ)
}

func (user *GostUser) generateB() {
	user.b = GetBigRandomWithLimit(MaxLimitB)
}

func (user *GostUser) generateA() {
	//for {
	//	user.A = GetBigRandomWithLimit(MaxLimitP)
	//
	//	tmp := big.NewInt(0).Exp(user.A, user.Q, user.P)
	//
	//	if tmp.Cmp(big.NewInt(1)) == 0 {
	//		break
	//	}
	//}

	for {
		limit := big.NewInt(0).Sub(MaxLimitP, big.NewInt(1))
		g := GenerateBigPrimeNumberWithLimit(limit)
		user.A = big.NewInt(0).Exp(g, user.b, user.P)
		if user.A.Cmp(big.NewInt(1)) > 0 {
			break
		}
	}
}

func (user *GostUser) generateX() {
	user.x = GetBigRandomWithLimit(MaxLimitP)
}

func (user *GostUser) generateY() {
	user.Y = big.NewInt(0).Exp(user.A, user.x, user.P)
}

func (user *GostUser) ComputeHash(filename string) {
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
	user.h = big.NewInt(0)
	user.h.SetString(hexstr, 16)
}

func (user *GostUser) generateK() {
	user.k = GetBigRandomWithLimit(MaxLimitQ)
}

func (user *GostUser) generateR() {
	for {
		user.generateK()
		user.R = big.NewInt(0).Mod(
			big.NewInt(0).Exp(
				user.A,
				user.k,
				user.P),
			user.Q)

		if user.R.Cmp(big.NewInt(0)) != 0 {
			break
		}
	}
}

func (user *GostUser) generateS() {
	for {
		user.generateR()
		user.S = big.NewInt(0).Mod(
			big.NewInt(0).Add(
				big.NewInt(0).Mul(
					user.x,
					user.R),
				big.NewInt(0).Mul(
					user.k,
					user.h)),
			user.Q)

		if user.S.Cmp(big.NewInt(0)) != 0 {
			break
		}
	}
}

func (user *GostUser) PrintOpenKeysToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		_ = fmt.Errorf("%v\n", err)
	}
	defer file.Close()

	_, _ = fmt.Fprintf(file,
		"P=%s\nQ=%s\nA=%s\nY=%s",
		user.P.Text(10),
		user.Q.Text(10),
		user.A.Text(10),
		user.Y.Text(10))
}

func GostGetOpenKeysFromFile(filename string) (P, Q, A, Y *big.Int) {
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
		Q, _ = big.NewInt(0).SetString(strings.Split(lines[1], "=")[1], 10)
		A, _ = big.NewInt(0).SetString(strings.Split(lines[2], "=")[1], 10)
		Y, _ = big.NewInt(0).SetString(strings.Split(lines[3], "=")[1], 10)
	}

	return
}

func (user *GostUser) PrintSignatureToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		_ = fmt.Errorf("%v\n", err)
	}
	defer file.Close()

	_, _ = fmt.Fprintf(file, "R=%s\nS=%s", user.R.Text(10), user.S.Text(10))
}

func GostGetSignatureFromFile(filename string) (R, S *big.Int) {
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

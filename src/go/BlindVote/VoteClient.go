package BlindVote

import (
	"crypto/sha1"
	"cryptocrouse/src/go/Fingerprints"
	"encoding/hex"
	"math/big"
)

type VoteClient struct {
	Name string
	rnd  *big.Int
	v    *big.Int
	N    string
	r    *big.Int
	s    *big.Int
	h    *big.Int
	H2   *big.Int
}

func (client *VoteClient) SetV(v *big.Int) {
	client.v = v
	client.generateRnd()
	client.generateN()
}

func (client *VoteClient) generateRnd() {
	client.rnd = big.NewInt(0).Exp(big.NewInt(2), big.NewInt(512), nil)
}

func (client *VoteClient) generateN() {
	client.N = client.rnd.Text(10) + client.v.Text(10)
}

func (client *VoteClient) GenerateR(N *big.Int) {
	for {
		client.r = Fingerprints.GetBigRandomWithLimit(MaxP)
		GCD := big.NewInt(0).GCD(
			nil,
			nil,
			client.r,
			N)
		if GCD.Cmp(big.NewInt(1)) == 0 {
			break
		}
	}
}

func (client *VoteClient) ComputeHash() {
	hasher := sha1.New()
	hasher.Write([]byte(client.N))
	checksum := hasher.Sum(nil)
	hexstr := hex.EncodeToString(checksum)
	client.h, _ = big.NewInt(0).SetString(hexstr, 16)
}

func (client *VoteClient) ComputeHash2(d, N *big.Int) {
	//client.H2 = big.NewInt(0).Mod(
	//	big.NewInt(0).Mul(
	//		client.h,
	//		big.NewInt(0).Exp(client.r, d, nil)),
	//	N)
	client.H2 = big.NewInt(0).Mod(
		big.NewInt(0).Mul(
			big.NewInt(0).Exp(client.h, big.NewInt(1), N),
			big.NewInt(0).Exp(client.r, d, N)),
		N)
}

func (client *VoteClient) ComputeS(s2 *big.Int) {
	n, _ := big.NewInt(0).SetString(client.N, 10)
	client.s = big.NewInt(0).Mod(
		big.NewInt(0).Mul(
			s2,
			Fingerprints.Inversion(client.r, n)),
		n)
}

func (client *VoteClient) GetNewsletter() Newsletter {
	n, _ := big.NewInt(0).SetString(client.N, 10)
	return Newsletter{
		N: n,
		S: client.s,
	}
}

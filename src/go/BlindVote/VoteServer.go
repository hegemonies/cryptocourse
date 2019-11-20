package BlindVote

import (
	"crypto/sha1"
	"cryptocrouse/src/go/Fingerprints"
	"encoding/hex"
	"errors"
	"math/big"
)

type VoteServer struct {
	p   *big.Int
	q   *big.Int
	N   *big.Int
	c   *big.Int
	D   *big.Int
	phi *big.Int
	Voted map[string]bool
	newsletters map[string]Newsletter
	GiveOutsNewsletter map[string]bool
	CorrectNewsletter []Newsletter
}

func InitVoteServer() *VoteServer {
	server := &VoteServer{}
	server.Voted = make(map[string]bool)
	server.newsletters = make(map[string]Newsletter)
	server.GiveOutsNewsletter = make(map[string]bool)
	server.CorrectNewsletter = make([]Newsletter, 0)

	return server
}

func (server *VoteServer) GenerateNumbers() {
	for {
		server.generateP()
		server.generateQ()
		if server.p.Cmp(server.q) != 0 {
			break
		}
	}

	server.generatePhi()
	server.generateN()

	for {
		server.generateC()
		server.generateD()
		if big.NewInt(0).Mod(
			big.NewInt(0).Mul(server.c, server.D),
			server.phi).
			Cmp(big.NewInt(1)) == 0 {
			break
		}
	}
}

func (server *VoteServer) generateP() {
	server.p = Fingerprints.GenerateBigPrimeNumberWithLimit(MaxP)
}

func (server *VoteServer) generateQ() {
	server.q = Fingerprints.GenerateBigPrimeNumberWithLimit(MaxP)
}

func (server *VoteServer) generateN() {
	server.N = big.NewInt(0).Mul(server.p, server.q)
}

func (server *VoteServer) generatePhi() {
	server.phi = big.NewInt(0).Mul(
		big.NewInt(0).Sub(
			server.p,
			big.NewInt(1)),
		big.NewInt(0).Sub(
			server.q,
			big.NewInt(1)))
}

func (server *VoteServer) generateC() {
	for {
		server.c = Fingerprints.GetBigRandomWithLimit(MaxP)
		GCD := big.NewInt(0).GCD(
			nil,
			nil,
			server.c,
			big.NewInt(0).Sub(server.p, big.NewInt(1)))
		if GCD.Cmp(big.NewInt(1)) == 0 {
			break
		}
	}
}

func (server *VoteServer) generateD() {
	server.D = Fingerprints.Inversion(server.c, server.phi)
}

func (server *VoteServer) ComputeS2(h2 *big.Int) *big.Int {
	return big.NewInt(0).Exp(h2, server.c, server.N)
}

func (server *VoteServer) CommitVote(name string)  {
	server.Voted[name] = true
}

func (server *VoteServer) AddNewsletter(name string, newsletter Newsletter) {
	server.newsletters[name] = newsletter
}

func (server *VoteServer) CheckCorrectNewsletter(name string) (err error) {
	if newsletter, ok := server.newsletters[name]; ok {
		lvalue := ComputeHash(newsletter.N.Text(10))
		rvalue := big.NewInt(0).Exp(
			newsletter.S,
			server.D,
			server.N)

		if lvalue.Cmp(rvalue) != 0 {
			return errors.New("Error check newsletter: not correct ")
		}

		server.CommitVote(name)
		server.CorrectNewsletter = append(server.CorrectNewsletter, newsletter)
	}

	return nil
}

func ComputeHash(str string) *big.Int {
	h := sha1.New()
	h.Write([]byte(str))
	checksum := hex.EncodeToString(h.Sum(nil))
	hash, _ := big.NewInt(0).SetString(checksum, 10)
	return hash
}

func (server *VoteServer) GiveOutNewsletterTo(name string) {
	server.GiveOutsNewsletter[name] = true
}

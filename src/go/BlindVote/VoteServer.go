package BlindVote

import (
	"cryptocrouse/src/go/Fingerprints"
	"math/big"
)

type VoteServer struct {
	p   *big.Int
	q   *big.Int
	N   *big.Int
	c   *big.Int
	D   *big.Int
	phi *big.Int
}

func InitVoteServer() *VoteServer {
	return &VoteServer{}
}

func (server *VoteServer) GenerateNumbers() {
	for {
		server.generateP()
		server.generateQ()
		if server.p.Cmp(server.q) != 0 {
			break
		}
	}

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



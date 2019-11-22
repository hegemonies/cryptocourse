package BlindVote

import (
	"errors"
	"fmt"
	"math/big"
)

type BlindVoter struct {
	voteServer         *VoteServer
	clientDb           *ClientDB
	correctNewsletters *CorrectNewsletters
}

func InitBlindVoter() *BlindVoter {
	bv := &BlindVoter{}
	bv.voteServer = InitVoteServer()
	bv.voteServer.GenerateNumbers()
	bv.clientDb = InitClientDB()
	bv.correctNewsletters = &CorrectNewsletters{}
	return bv
}

func (bv *BlindVoter) Vote(clientName string, candidate *big.Int) error {
	client, ok := bv.clientDb.RegisterClient(clientName)
	if !ok {
		return errors.New(fmt.Sprintf("Client with name - %s, already exists. Can not vote twice. ", clientName))
	}

	client.SetV(candidate)
	client.GenerateR(bv.voteServer.N)
	client.ComputeHash()
	client.ComputeHash2(bv.voteServer.D, bv.voteServer.N)

	client.ComputeS(bv.voteServer.ComputeS2(client.H2), bv.voteServer.N)
	newsletter := client.GetNewsletter()
	if !bv.voteServer.CheckCorrectNewsletter(newsletter) {
		return errors.New("Newsletter is not correct. ")
	}

	bv.correctNewsletters.Add(newsletter)

	return nil
}

func (bv *BlindVoter) PrintCorrectNewsletters() {
	bv.correctNewsletters.PrintAll()
}

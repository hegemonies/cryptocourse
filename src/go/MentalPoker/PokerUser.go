package MentalPoker

import (
	"cryptocrouse/src/go/Fingerprints"
	"fmt"
	"math/big"
	"math/rand"
)

type PokerUser struct {
	name      string
	cards     [2]*big.Int
	countWins int
	money     int
	c         *big.Int
	d         *big.Int
}

func (user *PokerUser) shuffleDeck(deck map[int]*big.Int)  {
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

// generate private and public numbers
func (user *PokerUser) GenerateNumbers(p *big.Int) {
	user.generateC(p)
	user.generateD(p)
}

func (user *PokerUser) generateC(p *big.Int) {
	for {
		user.c = Fingerprints.GetBigRandomWithLimit(MaxBound)
		GCD := big.NewInt(0).GCD(
			nil,
			nil,
			user.c,
			big.NewInt(0).Sub(
				p,
				big.NewInt(1)))
		if GCD.Cmp(big.NewInt(1)) == 0 {
			break
		}
	}
}

func (user *PokerUser) generateD(p *big.Int) {
	user.d = Fingerprints.Inversion(user.c, big.NewInt(0).Sub(
		p,
		big.NewInt(1)))
}

func (user *PokerUser) encodeDeck(deck map[int]*big.Int, p *big.Int) {
	for i := 0; i < len(deck); i++ {
		deck[i] = big.NewInt(0).Exp(
			deck[i],
			user.c,
			p)
	}
}

func (user *PokerUser) get2Cards(deck map[int]*big.Int) {
	for i := 0; i < 2; i++ {
		k := rand.Int() % (len(deck) - 1)
		user.cards[i] = deck[k]
		fmt.Println(deck[k])
		fmt.Println(user.cards[i])
		delete(deck, k)
	}
}

func (user *PokerUser) decode2Cards(twoCards [2]*big.Int, p *big.Int) {
	for i := 0; i < 2; i++ {
		twoCards[i] = big.NewInt(0).Exp(twoCards[i], user.d, p)
	}
}

func (user *PokerUser) PrintInfo() {
	fmt.Printf("%s 1=%v 2=%v c=%v d=%v\n",
		user.name,
		user.cards[0],
		user.cards[1],
		user.c,
		user.d)
}

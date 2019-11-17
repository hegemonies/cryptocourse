package MentalPoker

import (
	"cryptocrouse/src/go/Fingerprints"
	"fmt"
	"math/big"
	"math/rand"
)

type PokerUser struct {
	name      string
	cards     [2]*Card
	countWins int
	money     int
	c         *big.Int
	d         *big.Int
}

func (user *PokerUser) shuffleDeck(deck map[int]*Card)  {
	rand.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
}

// generate private and public numbers
func (user *PokerUser) GenerateNumbers(p *big.Int) {
	for {
		user.generateC(p)
		user.generateD(p)
		if big.NewInt(0).Mod(
			big.NewInt(0).Mul(
				user.c,
				user.d),
			big.NewInt(0).Sub(
				p,
				big.NewInt(1))).Cmp(big.NewInt(1)) == 0 {
			break
		}
	}
}
// c= 19015 d=-36405 p=113819
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
	user.d = Fingerprints.Inversion(
		user.c,
		big.NewInt(0).Sub(p, big.NewInt(1)))
}

func (user *PokerUser) encodeDeck(deck map[int]*Card, p *big.Int) {
	for i := 0; i < len(deck); i++ {
		deck[i].Num = big.NewInt(0).Exp(
			deck[i].Num,
			user.c,
			p)
	}
}

func (user *PokerUser) get2Cards(deck map[int]*Card) {
	countGettedCards := 0
	for countGettedCards != 2 {
		k := rand.Int() % (len(deck) - 1)
		card := deck[k]
		if card != nil {
			user.cards[countGettedCards] = card
			countGettedCards++
			delete(deck, k)
		}
	}
}

func (user *PokerUser) decode2Cards(twoCards [2]*Card, p *big.Int) {
	for i := 0; i < 2; i++ {
		twoCards[i].Num = big.NewInt(0).Exp(twoCards[i].Num, user.d, p)
	}
}

func (user *PokerUser) PrintInfo() {
	fmt.Printf("Name=%6s 1=%12s 2=%12s c=%6v d=%6v\n",
		user.name,
		user.cards[0].ToString(),
		user.cards[1].ToString(),
		user.c,
		user.d)
}

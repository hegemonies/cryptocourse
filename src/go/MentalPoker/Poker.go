package MentalPoker

import (
	"cryptocrouse/src/go/Fingerprints"
	"errors"
	"fmt"
	"math/big"
)

type PokerSystem struct {
	Users    []*PokerUser
	Deck     map[int]*big.Int
	croupier *PokerUser
	P        *big.Int
}

func (system *PokerSystem) InitSystem() {
	system.croupier = new(PokerUser)
	system.croupier.name = "Croupier"
}

func RegistrationRandomUsers(count int) (system *PokerSystem, err error) {
	if count < MinCountUsers || count > MaxCountUsers {
		return nil, errors.New("Count user should be between " + string(MinCountUsers) + " and " + string(MaxCountUsers))
	}

	system = new(PokerSystem)
	system.GenerateP()
	system.GenerateDeck()
	system.createUsers(count)

	return system, nil
}

func (system *PokerSystem) generateQ() *big.Int {
	return Fingerprints.GenerateBigPrimeNumberWithLimit(MinP)
}

func (system *PokerSystem) GenerateP() {
	system.P = big.NewInt(0)

	for {
		q := system.generateQ()
		system.P.Add(
			big.NewInt(0).Mul(
				big.NewInt(2),
				q),
			big.NewInt(1))
		if Fingerprints.IsPrimeRef(system.P) {
			if system.P.Cmp(MinP) > 0 && system.P.Cmp(MaxP) < 0 {
				break
			}
		}
	}
}

func (system *PokerSystem) createUsers(count int) {
	system.Users = make([]*PokerUser, count)
	for i := 0; i < count; i++ {
		user := &PokerUser{}
		user.name = Fingerprints.GetBigRandom().Text(10)
		user.GenerateNumbers(system.P)
		user.cards[0] = big.NewInt(0)
		user.cards[1] = big.NewInt(0)
		//system.Users = append(system.Users, user)
		system.Users[i] = user
	}
}

func (system *PokerSystem) GenerateDeck() {
	system.Deck = make(map[int]*big.Int)
	for i := 0; i < CountCards; i++ {
		tmp := Fingerprints.GenerateBigPrimeNumberWithLimit(MaxBound)

		if contains(system.Deck, tmp) {
			i--
			continue
		}

		system.Deck[i] = tmp
	}
}

func contains(deck map[int]*big.Int, num *big.Int) bool {
	for i := 0; i < len(deck); i++ {
		if num == deck[i] {
			return true
		}
	}

	return false
}

func (system *PokerSystem) Round() {
	for i := 0; i < len(system.Users); i++ {
		system.ShuffleDeck(system.Users[i])
		system.EncodeDeck(system.Users[i])
	}
	for i := 0; i < len(system.Users); i++ {
		system.Getting2Cards(system.Users[i])
		system.DecodeDeck(i)
	}
}

// everybody shuffle deck
func (system *PokerSystem) ShuffleDeck(user *PokerUser) {
	user.shuffleDeck(system.Deck)
}

// everybody encode deck
func (system *PokerSystem) EncodeDeck(user *PokerUser) {
	user.encodeDeck(system.Deck, system.P)
}

// everybody get 2 cards
func (system *PokerSystem) Getting2Cards(user *PokerUser) {
	user.get2Cards(system.Deck)
}

// everybody decode deck
func (system *PokerSystem) DecodeDeck(userIndex int) {
	for i := 0; i < len(system.Users); i++ {
		system.Users[userIndex].decode2Cards(system.Users[i].cards, system.P)
	}
}

func (system *PokerSystem) PrintDeck() {
	fmt.Println("Deck:")
	for i := 0; i < len(system.Deck); i++ {
		fmt.Printf("[%d] %v", i, system.Deck[i])
		if i % 7 == 0 {
			fmt.Println()
		}
	}
}

func (system *PokerSystem) PrintUsersCards() {
	for i := 0; i < len(system.Users); i++ {
		system.Users[i].PrintInfo()
	}
}

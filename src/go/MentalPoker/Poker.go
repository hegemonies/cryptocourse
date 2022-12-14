package MentalPoker

import (
	"cryptocrouse/src/go/Fingerprints"
	"errors"
	"fmt"
	"math/big"
)

type PokerSystem struct {
	Users    []*PokerUser
	Deck     []*Card
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
		for {
			found := false
			name := Fingerprints.GetBigRandomWithLimit(MaxBound).Text(10)
			for j := 0; j < i; j++ {
				if system.Users[j].name == name {
					found = true
					break
				}
			}
			if found == false {
				user.name = name
				break
			}
		}
		user.GenerateNumbers(system.P)
		user.Cards[0], user.Cards[1] = &Card{}, &Card{}
		system.Users[i] = user
	}
}

func (system *PokerSystem) GenerateDeck() {
	system.Deck = make([]*Card, CountCards)
	for i := 0; i < CountCards; i++ {
		tmp := Fingerprints.GenerateBigPrimeNumberWithLimit(MaxBound)

		if contains(system.Deck, tmp) {
			i--
			continue
		}

		system.Deck[i] = GetCard(i)
		system.Deck[i].Num = tmp
	}
}

func contains(deck []*Card, num *big.Int) bool {
	for i := 0; i < len(deck); i++ {
		if deck[i] != nil {
			if num.Cmp(deck[i].Num) == 0 {
				return true
			}
		}
	}

	return false
}

func (system *PokerSystem) Round() {
	for i := 0; i < len(system.Users); i++ {
		system.ShuffleDeck(system.Users[i])
		system.EncodeDeck(system.Users[i])
	}
	fmt.Printf("Encoded ")
	system.PrintDeck()
	for i := 0; i < len(system.Users); i++ {
		system.Getting2Cards(system.Users[i])
	}
	for i := 0; i < len(system.Users); i++ {
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
	system.Deck = user.get2Cards(system.Deck)
}

// everybody decode deck
func (system *PokerSystem) DecodeDeck(userIndex int) {
	for i := 0; i < len(system.Users); i++ {
		system.Users[userIndex].decode2Cards(system.Users[i].Cards, system.P)
	}
}

func (system *PokerSystem) PrintDeck() {
	fmt.Println("Deck:")
	for i := 0; i < len(system.Deck); i++ {
		if card := system.Deck[i]; card != nil {
			fmt.Printf("[%2d] %12s ", i, system.Deck[i].ToString())
		} else {
			fmt.Printf("[%2d]              ", i)
		}
		if (i+1) % 4 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func (system *PokerSystem) PrintUsersCards() {
	for i := 0; i < len(system.Users); i++ {
		fmt.Printf("[%2d] ", i)
		system.Users[i].PrintInfo()
	}
}

func (system *PokerSystem) Copy() *PokerSystem {
	copy := &PokerSystem{
		Users:    nil,
		Deck:     nil,
		croupier: nil,
		P:        system.P,
	}

	//copy.Users = make([]*PokerUser, len(system.Users))
	//for i := 0; i < len(system.Users); i++ {
	//	copy.Users[i] = &PokerUser{}
	//	copy.Users[i].name = system.Users[i].name
	//	copy.Users[i].c = system.Users[i].c
	//	copy.Users[i].d = system.Users[i].d
	//	for j := 0; j < 2; j++ {
	//		copy.Users[i].Cards[j] = &Card{}
	//		copy.Users[i].Cards[j].Num = big.NewInt(0)
	//		copy.Users[i].Cards[j].Num.SetString(system.Users[i].Cards[j].Num.Text(10), 10)
	//		copy.Users[i].Cards[j].Type = &CardType{}
	//		copy.Users[i].Cards[j].Type.Num = system.Users[i].Cards[j].Type.Num
	//		copy.Users[i].Cards[j].Type.Type = system.Users[i].Cards[j].Type.Type
	//	}
	//	//copy.Users[i].Cards[0] = &Card{
	//	//	Num:  big.NewInt(0).SetUint64(system.Users[i].Cards[0].Num.Uint64()),
	//	//	Type: &CardType{
	//	//		Num:  system.Users[i].Cards[0].Type.Num,
	//	//		Type: system.Users[i].Cards[0].Type.Type,
	//	//	},
	//	//}
	//	//copy.Users[i].Cards[1] = &Card{}
	//	//copy.Users[i].Cards[1] = &Card{
	//	//	Num:  big.NewInt(0).SetUint64(system.Users[i].Cards[1].Num.Uint64()),
	//	//	Type: &CardType{
	//	//		Num:  system.Users[i].Cards[1].Type.Num,
	//	//		Type: system.Users[i].Cards[1].Type.Type,
	//	//	},
	//	//}
	//}
	
	copy.Deck = make([]*Card, CountCards)
	for i := 0; i < CountCards; i++ {
		copy.Deck[i] = &Card{}
		copy.Deck[i].Num, _ = big.NewInt(0).SetString(system.Deck[i].Num.Text(10), 10)
		copy.Deck[i].Type = &CardType{
			Num:  system.Deck[i].Type.Num,
			Type: system.Deck[i].Type.Type,
		}
	}
	
	return copy
}

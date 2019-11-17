package MentalPoker

import "math/big"

type Card struct {
	Num *big.Int
	Type *CardType
}

type CardType struct {
	Num int
	Type string
}

func GetCard(num int) *Card {
	switch num {
	case 0:
		return &Card{Num:big.NewInt(0), Type:&CardType{TWO, getHearts()}}
	case 1:
		return &Card{Num:big.NewInt(0), Type:&CardType{TWO, getClubs()}}
	case 2:
		return &Card{Num:big.NewInt(0), Type:&CardType{TWO, getDiamonds()}}
	case 3:
		return &Card{Num:big.NewInt(0), Type:&CardType{TWO, getSpades()}}

	case 4:
		return &Card{Num:big.NewInt(0), Type:&CardType{THREE, getHearts()}}
	case 5:
		return &Card{Num:big.NewInt(0), Type:&CardType{THREE, getClubs()}}
	case 6:
		return &Card{Num:big.NewInt(0), Type:&CardType{THREE, getDiamonds()}}
	case 7:
		return &Card{Num:big.NewInt(0), Type:&CardType{THREE, getSpades()}}

	case 8:
		return &Card{Num:big.NewInt(0), Type:&CardType{FOUR, getHearts()}}
	case 9:
		return &Card{Num:big.NewInt(0), Type:&CardType{FOUR, getClubs()}}
	case 10:
		return &Card{Num:big.NewInt(0), Type:&CardType{FOUR, getDiamonds()}}
	case 11:
		return &Card{Num:big.NewInt(0), Type:&CardType{FOUR, getSpades()}}

	case 12:
		return &Card{Num:big.NewInt(0), Type:&CardType{FIVE, getHearts()}}
	case 13:
		return &Card{Num:big.NewInt(0), Type:&CardType{FIVE, getClubs()}}
	case 14:
		return &Card{Num:big.NewInt(0), Type:&CardType{FIVE, getDiamonds()}}
	case 15:
		return &Card{Num:big.NewInt(0), Type:&CardType{FIVE, getSpades()}}

	case 16:
		return &Card{Num:big.NewInt(0), Type:&CardType{SIX, getHearts()}}
	case 17:
		return &Card{Num:big.NewInt(0), Type:&CardType{SIX, getClubs()}}
	case 18:
		return &Card{Num:big.NewInt(0), Type:&CardType{SIX, getDiamonds()}}
	case 19:
		return &Card{Num:big.NewInt(0), Type:&CardType{SIX, getSpades()}}

	case 20:
		return &Card{Num:big.NewInt(0), Type:&CardType{SEVEN, getHearts()}}
	case 21:
		return &Card{Num:big.NewInt(0), Type:&CardType{SEVEN, getClubs()}}
	case 22:
		return &Card{Num:big.NewInt(0), Type:&CardType{SEVEN, getDiamonds()}}
	case 23:
		return &Card{Num:big.NewInt(0), Type:&CardType{SEVEN, getSpades()}}

	case 24:
		return &Card{Num:big.NewInt(0), Type:&CardType{EIGHT, getHearts()}}
	case 25:
		return &Card{Num:big.NewInt(0), Type:&CardType{EIGHT, getClubs()}}
	case 26:
		return &Card{Num:big.NewInt(0), Type:&CardType{EIGHT, getDiamonds()}}
	case 27:
		return &Card{Num:big.NewInt(0), Type:&CardType{EIGHT, getSpades()}}

	case 28:
		return &Card{Num:big.NewInt(0), Type:&CardType{NINE, getHearts()}}
	case 29:
		return &Card{Num:big.NewInt(0), Type:&CardType{NINE, getClubs()}}
	case 30:
		return &Card{Num:big.NewInt(0), Type:&CardType{NINE, getDiamonds()}}
	case 31:
		return &Card{Num:big.NewInt(0), Type:&CardType{NINE, getSpades()}}

	case 32:
		return &Card{Num:big.NewInt(0), Type:&CardType{T, getHearts()}}
	case 33:
		return &Card{Num:big.NewInt(0), Type:&CardType{T, getClubs()}}
	case 34:
		return &Card{Num:big.NewInt(0), Type:&CardType{T, getDiamonds()}}
	case 35:
		return &Card{Num:big.NewInt(0), Type:&CardType{T, getSpades()}}

	case 36:
		return &Card{Num:big.NewInt(0), Type:&CardType{J, getHearts()}}
	case 37:
		return &Card{Num:big.NewInt(0), Type:&CardType{J, getClubs()}}
	case 38:
		return &Card{Num:big.NewInt(0), Type:&CardType{J, getDiamonds()}}
	case 39:
		return &Card{Num:big.NewInt(0), Type:&CardType{J, getSpades()}}

	case 40:
		return &Card{Num:big.NewInt(0), Type:&CardType{Q, getHearts()}}
	case 41:
		return &Card{Num:big.NewInt(0), Type:&CardType{Q, getClubs()}}
	case 42:
		return &Card{Num:big.NewInt(0), Type:&CardType{Q, getDiamonds()}}
	case 43:
		return &Card{Num:big.NewInt(0), Type:&CardType{Q, getSpades()}}

	case 44:
		return &Card{Num:big.NewInt(0), Type:&CardType{K, getHearts()}}
	case 45:
		return &Card{Num:big.NewInt(0), Type:&CardType{K, getClubs()}}
	case 46:
		return &Card{Num:big.NewInt(0), Type:&CardType{K, getDiamonds()}}
	case 47:
		return &Card{Num:big.NewInt(0), Type:&CardType{K, getSpades()}}

	case 48:
		return &Card{Num:big.NewInt(0), Type:&CardType{A, getHearts()}}
	case 49:
		return &Card{Num:big.NewInt(0), Type:&CardType{A, getClubs()}}
	case 50:
		return &Card{Num:big.NewInt(0), Type:&CardType{A, getDiamonds()}}
	case 51:
		return &Card{Num:big.NewInt(0), Type:&CardType{A, getSpades()}}

	default:
		return nil
	}
}

// Черви
func getHearts() string {
	return "♥️"
}

// Бубны
func getDiamonds() string {
	return "♦️"
}

// Пики
func getSpades() string {
	return "♠️"
}

// Трефы
func getClubs() string {
	return "♣️"
}

func (card *Card) ToString() string {
	return card.Num.Text(10) + " " + card.Type.ToString()
}

func (cardType *CardType) ToString() string {
	result := ""
	switch cardType.Num {
	case TWO:
		result += "2"
	case THREE:
		result += "3"
	case FOUR:
		result += "4"
	case FIVE:
		result += "5"
	case SIX:
		result += "6"
	case SEVEN:
		result += "7"
	case EIGHT:
		result += "8"
	case NINE:
		result += "9"
	case T:
		result += "10"
	case J:
		result += "J"
	case Q:
		result += "Q"
	case K:
		result += "K"
	case A:
		result += "A"
	}

	result += cardType.Type
	return result
}

func (card *Card) Compare(card2 *Card) bool {
	return card.Num.Cmp(card2.Num) == 0 && card.Type.Compare(card2.Type)
}

func (cardType *CardType) Compare(cardType2 *CardType) bool {
	return cardType.Num == cardType2.Num && cardType.Type == cardType2.Type
}

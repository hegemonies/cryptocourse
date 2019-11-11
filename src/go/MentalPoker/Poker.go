package MentalPoker

import "errors"

type PokerSystem struct {
	Users []PokerUser
	Deck [52]int
	croupier PokerUser
}

func RegistrationRandomUsers(count int) (system *PokerSystem, err error) {
	if count < MinCountUsers || count > MaxCountUsers {
		return nil, errors.New("Count user should be between " + string(MinCountUsers) + " and " + string(MaxCountUsers))
	}

	system = new(PokerSystem)

	for i := 0; i < count; i++ {

	}

	return system, nil
}

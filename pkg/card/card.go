package card

type Card struct {
	Id       int64
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
}

type Service struct {
	BankName string
	Cards    []*Card
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

func (s *Service) NewCard(card *Card) {
	card.Issuer = s.BankName
	s.Cards = append(s.Cards, card)
}

func (s *Service) FindCardByNumber(number string) (card *Card) {
	for _, card := range s.Cards {
		if number == card.Number {
			return card
		}
	}

	return nil
}

package transfer

import (
	"github.com/apkraft/bgo_homework_2_1/pkg/card"
)

type Fee struct {
	FeeInPercents  int64
	MinFeeInCopeks int64
}

type Service struct {
	CardSvc           *card.Service
	InsideTheBank     Fee
	ToAnotherBank     Fee
	BetweenOtherBanks Fee
}

func NewService(cardSvc *card.Service, insideTheBank Fee, toAnotherBank Fee, betweenOtherBanks Fee) *Service {
	return &Service{
		CardSvc:           cardSvc,
		InsideTheBank:     insideTheBank,
		ToAnotherBank:     toAnotherBank,
		BetweenOtherBanks: betweenOtherBanks,
	}
}

func (s *Service) Card2Card(fromCard, toCard string, amount int64) (totalWithdrawal int64, ok bool) {
	ok = true

	transferFromCard := s.CardSvc.FindCardByNumber(fromCard)
	transferToCard := s.CardSvc.FindCardByNumber(toCard)

	fee := s.calculateFee(transferFromCard, transferToCard)
	totalWithdrawal = s.totalWithdrawal(amount, fee)

	if transferFromCard == nil && transferToCard == nil {
		return totalWithdrawal, true
	}

	if transferFromCard == nil && transferToCard != nil {
		transferToCard.Balance += amount
		return totalWithdrawal, true
	}

	if transferFromCard != nil && transferToCard == nil && transferFromCard.Balance >= totalWithdrawal {
		transferFromCard.Balance -= totalWithdrawal
		return totalWithdrawal, true
	}

	if transferFromCard.Balance < totalWithdrawal {
		return totalWithdrawal, false
	}

	transferFromCard.Balance -= totalWithdrawal
	transferToCard.Balance += amount

	return totalWithdrawal, true
}

func (s *Service) calculateFee(fromCard, toCard *card.Card) *Fee {
	if fromCard != nil && toCard != nil {
		return &s.InsideTheBank
	} else if fromCard != nil && toCard == nil {
		return &s.ToAnotherBank
	} else {
		return &s.BetweenOtherBanks
	}
}

func (s *Service) totalWithdrawal(amount int64, fee *Fee) int64 {
	finalFee := amount * fee.FeeInPercents / 100

	if finalFee < fee.MinFeeInCopeks {
		finalFee = fee.MinFeeInCopeks
	}
	return amount + finalFee
}

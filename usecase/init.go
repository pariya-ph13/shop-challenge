package usecase

import (
	"github.com/shopspring/decimal"
	"shopChallenge/domain"
	third_domain "shopChallenge/thirdparty/domain"
)

type UseCaseImpl struct {
	Repo domain.Repo
	Sms  third_domain.SMS
}

type card struct {
	domain.Cards
	rules  domain.TransactionRules
	amount decimal.Decimal
}

type TransferAccountRequest struct {
	SourceAcc *domain.Accounts
	TargetAcc *domain.Accounts
	Amount    decimal.Decimal
}

type finalAmount decimal.Decimal

func NewUseCase(repo domain.Repo,
	sms third_domain.SMS) domain.Usecase {
	return &UseCaseImpl{Repo: repo, Sms: sms}
}

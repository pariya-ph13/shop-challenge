package usecase

import (
	"fmt"
	"shopChallenge/domain"
)

func (u UseCaseImpl) GetLatestTXNsOfMostActiveUsers() ([]domain.LatestTXn, error) {
	txns, err := u.Repo.GetLatestTXNsOfActiveUsers()
	if err != nil {
		return []domain.LatestTXn{}, err
	}

	var s []domain.LatestTXn
	var c []domain.Transactions
	for i, txn := range txns {
		c, err = u.readLatestTxnOnly(txn.CardID)
		if err != nil {
			return []domain.LatestTXn{}, err
		}
		s = append(s, domain.LatestTXn{
			Name:       c[0].Card.Account.Customer.Name,
			CustomerID: c[0].Card.Account.CustomerID,
		})
		for i2, _ := range c {
			c[i2].Card = nil
			c[i2].ToCard = nil
		}
		s[i].Transactions = append(s[i].Transactions, c...)
	}
	fmt.Println(c)
	return s, nil
}

func (u UseCaseImpl) readLatestTxnOnly(cardID int) ([]domain.Transactions, error) {
	ac := domain.CardAccess{
		Contained: true,
		AccountAccess: domain.AccountAccess{
			Contained: true,
			CustomerAccess: domain.CustomerAccess{
				Contained: true,
			},
		},
	}
	c, err := u.Repo.GetLatestTXNs(cardID, ac)
	if err != nil {
		return []domain.Transactions{}, err
	}
	return c, nil
}

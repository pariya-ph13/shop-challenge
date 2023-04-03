package repository

import (
	"fmt"
	"shopChallenge/domain"
)

func (r RepoImpl) ReadTransactionRules(
	transaction string) (domain.TransactionRules, error) {

	txnRule := domain.TransactionRules{Transaction: transaction}

	res := r.db.Where("transaction = ?", transaction).
		Find(&txnRule)
	fmt.Println("-----rules:", txnRule)
	if res.Error != nil {
		return domain.TransactionRules{}, domain.ErrReadTransactionRules
	}
	return txnRule, nil
}

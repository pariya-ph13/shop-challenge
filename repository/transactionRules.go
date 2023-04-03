package repository

import (
	log "github.com/sirupsen/logrus"
	"shopChallenge/domain"
)

func (r RepoImpl) ReadTransactionRules(
	transaction string) (domain.TransactionRules, error) {

	txnRule := domain.TransactionRules{Transaction: transaction}

	res := r.DB.Where("transaction = ?", transaction).
		Find(&txnRule)
	if res.Error != nil {
		log.Error("transactionRules:", res.Error)
		return domain.TransactionRules{}, domain.ErrReadTransactionRules
	}
	return txnRule, nil
}

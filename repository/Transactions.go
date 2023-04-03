package repository

import (
	"fmt"
	"gorm.io/gorm"
	"shopChallenge/domain"
)

func (r RepoImpl) InsertTransaction(
	transactions domain.Transactions) error {
	transactions.Card = &domain.Cards{}
	transactions.ToCard = &domain.Cards{}
	res := r.db.Create(&transactions)
	if res.Error != nil {
		return domain.ErrInsertTransaction
	}
	return nil
}

func (r RepoImpl) GetLatestTXNsOfActiveUsers() ([]domain.LatestInfo, error) {
	var result []domain.LatestInfo

	r.db.Raw(`SELECT card_id, Count FROM  ( SELECT card_id,COUNT( card_id) AS Count
	                         from transactions
	                         where current_timestamp > created_at + interval '10 minutes'
	                         GROUP BY card_id
	                       ) sub
				ORDER BY Count desc
				limit 3`).Find(&result)

	fmt.Println(result)
	return result, nil
}

func (r RepoImpl) GetLatestTXNs(cardNo int, ac domain.CardAccess) (
	[]domain.Transactions, error) {
	var txn []domain.Transactions
	db := r.db.Where("Transactions.card_id = ?", cardNo).Limit(10)
	r.readRelatedToTxn(db, ac)
	res := db.Find(&txn)
	if res.Error != nil {
		return []domain.Transactions{}, nil
	}
	return txn, res.Error
}
func (r RepoImpl) readRelatedToTxn(tx *gorm.DB, c domain.CardAccess) {
	if c.Contained {
		tx.Joins("Card")
		if c.AccountAccess.Contained {
			tx.Preload("Card.Account")
			if c.AccountAccess.CustomerAccess.Contained {
				tx.Preload("Card.Account.Customer")
			}
		}
	}
}

package repository

import (
	"gorm.io/gorm"
	"shopChallenge/domain"
)

func (r RepoImpl) ReadCard(
	cardNo int, c domain.AccountAccess) (domain.Cards, error) {

	card := domain.Cards{}

	db := r.DB.Where("card_id = ?", cardNo)
	r.readRelatedToCard(db, c)

	res := db.Find(&card)
	if res.Error != nil {
		return domain.Cards{}, domain.ErrReadCard
	}
	return card, nil
}

func (r RepoImpl) readRelatedToCard(tx *gorm.DB, c domain.AccountAccess) {
	if c.Contained {
		tx.Joins("Account")
		if c.CustomerAccess.Contained {
			tx.Preload("Account.Customer")
		}
	}
}

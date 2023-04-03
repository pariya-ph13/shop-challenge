package repository

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"shopChallenge/domain"
)

func (r *RepoImpl) StartTransaction() error {
	r.DB.Begin()
	if r.DB.Error != nil {
		log.Error("Start txn err:", r.DB.Error)
		return domain.ErrNoStartTxn
	}
	return nil
}

func (r *RepoImpl) FinalizeTransaction(err error) error {
	if err == nil {
		r.DB.Commit()
		return r.DB.Error
	}
	return r.DB.Error
}

func (r RepoImpl) GetTxn() *gorm.DB {
	return r.DB
}

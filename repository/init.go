package repository

import (
	"gorm.io/gorm"
	"shopChallenge/domain"
)

type RepoImpl struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) domain.Repo {
	return &RepoImpl{db: db}
}

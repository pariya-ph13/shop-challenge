package repository

import (
	"gorm.io/gorm"
	"shopChallenge/domain"
)

type RepoImpl struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) domain.Repo {
	return &RepoImpl{DB: db}
}

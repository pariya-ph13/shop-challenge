package domain

import (
	"github.com/shopspring/decimal"
	"time"
)

type Customers struct {
	ID           uint64    `json:"id"`
	CustomerID   int       `json:"customerId" gorm:"index"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	MobileNumber string    `json:"mobileNumber"`
	IsActive     bool      `json:"isActive"`
}

func (Customers) TableName() string {
	return "customers"
}

type Accounts struct {
	ID         uint64          `json:"id"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
	AccountID  int             `json:"accountId"`
	Balance    decimal.Decimal `json:"balance"`
	CustomerID int             `json:"customerId"`
	Customer   Customers       `gorm:"references:CustomerID"`
	IsActive   bool            `json:"isActive"`
}

func (Accounts) TableName() string {
	return "accounts"
}

type Cards struct {
	ID        uint64    `json:"id,omitempty"`
	CardID    int       `json:"cardId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	AccountID int       `json:"accountId,omitempty"`
	Account   Accounts  `gorm:"references:AccountID" json:"account,omitempty"`
	IsActive  bool      `json:"isActive,omitempty"`
}

func (Cards) TableName() string {
	return "cards"
}

type TransactionRules struct {
	ID          uint64          `json:"id"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	Transaction string          `json:"transaction"`
	Fee         decimal.Decimal `json:"fee"`
	MaxLimit    decimal.Decimal `json:"maxLimit"`
	MinLimit    decimal.Decimal `json:"minLimit"`
	TemplateSms string          `json:"templateSms"`
	IsActive    bool            `json:"isActive"`
}

func (TransactionRules) TableName() string {
	return "transaction_rules"
}

type Transactions struct {
	ID        uint64    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	CardID    int       `json:"cardID,omitempty"`
	Card      *Cards    `gorm:"references:CardID" json:"card,omitempty"`
	ToCardID  int       `json:"ToCardID,omitempty"`
	ToCard    *Cards    `gorm:"foreignKey:ToCardID" json:"toCard,omitempty"`
	//`gorm:"references:CardID,foreignKey:ToCardID"`
	Amount   decimal.Decimal `json:"amount,omitempty"`
	IsActive bool            `json:"isActive,omitempty"`
}

func (Transactions) TableName() string {
	return "transactions"
}

type TransferRequest struct {
	SourceCardNo int             `json:"sourceCardNo"`
	TargetCardNo int             `json:"targetCardNo"`
	Amount       decimal.Decimal `json:"amount"`
}
type LatestInfo struct {
	CardID int
	Count  int
}

type LatestTXn struct {
	Name         string         `json:"name"`
	CustomerID   int            `json:"customerID"`
	Transactions []Transactions `json:"transactions,omitempty"`
}

type CardAccess struct {
	Contained     bool
	AccountAccess AccountAccess
}

type AccountAccess struct {
	Contained      bool
	CustomerAccess CustomerAccess
}
type CustomerAccess struct {
	Contained bool
}

/*
	c := domain.CardAccess{
		Contained: true,
		AccountAccess: domain.AccountAccess{
			Contained: true,
			CustomerAccess: domain.CustomerAccess{
				Contained: true,
			},
		},
	}
*/

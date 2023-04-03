package domain

import "github.com/pkg/errors"

var ErrNoUpdateAccount = errors.New("cannot update account")
var ErrReadTransactionRules = errors.New("cannot read transactionRules")
var ErrInsertTransaction = errors.New("cannot insert transaction")
var ErrAmountNotInRange = errors.New("amount not in limit range")
var ErrReadCard = errors.New("cannot Read card")
var ErrNotValidFromCard = errors.New("not valid source card")
var ErrNotValidToCard = errors.New("not valid target card")
var ErrNotEnoughCredit = errors.New("not enough credit in source account")
var ErrNoActiveSMSService = errors.New("there is no Active SMS service")

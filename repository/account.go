package repository

import (
	"shopChallenge/domain"
)

func (r RepoImpl) UpdateAccount(account domain.Accounts) error {
	account.Customer = domain.Customers{}
	res := r.DB.Save(account)
	if res.Error != nil || res.RowsAffected == 0 {
		return domain.ErrNoUpdateAccount
	}
	return nil
}

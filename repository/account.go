package repository

import (
	"fmt"
	"shopChallenge/domain"
)

func (r RepoImpl) UpdateAccount(account domain.Accounts) error {
	fmt.Println("update ----->", account)
	fmt.Println("update db----->", r.db)
	account.Customer = domain.Customers{}
	res := r.db.Save(account)
	if res.Error != nil || res.RowsAffected == 0 {
		return domain.ErrNoUpdateAccount
	}
	return nil
}

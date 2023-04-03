package repository

import "fmt"

func (r *RepoImpl) StartTransaction() error {
	r.db.Begin()
	if r.db.Error != nil {
		return nil
	}
	return r.db.Error
}

func (r *RepoImpl) FinalizeTransaction(err error) error {
	if err == nil {
		fmt.Println("commit part", err)
		r.db.Commit()
		return r.db.Error
	}
	r.db.Rollback()
	fmt.Println("rollback ----", err)
	return r.db.Error
}

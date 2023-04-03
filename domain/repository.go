package domain

type Repo interface {
	InsertTransaction(transactions Transactions) error
	GetLatestTXNsOfActiveUsers() ([]LatestInfo, error)
	GetLatestTXNs(cardNo int, access CardAccess) ([]Transactions, error)

	ReadCard(cardNo int, access AccountAccess) (Cards, error)
	//ReadOnlyCard(cardNo int) (Cards, error)

	ReadTransactionRules(transaction string) (TransactionRules, error)

	UpdateAccount(account Accounts) error

	StartTransaction() error
	FinalizeTransaction(err error) error
}

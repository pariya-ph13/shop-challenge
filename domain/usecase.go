package domain

type Usecase interface {
	GetLatestTXNsOfMostActiveUsers() ([]LatestTXn, error)
	Transfer(request TransferRequest) error
}

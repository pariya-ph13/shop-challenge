package delivery

import "shopChallenge/domain"

type DataHandlerImpl struct {
	usecase domain.Usecase
}

func NewDataHandler(usecase domain.Usecase) domain.DataHandler {
	return &DataHandlerImpl{usecase: usecase}
}

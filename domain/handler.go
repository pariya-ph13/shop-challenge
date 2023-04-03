package domain

import "github.com/gin-gonic/gin"

type DataHandler interface {
	GetLatestTXNsOfMostActiveUsers(ctx *gin.Context)
	TransferMoney(ctx *gin.Context)
}

package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shopChallenge/domain"
)

func (d *DataHandlerImpl) GetLatestTXNsOfMostActiveUsers(
	ctx *gin.Context) {
	res, err := d.usecase.GetLatestTXNsOfMostActiveUsers()
	if err != nil {
		ctx.JSON(500, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(200, gin.H{
		"data": res,
	})
}

func (d *DataHandlerImpl) TransferMoney(
	ctx *gin.Context) {
	var req domain.TransferRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"message": "error in bind request",
			"err":     err,
		})
		return
	}
	fmt.Println("++++++transition request is:", req)
	err := d.usecase.Transfer(&req)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "transaction is failed!.",
			"err":     err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "transaction is successful.",
		"request": req,
	})
}

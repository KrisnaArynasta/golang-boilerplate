package Controller

import (
	"fmt"
	"net/http"
	"strconv"
	Interface "testing-api/interface"
	Loghelper "testing-api/loghelper"
	ApiRequestModel "testing-api/model/api/request"
	ApiResponseModel "testing-api/model/api/response"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService Interface.TransactionService
}

func (tc *TransactionController) LoadData(c *gin.Context) {

	method := c.Query("method")

	data, err := tc.TransactionService.LoadData(c, method)
	if err != nil {
		Loghelper.WriteLog().Error().Msg(err.Error())
		fmt.Println("Error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (tc *TransactionController) PostData(c *gin.Context) {
	var dataRequest ApiRequestModel.PostDataRequest
	// 	// if err := c.ShouldBindJSON(&dataRequest); err != nil {
	// 	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	// 	return
	// 	// }
	dataRequest = ApiRequestModel.PostDataRequest{
		OrderId:                   c.PostForm("orderId"),
		Amount:                    c.PostForm("amount"),
		Currency:                  c.PostForm("currency"),
		DateTime:                  c.PostForm("dateTime"),
		Language:                  c.PostForm("language"),
		DepositorUserId:           c.PostForm("depositorUserId"),
		DepositMethod:             "THB_QR",
		CallbackUrl:               c.PostForm("callbackUrl"),
		RedirectUrl:               c.PostForm("redirectUrl"),
		CancelUrl:                 c.PostForm("redirectUrl"),
		MessageAuthenticationCode: c.PostForm("messageauthenticationcode"),
		DepositorName:             c.PostForm("depositorName"),
		DepositorBank:             c.PostForm("depositorBank"),
	}

	data, err := tc.TransactionService.PostData(c, dataRequest)
	if err != nil {
		Loghelper.WriteLog().Error().Msg(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if data.Correlation != "" || data.Message != "" {
		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Correlation: %s, Message: %s", data.Correlation, data.Message))
		c.JSON(http.StatusConflict, ApiResponseModel.ErrorResponse{Message: data.Message})
		return
	}

	err = tc.TransactionService.InsertToDatabase(c, data)
	if err != nil {
		Loghelper.WriteLog().Error().Msg(err.Error())
		c.JSON(http.StatusInternalServerError, ApiResponseModel.ErrorResponse{Message: "Something went wrong! Please contact Administrator"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (tc *TransactionController) LoadDataFromDatabase(c *gin.Context) {
	id := 0
	id, _ = strconv.Atoi(c.Query("id"))

	var idValue interface{}
	if id == 0 {
		idValue = nil
	} else {
		idValue = id
	}

	data, err := tc.TransactionService.GetFromDatabase(c, idValue)
	if err != nil {
		Loghelper.WriteLog().Error().Msg(err.Error())
		c.JSON(http.StatusInternalServerError, ApiResponseModel.ErrorResponse{Message: err.Error()})
		return
	}

	response := ApiResponseModel.TransactionDataResponse{
		Success: true,
		Message: "",
		Data:    data,
	}
	c.JSON(http.StatusOK, response)
}

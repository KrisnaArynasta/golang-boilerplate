package Interface

import (
	"context"
	ApiRequestModel "testing-api/model/api/request"
	ApiResponseModel "testing-api/model/api/response"
	DataBaseModel "testing-api/model/database"
)

type TransactionRepository interface {
	LoadData(c context.Context) (ApiResponseModel.PaymentMethodResponse, error)
	PostData(c context.Context, requestData ApiRequestModel.PostDataRequest) (ApiResponseModel.RequestDepositResponse, error)
}

type TransactionService interface {
	LoadData(c context.Context, method string) (ApiResponseModel.PaymentMethodResponse, error)
	PostData(c context.Context, requestData ApiRequestModel.PostDataRequest) (ApiResponseModel.RequestDepositResponse, error)
	InsertToDatabase(c context.Context, requestData ApiResponseModel.RequestDepositResponse) error
	GetFromDatabase(c context.Context, id interface{}) ([]DataBaseModel.TransactionDataFromDatabase, error)
}

type TransactionDatabase interface {
	Insert(c context.Context, data ApiResponseModel.RequestDepositResponse) error
	GetData(c context.Context, id interface{}) ([]DataBaseModel.TransactionDataFromDatabase, error)
}

package Mocks

import (
	"context"
	"errors"
	Interface "testing-api/interface"
	ApiRequestModel "testing-api/model/api/request"
	ApiResponseModel "testing-api/model/api/response"

	"github.com/stretchr/testify/mock"
)

type TransactionRepository struct {
	mock.Mock
}

func NewTransactionRepository() Interface.TransactionRepository {
	return &TransactionRepository{}
}

func (_m *TransactionRepository) LoadData(c context.Context) (ApiResponseModel.PaymentMethodResponse, error) {
	argument :=
		_m.Mock.Called(c)
	if argument.Get(0) == nil {
		return ApiResponseModel.PaymentMethodResponse{}, errors.New("Error!")
	} else {
		return argument.Get(0).(ApiResponseModel.PaymentMethodResponse), nil
	}
}

// PostData implements Domain.TestRepository.
func (_m *TransactionRepository) PostData(c context.Context, requestData ApiRequestModel.PostDataRequest) (ApiResponseModel.RequestDepositResponse, error) {
	argument :=
		_m.Mock.Called(c)
	if argument.Get(0) == nil {
		return ApiResponseModel.RequestDepositResponse{}, errors.New("Error!")
	} else {
		return argument.Get(0).(ApiResponseModel.RequestDepositResponse), nil
	}
}

package Mocks

import (
	"context"
	"errors"
	Interface "testing-api/interface"
	ApiResponseModel "testing-api/model/api/response"
	DataBaseModel "testing-api/model/database"

	"github.com/stretchr/testify/mock"
)

type TransactionDatabase struct {
	mock.Mock
}

func NewTransactionDatabase() Interface.TransactionDatabase {
	return &TransactionDatabase{}
}

func (_m *TransactionDatabase) Insert(c context.Context, data ApiResponseModel.RequestDepositResponse) error {
	argument :=
		_m.Mock.Called(c)
	if argument.Get(0) == nil {
		return nil
	} else {
		return errors.New("Error!")
	}
}

func (_m *TransactionDatabase) GetData(c context.Context, id interface{}) ([]DataBaseModel.TransactionDataFromDatabase, error) {
	argument :=
		_m.Mock.Called(c)
	if argument.Get(0) == nil {
		return []DataBaseModel.TransactionDataFromDatabase{}, errors.New("Error!")
	} else {
		return argument.Get(0).([]DataBaseModel.TransactionDataFromDatabase), nil
	}
}

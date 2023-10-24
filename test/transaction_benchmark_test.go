package test

//go test -v -bench. ./...

import (
	"context"
	"encoding/json"
	"testing"
	Mocks "testing-api/mocks"
	ApiRequestModel "testing-api/model/api/request"
	ApiResponseModel "testing-api/model/api/response"
	Service "testing-api/service"
	"time"

	"github.com/stretchr/testify/mock"
)

func BenchmarkLoadDataSpecifyPaymentMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mockTestingRepository := new(Mocks.TransactionRepository)
		mockTestingDatabase := new(Mocks.TransactionDatabase)
		service := Service.NewTransactionService(mockTestingRepository, time.Second*2, mockTestingDatabase)

		request := "thb_qr"
		var mockResponse ApiResponseModel.PaymentMethodResponse
		json.Unmarshal([]byte(mockJsonLoadData), &mockResponse)

		mockTestingRepository.Mock.On("LoadData", mock.Anything).Return(mockResponse, nil).Once()
		service.LoadData(context.Background(), request)
	}
}

func BenchmarkPostDataToProvider(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mockTestingRepository := new(Mocks.TransactionRepository)
		mockTestingDatabase := new(Mocks.TransactionDatabase)
		service := Service.NewTransactionService(mockTestingRepository, time.Second*2, mockTestingDatabase)

		request := ApiRequestModel.PostDataRequest{
			OrderId:         "CCCX001",
			Amount:          "10",
			Currency:        "THB",
			DateTime:        "20230811T085436+07:00",
			Language:        "en-Us",
			DepositorUserId: "DX000001",
			DepositMethod:   "THB_QR",
			CallbackUrl:     "https://7b2d-36-68-218-190.ngrok-free.app/api/qaicash/deposit-status",
			RedirectUrl:     "https://www.google.com",
			CancelUrl:       "https://www.google.com",
			//MessageAuthenticationCode: c.PostForm("messageauthenticationcode"),
			DepositorName: "John Smith",
			//DepositorBank:             c.PostForm("depositorBank"),
		}

		var mockResponse ApiResponseModel.RequestDepositResponse
		json.Unmarshal([]byte(mockJsonPostData), &mockResponse)

		mockTestingRepository.Mock.On("PostData", mock.Anything).Return(mockResponse, nil).Once()

		service.PostData(context.Background(), request)
		//mockTestingRepository.AssertExpectations(b)
	}
}

func BenchmarkLoadDataFromDatabase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mockTestingRepository := new(Mocks.TransactionRepository)
		mockTestingDatabase := new(Mocks.TransactionDatabase)
		service := Service.NewTransactionService(mockTestingRepository, time.Second*2, mockTestingDatabase)

		var request = 1

		mockTestingDatabase.Mock.On("GetData", mock.Anything).Return(mockTransactionData, nil).Once()
		service.GetFromDatabase(context.Background(), request)
	}
}

package test

//go test ./...

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
	Mocks "testing-api/mocks"
	ApiRequestModel "testing-api/model/api/request"
	ApiResponseModel "testing-api/model/api/response"
	Service "testing-api/service"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockJsonLoadData = `[
    {
        "method": "THB_QR",
        "details": [
            {
                "currency": "THB",
                "limits": [
                    {
                        "minTier": 0,
                        "maxTier": 3,
                        "minTransactionAmount": 1,
                        "maxTransactionAmount": 100000050,
                        "bankBased": false
                    }
                ]
            }
        ]
    }
]`

var mockJsonPostData = `{
    "depositTransaction": {
        "orderId": "CCCX001",
        "transactionId": 305362,
        "dateCreated": "2023-08-11T09:54:36.922+08:00",
        "depositMethod": "THB_QR",
        "processor": "MOCKPAY",
        "amount": 10,
        "currency": "THB",
        "status": "PENDING",
        "dateUpdated": "2023-08-11T09:54:36.993+08:00",
        "depositorUserId": "CCCX001",
        "notes": "2023-08-11 09:54:36.923 Transaction is created",
        "channel": "THB_QR_MOCKPAY_THB",
        "messageAuthenticationCode": "f61ac8a48f8b771995da231e71a1e04e23c1bc02cff538b292817f868396b223",
        "processingCurrency": "THB",
        "processingAmount": 10
    },
    "paymentPageSession": {
        "sessionToken": "2aac0186585777b9678e9b475869c846d81672d0",
        "expires": "2023-08-11T09:59:36.96+08:00",
        "sessionType": "DEPOSIT",
        "orderId": "CCCX001",
        "paymentPageUrl": "https://cashier.mekong-300.com/deposit/CCCX001/2aac0186585777b9678e9b475869c846d81672d0/en"
    },
    "success": true,
    "orderId": "CCCX001",
    "messageAuthenticationCode": "5ec958f4303d025b8f1c2f70b5c3d64c771e92ddb92fadca95ed1d39c8d57001",
    "correlation": "",
    "message": ""
}`

func TestLoadDataSpecifyPaymentMethod(t *testing.T) {
	mockTestingRepository := new(Mocks.TransactionRepository)
	service := Service.NewTransactionService(mockTestingRepository, time.Second*2)

	request := "thb_qr"
	// dataDetails := []Domain.Detail{}
	// dataDetail := Domain.Detail{
	// 	Currency: "THB",
	// 	Limits:   nil,
	// }
	// dataDetails = append(dataDetails, dataDetail)

	// dataPayment := Domain.PaymentMethodElement{
	// 	Method:  "THB_QR",
	// 	Details: dataDetails,
	// }
	// mockResponse := Domain.PaymentMethodResponse{}
	// mockResponse = append(mockResponse, dataPayment)
	var mockResponse ApiResponseModel.PaymentMethodResponse
	json.Unmarshal([]byte(mockJsonLoadData), &mockResponse)

	mockTestingRepository.Mock.On("LoadData", mock.Anything).Return(mockResponse, nil).Once()
	result, err := service.LoadData(context.Background(), request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, strings.ToLower(result[0].Method), strings.ToLower(request))

	//mockTestingRepository.AssertExpectations(t)
}

func TestPostDataSpecifyPaymentMethod(t *testing.T) {
	mockTestingRepository := new(Mocks.TransactionRepository)
	service := Service.NewTransactionService(mockTestingRepository, time.Second*2)

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
	result, err := service.PostData(context.Background(), request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.DepositTransaction.OrderId, request.OrderId)
}

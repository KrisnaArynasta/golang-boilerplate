package Service

import (
	"context"
	"strings"
	Interface "testing-api/interface"
	ApiRequestModel "testing-api/model/api/request"
	ApiResponseModel "testing-api/model/api/response"
	"time"
)

type transactionService struct {
	transactionRepository Interface.TransactionRepository
	contextTimeout        time.Duration
}

func NewTransactionService(transactionRepository Interface.TransactionRepository, timeout time.Duration) Interface.TransactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
		contextTimeout:        timeout,
	}
}

func (tu *transactionService) LoadData(c context.Context, method string) (ApiResponseModel.PaymentMethodResponse, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()

	data, err := tu.transactionRepository.LoadData(ctx)

	var result ApiResponseModel.PaymentMethodResponse
	if method != "" {
		result = ApiResponseModel.PaymentMethodResponse{}
		for _, paymentMethodElement := range data {
			if strings.EqualFold(paymentMethodElement.Method, method) {
				result = append(result, paymentMethodElement)
				break
			}
		}
		data = result
	}

	return data, err
}

func (tu *transactionService) PostData(c context.Context, requestData ApiRequestModel.PostDataRequest) (ApiResponseModel.RequestDepositResponse, error) {
	//ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	//defer cancel()
	return tu.transactionRepository.PostData(c, requestData)
}

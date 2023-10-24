package Repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	Bootstrap "testing-api/bootstrap"
	Helper "testing-api/helper"
	Interface "testing-api/interface"
	Loghelper "testing-api/loghelper"
	ApiRequestModel "testing-api/model/api/request"
	ApiResponseModel "testing-api/model/api/response"
	"time"
)

type transactionRepository struct {
	env *Bootstrap.Env
}

func NewTransactionRepository(env *Bootstrap.Env) Interface.TransactionRepository {
	return &transactionRepository{
		env: env,
	}
}

func (tr *transactionRepository) LoadData(c context.Context) (ApiResponseModel.PaymentMethodResponse, error) {

	var qaicashMerchantId = "1"
	var secretKey = "secret"
	var hash = Helper.HmacSHA256(qaicashMerchantId, secretKey)
	baseURL := tr.env.ExternalApiBaseUrl + qaicashMerchantId + "/deposit/routing/methods/complete"

	params := url.Values{}
	params.Set("userId", qaicashMerchantId)
	params.Set("deviceType", "PC")
	params.Set("hmac", hash)

	finalURL := baseURL + "?" + params.Encode()

	log.Println("get data from " + finalURL)
	response, err := http.Get(finalURL)
	if err != nil {
		log.Println("Error making HTTP GET request:", err)
		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Error making HTTP GET request: %s", err))
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Println("Request failed with status code:", response.StatusCode)
		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Request failed with status code: %d", response.StatusCode))
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error", err)
		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Error: %s", err))
	}

	var paymentMethod ApiResponseModel.PaymentMethodResponse
	err = json.Unmarshal(responseData, &paymentMethod)
	if err != nil {
		log.Println("Error unmarshaling JSON data:", err)
		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Error unmarshaling JSON data: %s", err))
	}

	return paymentMethod, err
}

func (tr *transactionRepository) PostData(c context.Context, requestData ApiRequestModel.PostDataRequest) (ApiResponseModel.RequestDepositResponse, error) {
	qaicashMerchantId := "1"
	secretKey := "secret"
	currentTime := time.Now().UTC()
	currentDate := currentTime.Format("20060102T150405Z")

	hashData := fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s",
		qaicashMerchantId,
		requestData.OrderId,
		requestData.Amount,
		requestData.Currency,
		currentDate,
		requestData.OrderId,
		requestData.DepositMethod,
	)
	hash := Helper.HmacSHA256(hashData, secretKey)

	// params := url.Values{}
	// params.Set("orderId", requestData.OrderId)
	// params.Set("amount", requestData.Amount)
	// params.Set("currency", requestData.Currency)
	// params.Set("dateTime", currentDate)
	// params.Set("language", "en-Us")
	// params.Set("depositorUserId", requestData.OrderId)
	// params.Set("depositMethod", requestData.DepositMethod)
	// params.Set("callbackUrl", requestData.CallbackUrl)
	// params.Set("redirectUrl", requestData.RedirectUrl)
	// params.Set("redirectUrl", requestData.CancelUrl)
	// params.Set("messageAuthenticationCode", hash)
	// params.Set("depositorName", requestData.DepositorName)
	// params.Set("depositorBank", requestData.DepositorBank)

	params := url.Values{
		"orderId":                   {requestData.OrderId},
		"amount":                    {requestData.Amount},
		"currency":                  {requestData.Currency},
		"dateTime":                  {currentDate},
		"language":                  {"en-Us"},
		"depositorUserId":           {requestData.OrderId},
		"depositMethod":             {requestData.DepositMethod},
		"callbackUrl":               {requestData.CallbackUrl},
		"redirectUrl":               {requestData.RedirectUrl},
		"cancelUrl":                 {requestData.CancelUrl},
		"messageAuthenticationCode": {hash},
		"depositorName":             {requestData.DepositorName},
		//"depositorBank":             {requestData.DepositorBank},
	}

	baseURL := "https://public-services.mekong-300.com/ago/integration/v2.0/" + qaicashMerchantId + "/deposit/"

	response, err := http.PostForm(baseURL, params)
	if err != nil {
		log.Println("Error making HTTP GET request:", err)
		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Request failed with status code: %s", err))
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		log.Println("Request failed with status code:", response.StatusCode)
		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Request failed with status code: %d", response.StatusCode))
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error", err)
		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Error: %s", err))
	}
	// bodyString := string(responseData)
	// log.Print(bodyString)

	var paymentMethod ApiResponseModel.RequestDepositResponse
	err = json.Unmarshal(responseData, &paymentMethod)
	if err != nil {
		log.Println("Error unmarshaling JSON data:", err)
		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Error unmarshaling JSON data: %s", err))
	}

	return paymentMethod, err
}

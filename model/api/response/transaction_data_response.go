package Model

import (
	DataBaseModel "testing-api/model/database"
)

type TransactionDataResponse struct {
	Success bool                                        `json:"success"`
	Message string                                      `json:"message"`
	Data    []DataBaseModel.TransactionDataFromDatabase `json:"data"`
}

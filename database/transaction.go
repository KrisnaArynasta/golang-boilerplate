package Database

import (
	"context"
	"databaseservice"
	"fmt"
	"log"
	"strconv"
	Interface "testing-api/interface"
	Loghelper "testing-api/loghelper"
	ApiResponseModel "testing-api/model/api/response"
	DataBaseModel "testing-api/model/database"
)

type transactionDatabase struct {
	database *databaseservice.DatabaseService
}

func NewTransactionDatabase(db *databaseservice.DatabaseService) Interface.TransactionDatabase {
	return &transactionDatabase{
		database: db,
	}
}

func (td *transactionDatabase) Insert(c context.Context, data ApiResponseModel.RequestDepositResponse) error {

	// currentTimeUTC := time.Now().UTC()
	// location := time.FixedZone("GMT+8", 8*60*60)
	// currentDate := currentTimeUTC.In(location)

	// query := "INSERT INTO TestingData (data) VALUES (?)"
	// _, err := td.database.Exec(query, currentDate)
	// //tr.database.ExecContext(c, query)
	// if err != nil {
	// 	fmt.Println("Error executing query:", err)
	// 	return false, err
	// }

	// stmt, err := td.database.PrepareContext(c,
	// 	"[dbo].[InsertTestingData] "+
	// 		"@amount = ?, "+
	// 		"@orderId = ?, "+
	// 		"@transactionId = ?",
	// )
	// if err != nil {
	//  Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Error executing query: %s", err))
	// 	return false, err
	// }

	// _, err = stmt.ExecContext(c,
	// 	data.DepositTransaction.Amount,
	// 	data.DepositTransaction.OrderId,
	// 	data.DepositTransaction.TransactionId,
	// )
	// if err != nil {
	//  Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Error executing query: %s", err))
	// 	return false, err
	// }
	// defer stmt.Close()

	// return true, err

	params := map[string]any{
		"@amount":        data.DepositTransaction.Amount,
		"@orderId":       data.DepositTransaction.OrderId,
		"@transactionId": strconv.Itoa(data.DepositTransaction.TransactionId),
	}

	log.Print(params)

	_, err := td.database.CallStoredProcedure(context.Background(), "[dbo].[InsertTransactionData]", params)
	if err != nil {
		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Error calling stored procedure: %s", err))
		return err
	}

	return nil
}

func (td *transactionDatabase) GetData(c context.Context, id interface{}) ([]DataBaseModel.TransactionDataFromDatabase, error) {
	var transactionDataList []DataBaseModel.TransactionDataFromDatabase

	params := map[string]any{
		"@id": id,
	}

	rows, err := td.database.CallStoredProcedure(context.Background(), "[dbo].[GetTransactionData]", params)
	if err != nil {
		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Error calling stored procedure: %s", err))
		return transactionDataList, err
	}
	defer rows.Close()

	for rows.Next() {
		var data DataBaseModel.TransactionDataFromDatabase
		if err := rows.Scan(&data.Id, &data.Amount, &data.OrderId, &data.TransactionId); err != nil {
			Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Error fetching rows: %s", err))
			return transactionDataList, err
		}
		transactionDataList = append(transactionDataList, data)
	}
	if err = rows.Err(); err != nil {
		Loghelper.WriteLog().Error().Msg(fmt.Sprintf("Error calling stored procedure: %s", err))
		return transactionDataList, err
	}
	return transactionDataList, err
}

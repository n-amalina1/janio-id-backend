package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"janio-id-backend/models"
	"net/http"
)

func GetOrdersIDProvider(db *sql.DB) (models.IDProviderOrdersParams, error) {
	var (
		IDProviderOrdersParams models.IDProviderOrdersParams
	)

	res, err := http.Get("http://localhost:8080/id/orders")

	if err != nil {
		panic(err.Error())
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(body, &IDProviderOrdersParams)

	return IDProviderOrdersParams, nil
}

func PostStatusIDProvider(db *sql.DB, status models.IDOrderStatus) (models.IDOrderStatus, error) {
	json, err := json.Marshal(
		models.IDOrderStatus{
			OrderID:     status.OrderID,
			OrderStatus: status.OrderStatus,
		})
	if err != nil {
		return models.IDOrderStatus{}, fmt.Errorf("update Status DB: %v", err)
	}

	_, errS := http.Post("http://localhost:8080/id/order/update", "application/json", bytes.NewBuffer(json))
	if errS != nil {
		return models.IDOrderStatus{}, fmt.Errorf("update Status DB: %v", errS)
	}

	return status, nil
}

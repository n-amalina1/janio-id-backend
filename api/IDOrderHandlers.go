package api

import (
	"database/sql"
	"encoding/json"
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

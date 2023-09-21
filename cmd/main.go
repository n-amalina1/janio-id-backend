package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"janio-id-backend/models"
	routes "janio-id-backend/routes"
)

var db *sql.DB

func main() {
	updateStatus()
	routes.SetupRoutes(db)

}

func updateStatus() {
	json, err := json.Marshal(models.IDOrderStatus{
		OrderID:     148,
		OrderStatus: "Completed",
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, errS := http.Post("http://localhost:8080/id/order/update", "application/json", bytes.NewBuffer(json))
	if errS != nil {
		log.Fatalln(errS)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	fmt.Printf("%s\n", sb)

}

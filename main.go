package main

import (
	"database/sql"
	"fmt"

	"github.com.caiommdev/go-intensivo/internal/infra/database"
	"github.com.caiommdev/go-intensivo/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRespository(db)

	calculateFinalPrice := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   13.34,
	}

	output, err := calculateFinalPrice.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}

package models

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestOrder(t *testing.T) {
	order := &Order{
		Products:  []Product{{Id: 1, Category: "Food", Img: "123.png", Title: "Shashlik", Price: "120.50", Count: 3}, {Id: 2, Category: "Food", Img: "333.png", Title: "Sauce", Price: "30.75", Count: 1}},
		Data:      OrderData{Adress: "Orel, sm. Per 10", Mobile: "89991002345671", Time: time.Now().Format("02.01.2006")},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	file, _ := os.Create("order.json")
	bytes, _ := json.Marshal(order)
	_, err := file.Write(bytes)
	if err != nil {
		log.Panic(err)
	}

	fmt.Print(order.String())

	return
}

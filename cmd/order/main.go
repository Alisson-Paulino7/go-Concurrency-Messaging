package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/Alisson-Paulino7/go-concurrence-messaging/internal/infra/database"
	"github.com/Alisson-Paulino7/go-concurrence-messaging/internal/usecase"
	"github.com/Alisson-Paulino7/go-concurrence-messaging/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3" //Usar o _ quando não estiver usando o pacote diretamente no código
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel) // escutando a fila / trava / T2

	rabbitmqWorker(msgRabbitmqChannel, uc)
}

func rabbitmqWorker(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {

	fmt.Println("Starting rabbitmq")

	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}
		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println("Mensagem processada e salva no banco", output)
	}
}

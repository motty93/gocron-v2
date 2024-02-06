package main

import (
	"log"
	"net/http"

	"github.com/motty93/gocron-v2/services"
)

func task() {
	println("Task")
}

func main() {
	ss, err := services.NewSchedulerService()
	if err != nil {
		panic(err)
	}

	ss.NewDailyJob(task)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
	"./app"
	"./app/repository"
	// "./app/usecases/entity"
	// log "github.com/sirupsen/logrus"
)

func main() {
	db := repository.GetConnectDataBase()
	app.Controller(db)

	// log.WithFields(log.Fields{
	// 	"animal": "walrus",
	// 	"number": 1,
	// 	"size":   10,
	// }).Info("A walrus appears")
	// person := entity.Person{2, "Billy", "Agrome", 12}
	// log.WithFields(log.Fields{
	// 	"id":    person.ID,
	// 	"fname": person.Fname,
	// 	"lname": person.Lname,
	// 	"age":   person.Age,
	// }).Info("A person data")
}

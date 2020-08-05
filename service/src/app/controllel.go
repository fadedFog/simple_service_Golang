package app

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"./usecases"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

var db *sql.DB

//Controller ...
func Controller(database *sql.DB) {
	e := echo.New()
	db = database

	e.GET("/", meet)
	e.GET("/people", showPeopleService)
	e.GET("/people/get_person", getPerson)
	e.GET("/people/get_people", getPeople)
	e.GET("/people/update", updatePerson)
	e.GET("/people/add", addPerson)
	e.GET("/people/delete", dropPerson)

	e.Logger.Fatal(e.Start(":1323"))
}

func meet(c echo.Context) error {
	log.WithFields(log.Fields{}).Info("A main page of meeting.")

	return c.String(http.StatusOK, "Welcome to the server!")
}

func showPeopleService(c echo.Context) error {
	log.WithFields(log.Fields{}).Info("A page servic people.")

	return c.String(http.StatusOK, "Page start service people.")
}

func getPerson(c echo.Context) error {
	idPerson := c.QueryParam("id")

	id, _ := strconv.Atoi(idPerson)
	dataP, isItP := usecases.FuncGetPerson(db, id)

	if isItP {
		logGetPerson("The data is transmitted. Id: " + idPerson)
		return c.JSON(http.StatusOK, fmt.Sprintln(dataP))
	}
	logGetPerson("The data was not transmitted, perhaps there is no such id")
	return c.String(http.StatusOK, fmt.Sprintf("Person with id: %v is not found!", id))
}
func logGetPerson(status string) {
	log.WithFields(log.Fields{
		"status": status,
		"method": "getPerson(e echo.Context)",
	}).Info("Operation - Get person by id")
}

func getPeople(c echo.Context) error {
	limitLine := c.QueryParam("limit")
	offsetLine := c.QueryParam("offset")

	limit, _ := strconv.Atoi(limitLine)
	offset, _ := strconv.Atoi(offsetLine)

	dataPs := usecases.FuncGetPeople(db, limit, offset)

	log.WithFields(log.Fields{
		"method":       "getPeople",
		"dataOfPeople": string(dataPs),
	}).Info("Data of people was transmitted")

	return c.JSON(http.StatusOK, fmt.Sprintln(dataPs))
}

func updatePerson(c echo.Context) error {
	idLine := c.QueryParam("id")
	fname := c.QueryParam("fname")
	lname := c.QueryParam("lname")
	ageLine := c.QueryParam("age")

	id, _ := strconv.Atoi(idLine)
	age, _ := strconv.Atoi(ageLine)

	wasUpdate := usecases.FuncUpdatePerson(db, id, fname, lname, age)

	if wasUpdate {
		logUpdatePerson("Person with " + idLine + " has new data")
		return c.String(http.StatusOK, fmt.Sprintf("Person with id: %v was update data", id))
	}
	logUpdatePerson("Failed to update user data. Probably, id: " + idLine + " not exist")
	return c.String(http.StatusOK, fmt.Sprintf("Person with id: %v was not update data (Probably, the persin is not real in db)", id))
}
func logUpdatePerson(status string) {
	log.WithFields(log.Fields{
		"status": status,
		"method": "updatePerson(e echo.Context)",
	}).Info("Operation - Update person by id")
}

func addPerson(c echo.Context) error {
	idLine := c.QueryParam("id")
	fname := c.QueryParam("fname")
	lname := c.QueryParam("lname")
	ageLine := c.QueryParam("age")

	id, _ := strconv.Atoi(idLine)
	age, _ := strconv.Atoi(ageLine)

	wasAdd := usecases.FuncAddPerson(db, id, fname, lname, age)

	if wasAdd {
		logAddPerson("The new person was added to the database.", []string{idLine, fname, lname, ageLine})
		return c.String(http.StatusOK, fmt.Sprintf("Person with id: %v was add in database. ", id))
	}
	logAddPerson("Having problems adding a new person.", []string{"error"})
	return c.String(http.StatusOK, fmt.Sprintf("Person with id: %v was not add in database. ", id))
}
func logAddPerson(status string, data []string) {
	log.WithFields(log.Fields{
		"status": status,
		"data":   data,
		"method": "addPerson(e echo.Context)",
	}).Info("Operation - Put new person's data")
}

func dropPerson(c echo.Context) error {
	idLine := c.QueryParam("id")

	id, _ := strconv.Atoi(idLine)

	wasDrop := usecases.FuncDropPerson(id, db)

	if wasDrop {
		logDropPerson("Person with id: " + idLine + " was deleted from database.")
		return c.String(http.StatusOK, fmt.Sprintf("Person with id: %v was drop in database. ", id))
	}
	logDropPerson("Person with id: " + idLine + " wasn't deleted from database. Probably, id: " + idLine + " not exist")
	return c.String(http.StatusOK, fmt.Sprintf("Person with id: %v was not drop in database! ", id))
}
func logDropPerson(status string) {
	log.WithFields(log.Fields{
		"status": status,
		"method": "dropPerson(e echo.Context)",
	}).Info("Operation - Delete person by id")
}

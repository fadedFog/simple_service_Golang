package usecases

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"./entity"
	log "github.com/sirupsen/logrus"
)

func isPerson(p entity.Person) bool {
	if p.ID <= 0 {
		return false
	}
	return true
}

func logAboutCRUD(info string) {
	log.WithFields(log.Fields{}).Info(info)
}

// FuncGetPerson ...
func FuncGetPerson(db *sql.DB, id int) ([]byte, bool) {
	var person entity.Person
	var data []byte
	for {
		person = getPerson(id, db)
		logAboutCRUD("Person  was geting from bd.")
		if !isPerson(person) {
			logAboutCRUD("It is not person.")
			break
		} else {
			data, _ = json.Marshal(&person)
			logAboutCRUD("Data pesron is passed in JSON type.")
			return data, true
		}
	}
	logAboutCRUD("Empty data is passed.")
	return data, false
}

func getPerson(idP int, db *sql.DB) entity.Person {
	rows, err := db.Query("SELECT * FROM `productdb`.`person` WHERE id = ?", idP)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var person entity.Person

	if rows.Next() {
		err := rows.Scan(&person.ID, &person.Fname, &person.Lname, &person.Age)
		if err != nil {
			fmt.Println(err)
		}
	}

	return person
}

// FuncGetPeople ...
func FuncGetPeople(db *sql.DB, limit, offset int) []byte {

	people := getPeople(limit, offset, db)
	data, _ := json.Marshal(&people)
	logAboutCRUD("Data people is passed in JSON type.")
	return data

}

func getPeople(limit, offset int, db *sql.DB) []entity.Person {
	rows, err := db.Query("SELECT * FROM `productdb`.`person`")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	slPeop := []entity.Person{}
	var counterL int
	var counterO int

	for rows.Next() {
		if counterO >= offset {
			counterL++
			person := entity.Person{}
			err = rows.Scan(&person.ID, &person.Fname, &person.Lname, &person.Age)
			if err != nil {
				fmt.Println(err)
				continue
			}
			slPeop = append(slPeop, person)

			if counterL >= limit {
				break
			}
		}
		counterO++
	}

	return slPeop
}

// FuncAddPerson ...
func FuncAddPerson(db *sql.DB, id int, fname, lname string, age int) bool {
	person := entity.Person{id, fname, lname, age}
	addPerson(person, db)
	logAboutCRUD("A new person was added")

	newPerson := getPerson(id, db)
	return newPerson == person
}

func addPerson(person entity.Person, db *sql.DB) error {
	_, err := db.Exec("INSERT INTO `productdb`.`person` (id, fname, lname, age) VALUES(?, ?, ?, ?)",
		person.ID, person.Fname, person.Lname, person.Age)

	if err != nil {
	}
	return err
}

// FuncUpdatePerson ...
func FuncUpdatePerson(db *sql.DB, id int, fname, lname string, age int) bool {
	person := entity.Person{id, fname, lname, age}
	updatePerson(person, db)
	logAboutCRUD("A new data was added")

	upPerson := getPerson(id, db)
	return upPerson == upPerson
}

func updatePerson(person entity.Person, db *sql.DB) {
	result, err := db.Exec("UPDATE `productdb`.`person` SET fname = ?, lname = ?, age = ? WHERE id = ?", person.Fname, person.Lname, person.Age, person.ID)

	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
}

// FuncDropPerson ...
func FuncDropPerson(id int, db *sql.DB) bool {
	person := getPerson(id, db)
	if !isPerson(person) {
		fmt.Printf("No one people have no id equals %v\n", id)
		logAboutCRUD("Person didn't drop.")
	}
	logAboutCRUD("Person did drop.")
	return dropPerson(id, db)
}

func dropPerson(id int, db *sql.DB) bool {
	result, err := db.Exec("DELETE FROM `productdb`.`person` WHERE id = ?", id)

	if err != nil {
		return false
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())
	return true
}

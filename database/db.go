package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Creating users table in database
func CreateTable() error {

	godotenv.Load()
	//Connecting to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer db.Close()

	//Creating users Table
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS users(ID SERIAL PRIMARY KEY, TIMESTAMP TIMESTAMP DEFAULT CURRENT_TIMESTAMP, USERNAME TEXT, CHAT_ID INT);`); err != nil {
		return err
	}

	return nil
}

func AddUser(username string, chatID int) error {
	//Connecting to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer db.Close()

	//Creating SQL command
	data := "INSERT INTO users (username, chat_id) VALUES($1, $2)"

	//Execute SQL command in database
	if _, err := db.Exec(data, username, chatID); err != nil {
		return err
	}

	return nil
}

// Create table for user in database
func CreateUserTable(username string) error {

	//Connecting to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS u_%s (NAME TEXT, DOB DATE, NUM INT, PN TEXT, ADDRESS TEXT, PROF TEXT)", username)); err != nil {
		return err
	}

	return nil
}

// Adding information about a person to the table
func AddPerson(username string, person *Person) error {

	//Connecting to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	defer db.Close()

	//Creating SQL command
	data := fmt.Sprintf("INSERT INTO u_%s (name, dob, num, pn, address, prof) VALUES($1, $2, $3, $4, $5, $6)", username)

	//Execute SQL command in database
	if _, err := db.Exec(data, person.Name, person.DoB, person.Num, person.PhoneNumber, person.Address, person.Profession); err != nil {
		return err
	}

	return nil
}

// Getting information about people by name
func GetPeople(username, name string) ([]Person, error) {

	var people []Person

	//Connecting to database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	//Counting number of users
	data := fmt.Sprintf("SELECT * FROM u_%s WHERE name = $1", username)
	rows, err := db.Query(data, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var person Person
		if err := rows.Scan(&person.Name, &person.DoB, &person.Num, &person.PhoneNumber, &person.Address, &person.Profession); err != nil {
			return people, err
		}
		people = append(people, person)
	}
	if err = rows.Err(); err != nil {
		return people, err
	}

	return people, nil
}

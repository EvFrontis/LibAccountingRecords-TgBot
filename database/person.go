package db

import (
	"fmt"
	"time"
)

type Person struct {
	//Age         int
	Name        string
	Num         int
	PhoneNumber string
	Birthdate   time.Time
	Address     string
	Profession  string
}

func (p Person) String() string {
	return fmt.Sprintf("Name: %s  Age: %d Num: %d", p.Name, time.Now().Year()-p.Birthdate.Year(), p.Num)
}

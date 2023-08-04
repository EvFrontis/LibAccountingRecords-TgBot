package db

import (
	"fmt"
	"time"
)

type Person struct {
	Name        string
	DoB         time.Time
	Num         int
	PhoneNumber string
	Address     string
	Profession  string
}

func (p Person) String() string {
	return fmt.Sprintf("%s Age: %d Number: %d", p.Name, time.Now().Year()-p.DoB.Year(), p.Num)
}

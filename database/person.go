package db

import "fmt"

type Person struct {
	Name string
	Age  int
	Num  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %d %d", p.Name, p.Age, p.Num)
}

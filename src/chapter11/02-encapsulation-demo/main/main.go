package main

import (
	"fmt"
	"go-project/src/chapter11/02-encapsulation-demo/model"
)

func main() {
	person := model.NewPerson("Jason")

	person.SetAge(26)
	person.SetSalary(30000)

	fmt.Println(person)
	fmt.Println("person.GetAge()=", person.GetAge())
	fmt.Println("person.GetSalary()=", person.GetSalary())
}

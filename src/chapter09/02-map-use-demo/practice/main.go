package main

import "fmt"

func main() {
	//演示一個key-value的value是`map`的案例，存放3個學生信息，每個學生有name和sex信息。
	var students = make(map[string]map[string]string)

	students["001"] = make(map[string]string)
	students["001"]["name"] = "Jason"
	students["001"]["sex"] = "M"

	students["002"] = make(map[string]string)
	students["002"]["name"] = "Juan"
	students["002"]["sex"] = "F"

	students["003"] = make(map[string]string)
	students["003"]["name"] = "Marry"
	students["003"]["sex"] = "F"

	fmt.Println("students=", students)
}

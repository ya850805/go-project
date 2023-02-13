package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name""` //使用tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	monster1 := Monster{"Monster A", 60, "Skill A"}
	//json.Marshal使用反射
	var jsonStr, _ = json.Marshal(monster1)
	fmt.Println("json string=", string(jsonStr))
}

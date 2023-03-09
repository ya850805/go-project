package main

import (
	"encoding/json"
	"fmt"
)

// Monster 使用tag指定json field名稱，反射機制
type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func unmarshalStruct() {
	//str在真實專案開發中是通過其他方式讀取到的
	str := "{\"name\":\"Monster1\",\"age\":120,\"birthday\":\"2011-11-11\",\"sal\":8000,\"skill\":\"Skill1\"}"

	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)

	if err != nil {
		fmt.Println("反序列化失敗，err=", err)
	} else {
		fmt.Println("反序列化後，monster=", monster)
	}
}

func unmarshalMap() {
	str := "{\"address\":\"TPE\",\"age\":35,\"name\":\"name1\"}"

	var a map[string]interface{}

	//反序列化map不需要make
	//因為make操作被封裝到Unmarshal函數
	err := json.Unmarshal([]byte(str), &a)

	if err != nil {
		fmt.Println("反序列化失敗，err=", err)
	} else {
		fmt.Println("反序列化後，a=", a)
	}
}

func unmarshalSlice() {
	str := "[{\"address\":\"TPE\",\"age\":7,\"name\":\"jack\"},{\"address\":[\"NY\",\"SF\"],\"age\":20,\"name\":\"tom\"}]"

	var slice []map[string]interface{}

	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("反序列化失敗，err=", err)
	} else {
		fmt.Println("反序列化後，slice=", slice)
	}
}

func main() {
	//將json字符串反序列化成結構體、map、切片
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}

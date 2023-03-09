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

func testStruct() {
	monster := Monster{
		Name:     "Monster1",
		Age:      120,
		Birthday: "2011-11-11",
		Sal:      8000,
		Skill:    "Skill1",
	}

	//將monster序列化
	data, err := json.Marshal(&monster)

	if err != nil {
		fmt.Println("序列化失敗 err=", err)
	} else {
		//輸出序列化後的結果
		fmt.Printf("monster序列化後=%v \n", string(data))
	}
}

func testMap() {
	var a map[string]interface{}
	//使用map需要先make
	a = make(map[string]interface{})

	a["name"] = "name1"
	a["age"] = 35
	a["address"] = "TPE"

	//將map序列化
	data, err := json.Marshal(a)

	if err != nil {
		fmt.Println("序列化失敗 err=", err)
	} else {
		//輸出序列化後的結果
		fmt.Printf("map序列化後=%v \n", string(data))
	}
}

func testSlice() {
	//切片是map[string]interface{}類型的
	var slice []map[string]interface{}

	var map1 = make(map[string]interface{})
	map1["name"] = "jack"
	map1["age"] = 7
	map1["address"] = "TPE"
	slice = append(slice, map1)

	var map2 = make(map[string]interface{})
	map2["name"] = "tom"
	map2["age"] = 20
	map2["address"] = []string{"NY", "SF"}
	slice = append(slice, map2)

	//將slice序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化失敗 err=", err)
	} else {
		//輸出序列化後的結果
		fmt.Printf("slice序列化後=%v \n", string(data))
	}
}

// 對基本類型序列化，對基本數據類型進行序列化意義不大
func testFloat64() {
	var num float64 = 8.8
	data, err := json.Marshal(num)
	if err != nil {
		fmt.Println("序列化失敗 err=", err)
	} else {
		//輸出序列化後的結果
		fmt.Printf("float64序列化後=%v \n", string(data))
	}
}

func main() {
	//將結構體、map、切片、基本類型序列化
	testStruct()
	testMap()
	testSlice()
	testFloat64()
}

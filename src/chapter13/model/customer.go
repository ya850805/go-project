package model

import "fmt"

// Customer 聲明一個Customer結構體，表示一個客戶信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// NewCustomer 編寫一個工廠模式，返回一個Customer實例
func NewCustomer(id int, name string, gender string, age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (c Customer) GetInfo() string {
	return fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", c.Id, c.Name, c.Gender, c.Age, c.Phone, c.Email)
}

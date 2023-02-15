package model

import "fmt"

type person struct {
	Name   string
	age    int
	salary float64
}

// NewPerson 工廠模式的函數，相當於構造函數
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年齡範圍不正確")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSalary(salary float64) {
	if salary >= 3000 && salary <= 30000 {
		p.salary = salary
	} else {
		fmt.Println("薪水範圍不正確")
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}

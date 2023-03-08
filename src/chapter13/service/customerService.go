package service

import "go-project/src/chapter13/model"

// CustomerService 完成對Customer的操作，包括增刪改查
type CustomerService struct {
	customers []model.Customer

	//聲明一個字段，表示當前切片含有多少個客戶
	//該字段還可以作為新客戶的id + 1
	customerNum int
}

// NewCustomerService 編寫一個方法可以返回*CustomerService
func NewCustomerService() *CustomerService {
	customerService := CustomerService{}
	customerService.customerNum = 1

	//為了顯示，我們初始化一個客戶
	customer := model.NewCustomer(1, "Jason", "M", 26, "999", "123@gmail.com")
	customerService.customers = append(customerService.customers, customer)

	return &customerService
}

// List 返回客戶切片
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

// Add 添加客戶
func (cs *CustomerService) Add(customer model.Customer) bool {
	//我們確定一個分配id的規則，就是添加的順序
	cs.customerNum++
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)

	return true
}

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

// Delete 根據id刪除客戶(從切片中刪除)
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)

	if index == -1 {
		return false
	} else {
		//從切片中刪除元素
		cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
		return true
	}
}

// Update 修改客戶信息
func (cs *CustomerService) Update(updateCustomer model.Customer) bool {
	index := cs.FindById(updateCustomer.Id)
	if index == -1 {
		return false
	} else {
		customer := &cs.customers[index]
		if updateCustomer.Name != "" {
			(*customer).Name = updateCustomer.Name
		}
		if updateCustomer.Gender != "" {
			(*customer).Gender = updateCustomer.Gender
		}
		if updateCustomer.Age != 0 {
			(*customer).Age = updateCustomer.Age
		}
		if updateCustomer.Phone != "" {
			(*customer).Phone = updateCustomer.Phone
		}
		if updateCustomer.Email != "" {
			(*customer).Email = updateCustomer.Email
		}
		return true
	}
}

// FindById 根據id查找客戶在切片中對應下標，如果沒有該客戶，返回-1
func (cs *CustomerService) FindById(id int) int {
	index := -1

	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			index = i
			break
		}
	}

	return index
}

package repository

import (
	"fmt"
	"github.com/myrachanto/power/httperors"
	"github.com/myrachanto/power/model"
	"github.com/myrachanto/power/support"
)
  
var (
	Customerrepo customerrepo = customerrepo{}
)

///curtesy to gorm
type customerrepo struct{}

func (customerRepo customerrepo) Create(customer *model.Customer) (*model.Customer, *httperors.HttpError) {
	if err := customer.Validate(); err != nil {
		return nil, err
	}
	ok := customer.ValidateEmail(customer.Email)
	if !ok {
		return nil, httperors.NewNotFoundError("Your email format is wrong!")
	}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Create(&customer)
	IndexRepo.DbClose(GormDB)
	return customer, nil
}
func (customerRepo customerrepo) GetOne(id int) (*model.Customer, *httperors.HttpError) {
	ok := customerRepo.customerUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("customer with that id does not exists!")
	}
	customer := model.Customer{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	fmt.Println(customer)
	GormDB.Model(&customer).Where("id = ?", id).First(&customer)
	IndexRepo.DbClose(GormDB)
	
	return &customer, nil
}
func (customerRepo customerrepo) GetAll(search *support.Search) ([]model.Customer,*httperors.HttpError) {
	// customers := []model.Customer{} 
	// results, err1 := customerRepo.Search(search, customers)
	// if err1 != nil {
	// 		return nil, err1
	// 	}
	return nil, nil
}

// func (customerRepo customerrepo) GetAll(search *support.Search) ([]interface{}, *httperors.HttpError) {
// 	customer := model.Customer{}
// 	// customers := []model.Customer{}
// 	// results, err1 := customerRepo.Search(search, customer)
// 	 results, err1 := support.SearchQuery(search, customer)
// 	if err1 != nil {
// 			return nil, err1
// 		}
// 	return results, nil 
// }

func (customerRepo customerrepo) Update(id int, customer *model.Customer) (*model.Customer, *httperors.HttpError) {
	ok := customerRepo.customerUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("customer with that id does not exists!")
	}
	
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	acustomer := model.Customer{}
	
	GormDB.Where("id = ?", id).First(&acustomer)
	if customer.Name  == "" {
		customer.Name = acustomer.Name
	}
	if customer.Company  == "" {
		customer.Company = acustomer.Company
	}
	if customer.Phone  == "" {
		customer.Phone = acustomer.Phone
	}
	if customer.Email  == "" {
		customer.Email = acustomer.Email
	}
	if customer.Address  == "" {
		customer.Address = acustomer.Address
	}
	GormDB.Model(&acustomer).Where("id = ?", id).Save(&customer)
	
	IndexRepo.DbClose(GormDB)

	return customer, nil
}
func (customerRepo customerrepo) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	ok := customerRepo.customerUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("customer with that id does not exists!")
	}
	customer := model.Customer{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	} 
	GormDB.Model(&customer).Where("id = ?", id).First(&customer)
	GormDB.Delete(customer)
	IndexRepo.DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (customerRepo customerrepo)customerUserExistByid(id int) bool {
	customer := model.Customer{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return false
	}	
	res := GormDB.First(&customer, "id =?", id)
	if res.Error != nil{
	   return false
	}
	IndexRepo.DbClose(GormDB)
	return true
	
}
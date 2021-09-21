package controllers

import(
	"fmt"
	"strconv"	
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/myrachanto/power/httperors"
	"github.com/myrachanto/power/model"
	"github.com/myrachanto/power/support"
	"github.com/myrachanto/power/service"
) 
  
var (
	CustomerController customerController = customerController{}
)
type customerController struct{ }
/////////controllers/////////////////
func (controller customerController) Create(c echo.Context) error {
	customer := &model.Customer{}
	if err := c.Bind(customer); err != nil {
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}
	fmt.Println(customer.Phone)
	createdcustomer, err1 := service.Customerservice.Create(customer)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, createdcustomer)
}
func (controller customerController) GetAll(c echo.Context) error {
	//  customers := []model.Customer{} 
	column := string(c.QueryParam("column"))
	direction := string(c.QueryParam("direction"))
	search_column := string(c.QueryParam("search_column"))
	search_operator := string(c.QueryParam("search_operator"))
	search_query_1 := string(c.QueryParam("search_query_1"))
	search_query_2 := string(c.QueryParam("search_query_2"))
	per_page, err := strconv.Atoi(c.QueryParam("per_page"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid per_page number")
		return c.JSON(httperror.Code, httperror)
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid page number")
		return c.JSON(httperror.Code, httperror)
	}
	fmt.Println("------------------------")
	search := &support.Search{Column:column, Direction:direction,Search_column:search_column,Search_operator:search_operator,Search_query_1:search_query_1,Search_query_2:search_query_2,Per_page:per_page,Page:page}
	Customers, err3 := service.Customerservice.GetAll(search)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, Customers)
} 
func (controller customerController) GetOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	fmt.Println(id)
	customer, problem := service.Customerservice.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, customer)	
}

func (controller customerController) Update(c echo.Context) error {
		
	customer :=  &model.Customer{}
	if err := c.Bind(customer); err != nil {
		httperror := httperors.NewBadRequestError("Invalid json body")
		return c.JSON(httperror.Code, httperror)
	}	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	updatedcustomer, problem := service.Customerservice.Update(id, customer)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, updatedcustomer)
}

func (controller customerController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	success, failure := service.Customerservice.Delete(id)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}

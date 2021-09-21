package controllers

import(
	"fmt"
	"strconv"	
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/myrachanto/power/httperors"
	"github.com/myrachanto/power/model"
	"github.com/myrachanto/power/service"
) 
//workspaceController ...
var (
	WorkspaceController workspaceController = workspaceController{}
)
type workspaceController struct{ }
/////////controllers/////////////////
func (controller workspaceController) Create(c echo.Context) error {
	workspace := &model.Workspace{}	
	workspace.Name = c.FormValue("name")
	workspace.Description = c.FormValue("description")
	workspace.Usercode = c.FormValue("usercode")
	createdworkspace, err1 := service.WorkspaceService.Create(workspace)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, createdworkspace)
}
func (controller workspaceController) Getlist(c echo.Context) error {
	code := c.Param("code")
	results, err3 := service.WorkspaceService.GetList(code)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, results)
} 
func (controller workspaceController) GetAll(c echo.Context) error {
	dated := c.QueryParam("dated")
	searchq2 := c.QueryParam("searchq2")
	searchq3 := c.QueryParam("searchq3")
	results, err3 := service.WorkspaceService.GetAll(dated,searchq2,searchq3)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, results)
} 
func (controller workspaceController) GetOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	fmt.Println(id)
	workspace, problem := service.WorkspaceService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, workspace)	
}

func (controller workspaceController) Update(c echo.Context) error {
		
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	updatedworkspace, problem := service.WorkspaceService.Update(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, updatedworkspace)
}

func (controller workspaceController) Delete(c echo.Context) error {
	code := c.Param("code")
	success, failure := service.WorkspaceService.Delete(code)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}

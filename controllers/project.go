package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/myrachanto/power/httperors"
	"github.com/myrachanto/power/model"
	"github.com/myrachanto/power/service"
)

//projectController ...
var (
	ProjectController projectController = projectController{}
)
type projectController struct{ }
/////////controllers/////////////////
func (controller projectController) Create(c echo.Context) error {
	project := &model.Project{}	
	project.Name = c.FormValue("name")
	project.Workspacecode = c.FormValue("workspacecode")
	project.Workspacename = c.FormValue("workspace")
	project.Description = c.FormValue("description")
	project.Usercode = c.FormValue("usercode")
	createdproject, err1 := service.ProjectService.Create(project)
	if err1 != nil {
		return c.JSON(err1.Code, err1)
	}
	return c.JSON(http.StatusCreated, createdproject)
}
func (controller projectController) GetAll(c echo.Context) error {
	dated := c.QueryParam("dated")
	searchq2 := c.QueryParam("searchq2")
	searchq3 := c.QueryParam("searchq3")
	results, err3 := service.ProjectService.GetAll(dated,searchq2,searchq3)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, results)
} 

func (controller projectController) Upload(c echo.Context) error {
	upload := &model.Projectupload{}
	upload.Project = c.FormValue("project")
	upload.Usercode = c.FormValue("usercode")
	upload.Workspace = c.FormValue("workspace")
	file, err2 := c.FormFile("file")
	// fmt.Println("66666666666666666666666666666", file)
//    fmt.Println(pic.Filename)
	if err2 != nil {
			httperror := httperors.NewBadRequestError("Invalid file")
			return c.JSON(httperror.Code, err2)
		}	
		src, err := file.Open()
		if err != nil {
			httperror := httperors.NewBadRequestError("the file is corrupted")
			return c.JSON(httperror.Code, err)
		}	
		defer src.Close()
		path := "public/workspaces/"+upload.Project
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.Mkdir(path, 0755)
		}
		filePath := path+"/" + file.Filename
		// filePath := "./public/workspaces/" + file.Filename
		// Destination
		dst, err4 := os.Create(filePath)
		if err4 != nil {
			httperror := httperors.NewBadRequestError("the Directory is mess")
			return c.JSON(httperror.Code, err4)
		}
	defer dst.Close()
	upload.Filename = filePath
	//  copy
	if _, err = io.Copy(dst, src); err != nil {
		if err2 != nil {
			httperror := httperors.NewBadRequestError("error filling")
			return c.JSON(httperror.Code, httperror)
		}
	} 
	
   
	results, err3 := service.ProjectService.Upload(upload)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, results)
} 

func (controller projectController) Getlist(c echo.Context) error {
	
	usercode := c.QueryParam("usercode")
	workspace := c.QueryParam("workspace")
	fmt.Println(">>>>>>>>>>>>", usercode,workspace)
	results, err3 := service.ProjectService.GetList(usercode,workspace)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, results)
} 

func (controller projectController) Getinfo(c echo.Context) error {
	
	usercode := c.QueryParam("usercode")
	project := c.QueryParam("project")
	fmt.Println(">>>>>>>>>>>>", usercode,project)
	results, err3 := service.ProjectService.Getinfo(usercode,project)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, results)
} 
func (controller projectController) Results(c echo.Context) error {
	
	usercode := c.QueryParam("usercode")
	project := c.QueryParam("project")
	fmt.Println(">>>>>>>>>>>>", usercode,project)
	results, err3 := service.ProjectService.Results(usercode,project)
	if err3 != nil {
		return c.JSON(err3.Code, err3)
	}
	return c.JSON(http.StatusOK, results)
} 
func (controller projectController) GetOne(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	fmt.Println(id)
	project, problem := service.ProjectService.GetOne(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, project)	
}

func (controller projectController) Update(c echo.Context) error {
		
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httperror := httperors.NewBadRequestError("Invalid ID")
		return c.JSON(httperror.Code, httperror)
	}
	updatedproject, problem := service.ProjectService.Update(id)
	if problem != nil {
		return c.JSON(problem.Code, problem)
	}
	return c.JSON(http.StatusOK, updatedproject)
}

func (controller projectController) Delete(c echo.Context) error {
	code := c.Param("code")
	success, failure := service.ProjectService.Delete(code)
	if failure != nil {
		return c.JSON(failure.Code, failure)
	}
	return c.JSON(success.Code, success)
		
}

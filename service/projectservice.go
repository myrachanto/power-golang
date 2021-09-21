package service

import (
	// "fmt"
	"github.com/myrachanto/power/httperors"
	"github.com/myrachanto/power/model"
	e "github.com/myrachanto/power/excel"
	r "github.com/myrachanto/power/repository"
)
//projectService ...
var (
	ProjectService projectService = projectService{}

) 
type projectService struct {
}

func (service projectService) Create(project *model.Project) (*model.Project, *httperors.HttpError) {
	if err := project.Validate(); err != nil {
		return nil, err
	}	
	project, err1 := r.Projectrepo.Create(project)
	if err1 != nil {
		return nil, err1
	}
	 return project, nil

}
func (service projectService) GetOne(id int) (*model.Project, *httperors.HttpError) {
	project, err1 := r.Projectrepo.GetOne(id)
	if err1 != nil {
		return nil, err1
	}
	return project, nil
}

func (service projectService) Upload(upload *model.Projectupload) (*model.Project, *httperors.HttpError) {
	e.Excelling(upload.Usercode,upload.Project, upload.Filename)
	results, err := r.Projectrepo.Upload(upload)
	return results, err
}
func (service projectService) GetList(usercode,workspace string) ([]model.Project, *httperors.HttpError) {
	results, err := r.Projectrepo.GetList(usercode,workspace)
	return results, err
}
func (service projectService) Getinfo(usercode,project string) (*model.InfoView, *httperors.HttpError) {
	results, err := r.Projectrepo.Getinfo(usercode,project)
	return results, err
}

func (service projectService) Results(usercode,project string) ([]model.Results, *httperors.HttpError) {
	results, err := r.Projectrepo.Results(usercode,project)
	return results, err
}
func (service projectService) GetAll(dated,searchq2,searchq3 string) ([]model.Project, *httperors.HttpError) {
	results, err := r.Projectrepo.GetAll(dated,searchq2,searchq3)
	return results, err
}
func (service projectService) Update(id int) (*model.Project, *httperors.HttpError) {
	project, err1 := r.Projectrepo.Update(id)
	if err1 != nil {
		return nil, err1
	}
	
	return project, nil
}
func (service projectService) Delete(code string) (*httperors.HttpSuccess, *httperors.HttpError) {
		success, failure := r.Projectrepo.Delete(code)
		return success, failure
}

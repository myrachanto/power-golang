package service

import (
	// "fmt"
	"github.com/myrachanto/power/httperors"
	"github.com/myrachanto/power/model"
	r "github.com/myrachanto/power/repository"
)
//workspaceService ...
var (
	WorkspaceService workspaceService = workspaceService{}

) 
type workspaceService struct {
}

func (service workspaceService) Create(workspace *model.Workspace) (*model.Workspace, *httperors.HttpError) {
	if err := workspace.Validate(); err != nil {
		return nil, err
	}	
	workspace, err1 := r.Workspacerepo.Create(workspace)
	if err1 != nil {
		return nil, err1
	}
	 return workspace, nil

}
func (service workspaceService) GetOne(id int) (*model.Workspace, *httperors.HttpError) {
	workspace, err1 := r.Workspacerepo.GetOne(id)
	if err1 != nil {
		return nil, err1
	}
	return workspace, nil
}
func (service workspaceService) GetList(code string) ([]model.Workspaceview, *httperors.HttpError) {
	results, err := r.Workspacerepo.GetList(code)
	return results, err
}
func (service workspaceService) GetAll(dated,searchq2,searchq3 string) ([]model.Workspace, *httperors.HttpError) {
	results, err := r.Workspacerepo.GetAll(dated,searchq2,searchq3)
	return results, err
}
func (service workspaceService) Update(id int) (*model.Workspace, *httperors.HttpError) {
	workspace, err1 := r.Workspacerepo.Update(id)
	if err1 != nil {
		return nil, err1
	}
	
	return workspace, nil
}
func (service workspaceService) Delete(code string) (*httperors.HttpSuccess, *httperors.HttpError) {
		success, failure := r.Workspacerepo.Delete(code)
		return success, failure
}

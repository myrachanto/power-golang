package model

import (
	"github.com/myrachanto/power/httperors"
	"gorm.io/gorm"
)

//Blog ...
type Workspace struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Quantity    int64  `json:"quantity"`
	Usercode    string `json:"usercode"`
	gorm.Model
}
type Workspaceview struct {
	Workspacename string    `json:"workspacename"`
	Workspacecode string    `json:"workspacecode"`
	Projects      []Project `json:"projects"`
}

//Validate ...
func (doc Workspace) Validate() *httperors.HttpError {

	if doc.Name == "" {
		return httperors.NewNotFoundError("Invalid Name")
	}
	if doc.Description == "" {
		return httperors.NewNotFoundError("Invalid Description")
	}
	return nil
}

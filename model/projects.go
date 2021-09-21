package model

import (
	"github.com/myrachanto/power/httperors"
	"gorm.io/gorm"
)

//Blog ...
type Project struct {
	Workspacecode string `json:"workspacecode"`
	Workspacename string `json:"workspacename"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Code          string `json:"code"`
	Quantity      string `json:"quantity"`
	Usercode      string `json:"usercode"`
	Uploaded      bool   `json:"uploaded"`
	Filename      string `json:"filename"`
	gorm.Model
}
type Info struct {
	Name        string `json:"name"`
	Usercode        string `json:"usercode"`
	Information string `json:"information"`
	gorm.Model
}
type InfoView struct {
	Header      []string `json:"headers"`
	Body      [][]string `json:"body"`
}
type Response struct {
	Name    string         `json:"name"`
	Results []string `json:"results"`
}
type Results struct {
	Name    string         `json:"name"`
	Results  []Content`json:"results"`
}
type Content struct {
	Name string `json:"name"`
	Quantity int `json:"quantity"`
}
type Body struct {
	Header      []string `json:"headers"`
}
type Projectupload struct {
	Workspace string `json:"workspace"`
	Project   string `json:"project"`
	Usercode  string `json:"usercode"`
	Filename  string `json:"filename"`
}

//Validate ...
func (doc Project) Validate() *httperors.HttpError {
	if doc.Name == "" {
		return httperors.NewNotFoundError("Invalid Name")
	}
	if doc.Description == "" {
		return httperors.NewNotFoundError("Invalid Description")
	}
	return nil
}

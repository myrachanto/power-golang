package repository

import (
	"strconv"
	"time"

	"github.com/myrachanto/power/httperors"
	"github.com/myrachanto/power/model"
)

//workspacerepo ...
var (
	Workspacerepo workspacerepo = workspacerepo{}
)

///curtesy to gorm
type workspacerepo struct {
}

func (m workspacerepo) Create(workspace *model.Workspace) (*model.Workspace, *httperors.HttpError) {
	if err := workspace.Validate(); err != nil {
		return nil, err
	}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	code1, x := m.GeneCode()
	if x != nil {
		return nil, x
	}
	workspace.Code = code1
	GormDB.Create(&workspace)
	IndexRepo.DbClose(GormDB)
	return workspace, nil
}
func (workspaceRepo workspacerepo) GetOne(id int) (*model.Workspace, *httperors.HttpError) {
	ok := workspaceRepo.WorksapceExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("workspace with that id does not exists!")
	}
	workspace := model.Workspace{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}

	GormDB.Model(&workspace).Where("id = ? ", id).First(&workspace)
	IndexRepo.DbClose(GormDB)

	return &workspace, nil
}
func (workspaceRepo workspacerepo) GetList(code string) ([]model.Workspaceview, *httperors.HttpError) {
	workspace := []model.Workspace{}
	projects := []model.Project{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
     work := model.Workspaceview{}
	GormDB.Where("usercode = ? ", code).Find(&workspace)
	works := []model.Workspaceview{}
	for _,k := range workspace {
		projs := []model.Project{}
		GormDB.Where("usercode = ? AND workspacecode = ?", code, k.Code).Find(&projects)
		work.Workspacename = k.Name
		work.Workspacecode = k.Code
		projs = append(projs , projects...)
		work.Projects = projs
		works = append(works, work)
	}
	IndexRepo.DbClose(GormDB)
	return works, nil
}
func (m workspacerepo) GetAll(dated, searchq2, searchq3 string) (results []model.Workspace, e *httperors.HttpError) {

	now := time.Now()
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	if dated != "custom" {
		if dated == "In the last 24hrs" {
			d := now.AddDate(0, 0, -1)
			GormDB.Where("updated_at > ?", d).Find(&results)
		}
		if dated == "In the last 7days" {
			d := now.AddDate(0, 0, -7)
			GormDB.Where("updated_at > ?", d).Find(&results)
		}
		if dated == "In the last 15day" {
			d := now.AddDate(0, 0, -15)
			GormDB.Where("updated_at > ?", d).Find(&results)
		}
		if dated == "In the last 30days" {
			d := now.AddDate(0, 0, -30)
			GormDB.Where("updated_at > ?", d).Find(&results)
		}
	}
	if dated == "custom" {
		start, err := time.Parse(Layout, searchq2)
		if err != nil {
			return nil, httperors.NewNotFoundError("Something went wrong parsing date1!")
		}
		end, err1 := time.Parse(Layout, searchq3)
		if err1 != nil {
			return nil, httperors.NewNotFoundError("Something went wrong parsing date1!")
		}
		GormDB.Where("updated_at BETWEEN ? AND ?", start, end).Find(&results)
	}
	IndexRepo.DbClose(GormDB)
	return results, nil
}
func (workspaceRepo workspacerepo) AllSearch(dbname, dated, searchq2, searchq3 string) (results []model.Workspace, r *httperors.HttpError) {

	now := time.Now()
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	if dated != "custom" {
		if dated == "In the last 24hrs" {
			d := now.AddDate(0, 0, -1)
			GormDB.Where("updated_at > ? AND shopalias = ?", d, dbname).Find(&results)
		}
		if dated == "In the last 7days" {
			d := now.AddDate(0, 0, -7)
			GormDB.Where("updated_at > ? AND shopalias = ?", d, dbname).Find(&results)
		}
		if dated == "In the last 15day" {
			d := now.AddDate(0, 0, -15)
			GormDB.Where("updated_at > ? AND shopalias = ?", d, dbname).Find(&results)
		}
		if dated == "In the last 30days" {
			d := now.AddDate(0, 0, -30)
			GormDB.Where("updated_at > ? AND shopalias = ?", d, dbname).Find(&results)
		}
	}
	if dated == "custom" {
		start, err := time.Parse(Layout, searchq2)
		if err != nil {
			return nil, httperors.NewNotFoundError("Something went wrong parsing date1!")
		}
		end, err1 := time.Parse(Layout, searchq3)
		if err1 != nil {
			return nil, httperors.NewNotFoundError("Something went wrong parsing date1!")
		}
		GormDB.Where("shopalias = ? AND updated_at BETWEEN ? AND ?", dbname, start, end).Find(&results)
	}
	IndexRepo.DbClose(GormDB)
	return results, nil

}
func (workspaceRepo workspacerepo) Update(id int) (*model.Workspace, *httperors.HttpError) {
	workspace := &model.Workspace{}
	ok := workspaceRepo.WorksapceExistByid( id)
	if !ok {
		return nil, httperors.NewNotFoundError("workspace with that id does not exists!")
	}

	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&workspace).Where("id = ?", id).Update("read", true)

	IndexRepo.DbClose(GormDB)

	return workspace, nil
}
func (workspaceRepo workspacerepo) Delete(code string) (*httperors.HttpSuccess, *httperors.HttpError) {
	ok := workspaceRepo.WorksapceBycode( code)
	if !ok {
		return nil, httperors.NewNotFoundError("Workspace with that code does not exists!")
	}
	workspace := model.Workspace{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Where("code = ?", code).Delete(&workspace)
	IndexRepo.DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (workspaceRepo workspacerepo) WorksapceBycode( code string) bool {
	workspace := model.Workspace{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return false
	}
	res := GormDB.Where("code = ?", code).First(&workspace)
	if res.Error != nil {
		return false
	}
	IndexRepo.DbClose(GormDB)
	return true

}
func (workspaceRepo workspacerepo) WorksapceExistByid( id int) bool {
	workspace := model.Workspace{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return false
	}
	res := GormDB.Where("id = ?", id).First(&workspace)
	if res.Error != nil {
		return false
	}
	IndexRepo.DbClose(GormDB)
	return true

}
func (workspaceRepo workspacerepo) WorksapceExistBycode( id string) *model.Workspace {
	workspace := model.Workspace{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil
	}
	res := GormDB.Where("code = ?", id).First(&workspace)
	if res.Error != nil {
		return nil
	}
	IndexRepo.DbClose(GormDB)
	return &workspace

}
func (workspaceRepo workspacerepo) WorksapceExistByname( id string) *model.Workspace {
	workspace := model.Workspace{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil
	}
	res := GormDB.Where("name = ?", id).First(&workspace)
	if res.Error != nil {
		return nil
	}
	IndexRepo.DbClose(GormDB)
	return &workspace

}

func (m workspacerepo) GeneCode() (string, *httperors.HttpError) {
	work := model.Workspace{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return "", err1
	}
	err := GormDB.Last(&work)
	if err.Error != nil {
		var c1 uint = 1
		code := "workspaceCode" + strconv.FormatUint(uint64(c1), 10)
		return code, nil
	}
	c1 := work.ID + 1
	code := "workspaceCode" + strconv.FormatUint(uint64(c1), 10)
	IndexRepo.DbClose(GormDB)
	return code, nil

}

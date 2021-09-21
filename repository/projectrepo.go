package repository

import (
	"fmt"
	"strconv"
	"time"
	"strings"
	"os"
	"encoding/csv"

	"github.com/myrachanto/power/httperors"
	"github.com/myrachanto/power/model"
)

//projectrepo ...
var (
	Projectrepo projectrepo = projectrepo{}
)

///curtesy to gorm
type projectrepo struct {
}

func (m projectrepo) Create(project *model.Project) (*model.Project, *httperors.HttpError) {
	if err := project.Validate(); err != nil {
		return nil, err
	}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	workspace := Workspacerepo.WorksapceExistBycode(project.Workspacecode)
	code1, x := m.GeneCode()
	if x != nil {
		return nil, x
	}
	fmt.Println("7777777777777777777", workspace)
	project.Workspacename = workspace.Name
	project.Code = code1
	GormDB.Create(&project)
	IndexRepo.DbClose(GormDB)
	return project, nil
}
func (m projectrepo) Createprojects(project *model.Project) (*model.Project, *httperors.HttpError) {
	if err := project.Validate(); err != nil {
		return nil, err
	}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	workspace := Workspacerepo.WorksapceExistBycode(project.Workspacecode)
	code1, x := m.GeneCode()
	if x != nil {
		return nil, x
	}
	fmt.Println("7777777777777777777", workspace)
	project.Workspacename = workspace.Name
	project.Code = code1
	GormDB.Create(&project)
	IndexRepo.DbClose(GormDB)
	return project, nil
}
func (p projectrepo) Upload(u *model.Projectupload) (*model.Project, *httperors.HttpError) {
	project := model.Project{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}

	GormDB.Where("usercode = ? AND workspacecode = ? AND code = ?", u.Usercode, u.Workspace, u.Project).First(&project)
	if project.Code == "" {
		return nil,  httperors.NewNotFoundError("That project does not exists!")
	}
	GormDB.Model(&project).Where("usercode = ? AND workspacecode = ? AND code = ?", u.Usercode, u.Workspace, u.Project).Updates(model.Project{Uploaded: true, Filename: u.Filename})
	IndexRepo.DbClose(GormDB)

	return &project, nil
}
func (p projectrepo) GetList(usercode,workspace string) ([]model.Project, *httperors.HttpError) {
	projects := []model.Project{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}

	GormDB.Where("usercode = ? AND workspacename = ?", usercode, workspace).Find(&projects)
	IndexRepo.DbClose(GormDB)

	return projects, nil
}
func (p projectrepo) Getinfo(usercode,project string) (*model.InfoView, *httperors.HttpError) {
	projects := []model.Info{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}

	GormDB.Where("name = ?", project).Find(&projects)
	// GormDB.Where("usercode = ? AND name = ?", usercode, project).Find(&projects)
	infoview := model.InfoView{}
	for k,v := range projects{
		if k == 0 {
			infoview.Header = strings.Split(v.Information, "|")
		}else{
			res := strings.Split(v.Information, "|")
			infoview.Body = append(infoview.Body, res)
		}

	}

	IndexRepo.DbClose(GormDB)

	return &infoview, nil
}
func (p projectrepo) Results(usercode,project string) ([]model.Results, *httperors.HttpError) {
	projects := []model.Info{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}

	GormDB.Where("name = ?", project).Find(&projects)
	// GormDB.Where("usercode = ? AND name = ?", usercode, project).Find(&projects)
	infoview := model.InfoView{}
	for k,v := range projects{
		if k == 0 {
			infoview.Header = strings.Split(v.Information, "|")
		}else{
			res := strings.Split(v.Information, "|")
			infoview.Body = append(infoview.Body, res)
		}

	}
	pj := model.Project{}
	GormDB.Where("code = ?", project).First(&pj)
	IndexRepo.DbClose(GormDB)
	// fmt.Println("bbbbbbbbbbb",infoview.Body[1][6],infoview.Body[2][6],infoview.Body[3][6],infoview.Body[4][6])
	// count := len(infoview.Body)/len(infoview.Header)
	// resa := []string{}
	// fmt.Println("bbbbbbbbbbbbbbbbbbbb", infoview.Body[1])
	resps := []model.Response{}
	for a,b := range infoview.Header{
		resp := model.Response{}
		fiti := []string{}
			for i := 0; i < len(infoview.Body); i++ {
				fiti = append(fiti, infoview.Body[i][a])
			}
			resp.Results = fiti
			resp.Name = b
			resps = append(resps, resp)
	}
	total := len(resps)
	// result1 := model.Results{}
	// content1 := model.Content{}
	// contents1 := []model.Content{}
	// // fmt.Println("bbbbbbbbbbb",resps[0], len(resps[0].Results), len(infoview.Body),len(infoview.Header))
	// fmt.Println("ffffffffffffffffff", resps[4])
	// dup_map := p.dup_count(resps[1].Results)
	// // fmt.Println("xxxxxxxxxxxx", dup_map)
	// // reso := make(map[string]int)
	// for x,y := range dup_map{
	// 	if y > 1 && y != total {
	// 		// fmt.Println("vvvvvvvvvvvvvvvvvvv", x, y)
	// 		content1.Name = x
	// 		content1.Quantity = y
	// 		contents1 = append(contents1, content1)
	// 	}
	// }
	// result1.Name = resps[1].Name
	// result1.Results = append(result1.Results, contents1...)
	// // fmt.Println("xxxxxxxxxxxx", result1)
	results := []model.Results{}
	for i := 0; i< len(infoview.Header); i++{
		result := model.Results{}
		contents := []model.Content{}
		content := model.Content{}
		dup_map := p.dup_count(resps[i].Results)
		for x,y := range dup_map{
			if y > 1 && y != total {
				// fmt.Println("vvvvvvvvvvvvvvvvvvv", x, y)
				content.Name = x
				content.Quantity = y
				contents = append(contents, content)
			}
		}
		if len(contents) > 1 &&  resps[i].Name != ""{
			result.Name = resps[i].Name
			result.Results = append(result.Results, contents...)
			results = append(results, result)
		}
	}
//    fmt.Println("vvvvvvvvvvvvvvvvvvv", results[0])


	return results, nil
}
func (p projectrepo) dup_count(list []string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}

func (projectRepo projectrepo) GetOne(id int) (*model.Project, *httperors.HttpError) {
	ok := projectRepo.ProductUserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("project with that id does not exists!")
	}
	project := model.Project{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}

	GormDB.Model(&project).Where("id = ? ", id).First(&project)
	IndexRepo.DbClose(GormDB)

	return &project, nil
}
func (m projectrepo) GetAll(dated, searchq2, searchq3 string) (results []model.Project, e *httperors.HttpError) {

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
func (projectRepo projectrepo) AllSearch(dbname, dated, searchq2, searchq3 string) (results []model.Project, r *httperors.HttpError) {

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
func (projectRepo projectrepo) Update(id int) (*model.Project, *httperors.HttpError) {
	project := &model.Project{}
	ok := projectRepo.ProductUserExistByid( id)
	if !ok {
		return nil, httperors.NewNotFoundError("project with that id does not exists!")
	}

	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&project).Where("id = ?", id).Update("read", true)

	IndexRepo.DbClose(GormDB)

	return project, nil
}
func (projectRepo projectrepo) Delete(code string) (*httperors.HttpSuccess, *httperors.HttpError) {
	ok := projectRepo.ProjectUserExistBycode( code)
	if !ok {
		return nil, httperors.NewNotFoundError("Project with that code does not exists!")
	}
	project := model.Project{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Where("code = ?", code).Delete(&project)
	IndexRepo.DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (projectRepo projectrepo) ProjectUserExistBycode( code string) bool {
	project := model.Project{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return false
	}
	res := GormDB.Where("code = ?", code).First(&project)
	if res.Error != nil {
		return false
	}
	IndexRepo.DbClose(GormDB)
	return true

}
func (projectRepo projectrepo) ProductUserExistByid( id int) bool {
	project := model.Project{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return false
	}
	res := GormDB.Where("id = ?", id).First(&project)
	if res.Error != nil {
		return false
	}
	IndexRepo.DbClose(GormDB)
	return true

}

func (m projectrepo) GeneCode() (string, *httperors.HttpError) {
	work := model.Project{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return "", err1
	}
	err := GormDB.Last(&work)
	if err.Error != nil {
		var c1 uint = 1
		code := "projectCode" + strconv.FormatUint(uint64(c1), 10)
		return code, nil
	}
	c1 := work.ID + 1
	code := "projectCode" + strconv.FormatUint(uint64(c1), 10)
	IndexRepo.DbClose(GormDB)
	return code, nil

}

// func (p projectrepo)chunkSlice(slice []string, chunkSize int) [][]int {
// 	var chunks [][]int
// 	var resp := []slice
// 	for a,b := range slice{
// 		end := a + chunkSize
// 		if end > len(slice) {
// 			end = len(slice)
// 		}
// 		chunks = append(chunks, b)
// 	}
// 	for i := 0; i < len(slice); i += chunkSize {
// 		end := i + chunkSize

// 		// necessary check to avoid slicing beyond
// 		// slice capacity
// 		if end > len(slice) {
// 			end = len(slice)
// 		}

// 		chunks = append(chunks, slice[i:end])
// 	}

// 	return chunks
// }
func (p projectrepo) analysedata(header,res []string) () {
	// for a,b := range 
	headcount := len(header)
	bodycount := len(res)
	counting := bodycount/headcount
	// results := [][]string{}
	fmt.Println("<<<<<<<<<<<<<<<<<", headcount,bodycount,counting)
	// for x,y := range res {

	// }

	for a,_ := range header {
		// a := 0
		// tg := a
		ress := []string{}
		// tg := a + counting
		for c,d := range res{
			headcount += headcount
			if c == a{
				ress = append(ress, d)
			}
			if c == a + headcount{
				ress = append(ress, d)
			}
		  }
		// results = append(results, ress)
		fmt.Println("<<<<<<<<<<<<<<<<<", ress)
	}
	// fmt.Println("<<<<<<<<<<<<<<<<<", results)
	return 

}
func (p projectrepo) mapdata(filo string) ([]string, *httperors.HttpError) {

	file, err := os.Open(filo)
    if err != nil {
		return nil, httperors.NewNotFoundError("That project does not exists!")
    }
    reader := csv.NewReader(file)
    records, _ := reader.ReadAll()
	header := []string{}
	body := [][]string{}
    for k,g := range records {
		bod := []string{}
		for _, l := range g {
			if k == 0 {
				header = append(header, l)
			}else{
				bod = append(bod, l)
			}
			body =append(body, bod)
		}

    }
	res := []string{}
	for _,b := range body{
		for c,_ := range header {
			for e,f := range b {
				if c == e {
					res = append(res, f)
				}
			}
		}
	}
    // fmt.Println("heading >>>>>>>>>>>>",head)
	return res, nil
}
	// g := p.dup_count(resa)

	// fmt.Println(">>>>>>>>>>>>>>",g)
	// results := make(map[string]int)
	// for m,n := range g {
	// 	if n > 1  && n != len(infoview.Body){
	// 		results[m] = n
	// 	}
	// }
	// results2 := make(map[string]int)
	// for m,n := range g {
	// 	if n > 1  && n != len(infoview.Body){
	// 		results[m] = n
	// 	}
	// }
	// charts := []model.Results{}
	// chart := model.Results{}
	// results2 := make(map[string]int)
	// // fmt.Println("aaaaaaaaaaaaaaaaaa", results)
	// for s,l := range infoview.Header {
	// 	// fmt.Println("111111111111111111111", s,l)
	// 	for g,f := range results {
	// 		// fmt.Println("22222222222222222222", g,f)
	// 		for _, v := range infoview.Body{					
	// 			for w, h := range v {
	// 				if s == w && h == g{
	// 					results2[g] = f
	// 				}
	// 			}

	// 		}
	// 		chart.Name = l
	// 		chart.Results = results2
	// 	}
	// 	// chart.Name = l
	// 	// chart.Results = results2
	// charts = append(charts, chart)
	// }

	// fmt.Println("wwwwwwwwwwwwwwwwwww", charts)
	
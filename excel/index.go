package excel

import (
	"encoding/csv"
	"fmt"
	"os"
	// "regexp"
	"strings"

	// "reflect"
	// "log"
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"

	"github.com/myrachanto/power/httperors"
	"github.com/myrachanto/power/model"
	r "github.com/myrachanto/power/repository"
	// "github.com/myrachanto/power/model"
	// "strings"
)
 
func Excelling(usercode,models, csvfile string)  *httperors.HttpError{
 
    file, err := os.Open(csvfile)
    if err != nil {
		return httperors.NewNotFoundError("Could not open the csv file")
    }
    reader := csv.NewReader(file)
    records, _ := reader.ReadAll()
    headers := records[0]
    head := []string{}
    for _, r := range headers {
        head = append(head, strings.Title(r))
    }
	GormDB, err1 := r.IndexRepo.Getconnected()
	if err1 != nil {
		return err1
	}    
    for _,g := range records {
        strs := getName(g)
        info := model.Info{}
        info.Name = models
        info.Usercode = usercode
        info.Information = strs
        GormDB.Create(&info)
    }
    fmt.Println("heading >>>>>>>>>>>>",head)
    // modeling := "./model/"+ models+ ".go"

	r.IndexRepo.DbClose(GormDB)
    return nil
}
func getName(params []string) string{
    info := ""
    for _,v := range params {
        info += v+"|"
    }
    return info
}
// func initstruct(head []string){
//     t := reflect.StructOf([]reflect.StructField{
//         for _,key := range head {
//             key1,_ := cleanser(key)
//             js :=  "`json:\""+ key1 +"\"`"
//             _, _ = fmt.Fprintln(f, key1," string",js)
//         }
//         {
//             Name: "Name",
//             Type: reflect.TypeOf(""), // string
//         },
//         {
//             Name: "Age",
//             Type: reflect.TypeOf(0), // int
//         },
//     })
    
//     fmt.Println(t)
    
//     v := reflect.New(t)
//     fmt.Printf("%+v\n", v)
//     v.Elem().FieldByName("Name").Set(reflect.ValueOf("Bob"))
//     v.Elem().FieldByName("Age").Set(reflect.ValueOf(12))
    
//     fmt.Printf("%+v\n", v)
// }
// func inserter(dbname string, names, values []string){

//     db, err := sql.Open("mysql", "user7:s$cret@tcp(127.0.0.1:3306)/testdb")
//     defer db.Close()

//     if err != nil {
//         log.Fatal(err)
//     }

//     CREATE TABLE "+dbname+"(id INT PRIMARY KEY AUTO_INCREMENT, name VARCHAR(255), population INT);
//     sql := "INSERT INTO "+dbname+"("+names+") VALUES ('Moscow', 12506000)"
//     res, err := db.Exec(sql)

//     if err != nil {
//         panic(err.Error())
//     }

//     lastId, err := res.LastInsertId()

//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Printf("The last inserted row id: %d\n", lastId)
// }
// func createmodel(modelname,files string,head []string) *httperors.HttpError{
//     f, err := os.Create(files)    
//     if err != nil {
//         httperors.NewNotFoundError("Could not create the file")
//     }   
//     fmt.Println("|||||||||||||||||||||||||||||") 
//     defer f.Close()   
//     _, _ = fmt.Fprintln(f,"package model")
//     _, _ = fmt.Fprintln(f,"")
//     _, _ = fmt.Fprintln(f,"import (")
//     _, _ = fmt.Fprintln(f,"\"gorm.io/gorm\"")
//     _, _ = fmt.Fprintln(f,")")
//     _, _ = fmt.Fprintln(f,"")
//     _, _ = fmt.Fprintln(f,"type ", modelname, " struct {")
//         for _,key := range head {
//             key1,_ := cleanser(key)
//             js :=  "`json:\""+ key1 +"\"`"
//             _, _ = fmt.Fprintln(f, key1," string",js)
//         }
//     _, _ = fmt.Fprintln(f,"gorm.Model")
//     _, err = fmt.Fprintln(f,"}")
//     if err != nil {
//     httperors.NewNotFoundError("Could not write into the file")
//     }
//     return nil
// }
// func cleanser(str string) (string, *httperors.HttpError){
//     reg, err := regexp.Compile("[^a-zA-Z]+")
//     if err != nil {
//         return "", httperors.NewNotFoundError("Could not write into the file")
//     }
//     st := reg.ReplaceAllString(str, "")
// 	return st, nil
// }
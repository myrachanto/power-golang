package repository

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	// "os"
	"strconv"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	// "github.com/joho/godotenv"
	"github.com/myrachanto/power/httperors"
	"github.com/myrachanto/power/model"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

//Userrepo ...
var (
	Userrepo userrepo = userrepo{}
)
type updatuser struct {
	Admin      string `json:"admin,omitempty"`
	Supervisor string `json:"supervisor,omitempty"`
	Employee   string `json:"employee,omitempty"`
}
type updatuser2 struct {
	Admin      bool `json:"admin,omitempty"`
	Supervisor bool `json:"supervisor,omitempty"`
	Employee   bool `json:"employee,omitempty"`
}
///curtesy to gorm
type userrepo struct{
	Bizname string `json:"bizname,omitempty"`
}
func (userRepo userrepo) Create(user *model.User) (string, *httperors.HttpError) {
	if err := user.Validate(); err != nil {
		return "", err
	}
	ok, err1 := user.ValidatePassword(user.Password)
	if !ok {
		return "", err1
	}
	ok = user.ValidateEmail(user.Email)
	if !ok {
		return "", httperors.NewNotFoundError("Your email format is wrong!")
	}
	ok = userRepo.UserExist(user.Email)
	if ok {
		return "", httperors.NewNotFoundError("Your email already exists!")
	}
	
	hashpassword, err2 := user.HashPassword(user.Password)
	if err2 != nil {
		return "", err2
	}
	user.Password = hashpassword
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return "", err1
	}
	code, x := userRepo.GeneCode()
	if x != nil {
		return "", x
	}
	user.Usercode = code
	// fmt.Println(user)
	// code, x := userRepo.GeneCode()
	// if x != nil {
	// 	return "", x
	// }
	// user.Usercode = code
	// fmt.Println(user)
	user.Admin = "admin"
	user.Employee = "employee"
	user.Supervisor = "supervisor"
	user.Accesslevel = "level3"
	// _, err := Businessrepo.Create(&model.Business{Name: user.Business})
	// if err != nil {
	// 	return "", err
	// }
	GormDB.Create(&user)
	IndexRepo.DbClose(GormDB)
	return "user created successifully", nil
}
func (userRepo userrepo) CreateUser(user *model.User) (string, *httperors.HttpError) {
	if err := user.Validate(); err != nil {
		return "", err
	}
	ok, err1 := user.ValidatePassword(user.Password)
	if !ok {
		return "", err1
	}
	ok = user.ValidateEmail(user.Email)
	if !ok {
		return "", httperors.NewNotFoundError("Your email format is wrong!")
	}
	ok = userRepo.UserExist(user.Email)
	if ok {
		return "", httperors.NewNotFoundError("Your email already exists!")
	}
	
	hashpassword, err2 := user.HashPassword(user.Password)
	if err2 != nil {
		return "", err2
	}
	user.Password = hashpassword
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return "", err1
	}
	code, x := userRepo.GeneCode()
	if x != nil {
		return "", x
	}
	user.Usercode = code
	// fmt.Println(user)
	// user.Admin = ""
	user.Employee = "employee"
	user.Supervisor = "supervisor"
	user.Accesslevel = "level3"
	// _, err := Businessrepo.Create(&model.Business{Name: user.Business})
	// if err != nil {
	// 	return "", err
	// }
	GormDB.Create(&user)
	IndexRepo.DbClose(GormDB)
	return "user created successifully", nil
}
func (userRepo userrepo) Login(auser *model.LoginUser) (*model.Auth, *httperors.HttpError) {
	if err := auser.Validate(); err != nil {
		return nil, err
	}
	///configure the db first and alias position
	//  := userRepo.Bizname
	sysDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	// shopsy := model.Shop{}
	// sysDB.Where("alias = ?", alias).First(&shopsy)
	// Systemadmin.DbClose(sysDB)
	//  := shopsy.
	// // fmt.Println("swwwwwwwwwwwwwwwww", shopsy)
	// GormDB, err1 := IndexRepo.Getconnected()
	// if err1 != nil {
	// 	return nil, err1
	// }
	// users := []model.User{}
	// GormDB.Find(&users)
	// fmt.Println("svvvvvvvvvvvvvvvvvvvvv", users)

	ok := userRepo.UserExist(auser.Email)
	if !ok {
		return nil, httperors.NewNotFoundError("Your email does not exists!")
	}
	
	user := model.User{}

	sysDB.Model(&user).Where("email = ?", auser.Email).First(&user)
	ok = user.Compare(auser.Password, user.Password)
	if !ok {
		return nil, httperors.NewNotFoundError("wrong email password combo!")
	}
	//get all the data into the token
	// shop := Businessrepo.BusinessExistByAlias(alias)
	tk := &model.Token{
		UserID: user.ID,
		UName: user.UName,
		Usercode:user.Usercode,
		Admin: user.Admin,
		Employee: user.Employee, 
		Supervisor: user.Supervisor,
		Role: user.Accesslevel,
		// Bizstatus: shop.Active,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: model.ExpiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	userkey, err := userRepo.Loaduserkey()
	if err != nil {
        log.Fatal("cannot load config:", err)
    }
	// encyKey := os.Getenv("EncryptionKey")
	tokenString, error := token.SignedString([]byte(userkey.EncryptionKey))
	if error != nil { 
		fmt.Println(error)
	}
	up := updatuser2{}
	up.Admin = true
	up.Supervisor = true
	up.Employee = true
	auth := &model.Auth{UserID:user.ID, UName:user.UName, Usercode:user.Usercode, Admin:user.Admin, Supervisor: user.Supervisor, Employee: user.Employee, Level: user.Accesslevel, Picture:user.Picture, Token:tokenString}
	IndexRepo.DbClose(sysDB)
	
	return auth, nil
}

func (userRepo userrepo) All( string) (t []model.User, r *httperors.HttpError) {

	user := model.User{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&user).Find(&t)
	IndexRepo.DbClose(GormDB)
	return t, nil

}
func (userRepo userrepo) Logout(token string) (*httperors.HttpError) {
	auth := model.Auth{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return err1
	}
	res := GormDB.First(&auth, "token =?", token)
	if res.Error != nil {
		return httperors.NewNotFoundError("Something went wrong logging out!")
	 }
	
	GormDB.Model(&auth).Where("token =?", token).First(&auth)
	
	GormDB.Delete(auth)
	IndexRepo.DbClose(GormDB)
	
	return  nil
}
func (userRepo userrepo)GeneCode() (string, *httperors.HttpError) {
	user := model.User{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return "", err1
	}
	// erre := godotenv.Load()
	// if erre != nil {
	// 	log.Fatal("Error loading .env file in routes")
	// }

	// key := os.Getenv("EncryptionKey")
	userkey, e := userRepo.Loaduserkey()
	if e != nil {
        log.Fatal("cannot load config:", e)
    }
	hash := sha256.New()
	err := GormDB.Last(&user)
	if err.Error != nil {
		var c1 uint = 1
		code1 := "UserCode"+strconv.FormatUint(uint64(c1), 10)+userkey.EncryptionKey
		// code = fmt.Sprintf( "%x",sha256.Sum256([]byte(code)))
		cod := strings.NewReader(code1)
		if _, er := io.Copy(hash, cod); er != nil {
		log.Fatal("Error hashing user")
		}
		code := fmt.Sprintf( "%x",hash.Sum(nil))
		return code, nil
	 }
	c1 := user.ID + 1
	code1 := "UserCode"+strconv.FormatUint(uint64(c1), 10)+userkey.EncryptionKey
	// code = fmt.Sprintf( "%x",sha256.Sum256([]byte(code)))cod := strings.NewReader(code1)
		cod := strings.NewReader(code1)
		if _, er := io.Copy(hash, cod); er != nil {
			log.Fatal("Error hashing user")
			}
			code := fmt.Sprintf( "%x",hash.Sum(nil))
	IndexRepo.DbClose(GormDB)
	return code, nil
	
}
func (userRepo userrepo)UserExistbycode(code string) bool {
	u := model.User{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return false
	}
	GormDB.Where("usercode = ?", code).First(&u)
	if u.ID == 0 {
	   return false
	}
	IndexRepo.DbClose(GormDB)
	return true
	
}
func (userRepo userrepo)Userbycode(code string) *model.User {
	u := model.User{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil
	}
	GormDB.Where("usercode = ?", code).First(&u)
	if u.ID == 0 {
	   return nil
	}
	IndexRepo.DbClose(GormDB)
	return &u
	
}
func (userRepo userrepo) GetAll(search string, page,pagesize int) ([]model.User, *httperors.HttpError) {
	results := []model.User{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	if search == ""{
		GormDB.Find(&results)
	}
	// db.Scopes(Paginate(r)).Find(&users)
	GormDB.Scopes(Paginate(page,pagesize)).Where("name LIKE ?", "%"+search+"%").Or("email LIKE ?", "%"+search+"%").Or("company LIKE ?", "%"+search+"%").Find(&results)

	IndexRepo.DbClose(GormDB)
	return results, nil
}

func (userRepo userrepo) UpdateRole(code,admin,supervisor,employee,level, usercode string) (string, *httperors.HttpError) {
	fmt.Println("vrrrrrrrrrrrrrrrrrrrrrrrrrr", code)
	user := model.User{}
	ok := Userrepo.UserExistbycode(code)
	if !ok {
		return "", httperors.NewNotFoundError("user with that code does not exists!")
	}
	
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return "", err1
	}

	adm, err := strconv.ParseBool(admin)
	if err != nil {
		return "", httperors.NewNotFoundError("Something went wrong parsing the admin!")
	}
	sup, err := strconv.ParseBool(supervisor)
	if err != nil {
		return "", httperors.NewNotFoundError("Something went wrong parsing the supervisor!")
	}
	// emp, err := strconv.ParseBool(employee)
	// if err != nil {
	// 	return "", httperors.NewNotFoundError("Something went wrong parsing the employee!")
	// }

	up := updatuser{}
	if adm == true {
		up.Admin = "admin"
		up.Supervisor = "supervisor"
		up.Employee = "employee"
		GormDB.Model(&user).Where("usercode = ?", code).Updates(model.User{Admin: up.Admin, Supervisor: up.Supervisor, Employee: up.Employee, Accesslevel:level})
		return "user updated succesifully", nil
	}
	if sup == true {
		up.Admin = "notadmin"
		up.Supervisor = "supervisor"
		up.Employee = "employee"
		GormDB.Model(&user).Where("usercode = ?", code).Updates(model.User{Admin: up.Admin, Supervisor: up.Supervisor, Employee: up.Employee, Accesslevel:level})
		return "user updated succesifully", nil
	}
	IndexRepo.DbClose(GormDB)

	up.Admin = "notadmin"
	up.Supervisor = "notsupervisor"
	up.Employee = "employee"
	GormDB.Model(&user).Where("usercode = ?", code).Updates(model.User{Admin: up.Admin, Supervisor: up.Supervisor, Employee: up.Employee, Accesslevel:level})
	return "user updated succesifully", nil
}
func (userRepo userrepo) Update(id int, user *model.User) (*model.User, *httperors.HttpError) {
	ok := userRepo.UserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("User with that id does not exists!")
	}
	
	hashpassword, err2 := user.HashPassword(user.Password)
	if err2 != nil {
		return nil, err2
	}
	user.Password = hashpassword
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	User := model.User{}
	uuser := model.User{}
	
	GormDB.Model(&User).Where("id = ?", id).First(&uuser)
	if user.FName  == "" {
		user.FName = uuser.FName
	}
	if user.LName  == "" {
		user.LName = uuser.LName
	}
	if user.UName  == "" {
		user.UName = uuser.UName
	}
	if user.Phone  == "" {
		user.Phone = uuser.Phone
	}
	if user.Address  == "" {
		user.Address = uuser.Address
	}
	if user.Picture  == "" {
		user.Picture = uuser.Picture
	}
	if user.Email  == "" {
		user.Email = uuser.Email
	}
	// if user.Admin  == false {
	// 	user.Admin = true
	// }
	GormDB.Save(&user)
	
	IndexRepo.DbClose(GormDB)

	return user, nil
}
func (userRepo userrepo) Delete(id int) (*httperors.HttpSuccess, *httperors.HttpError) {
	ok := userRepo.UserExistByid(id)
	if !ok {
		return nil, httperors.NewNotFoundError("User with that id does not exists!")
	}
	user := model.User{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil, err1
	}
	GormDB.Model(&user).Where("id = ?", id).First(&user)
	GormDB.Delete(user)
	IndexRepo.DbClose(GormDB)
	return httperors.NewSuccessMessage("deleted successfully"), nil
}
func (userRepo userrepo)UserExist(email string) bool {
	user := model.User{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return false
	}
	res := GormDB.Where("email = ?", email).First(&user)
	// res := GormDB.First(&user, "email =?", email)
	if res.Error != nil {
	   return false
	}
	IndexRepo.DbClose(GormDB)
	return true
	
}
func (userRepo userrepo)UserExistBycode(code string) (*model.User, *httperors.HttpError) {
	user := model.User{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return nil,err1
	}
	res := GormDB.Where("usercode =?", code).First(&user)
	if res.Error != nil {
	   return nil, httperors.NewBadRequestError("soemthing went wrong")
	}
	IndexRepo.DbClose(GormDB)
	return &user, nil
	
}
func (userRepo userrepo)UserExistByid(id int) bool {
	user := model.User{}
	GormDB, err1 := IndexRepo.Getconnected()
	if err1 != nil {
		return false
	}
	res := GormDB.First(&user, "id =?", id)
	if res.Error != nil{
	   return false
	}
	IndexRepo.DbClose(GormDB)
	return true
	
}
type Userkey struct {
	EncryptionKey string `mapstructure:"EncryptionKey"`
}

func (userRepo userrepo) Loaduserkey() (userkey Userkey, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&userkey)
	return
}
func Paginate(page, pagesize int) func(GormDB *gorm.DB) *gorm.DB {
	return func(GormDB *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		switch {
		case pagesize > 100:
			pagesize = 100
		case pagesize <= 0:
			pagesize = 10
		}

		offset := (page - 1) * pagesize
		return GormDB.Offset(offset).Limit(pagesize)
	}
}
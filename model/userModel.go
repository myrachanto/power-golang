package model

import (
	"github.com/jinzhu/gorm"
	"time"
	"golang.org/x/crypto/bcrypt"
	jwt "github.com/dgrijalva/jwt-go"
	"regexp"
	"github.com/myrachanto/power/httperors"
)
var ExpiresAt = time.Now().Add(time.Minute * 100000).Unix()

type User struct {
	FName       string     `json:"f_name"`
	LName       string     `json:"l_name"`
	UName       string     `json:"u_name"`
	Phone       string     `json:"phone"`
	Address     string     `json:"address"`
	Dob         *time.Time `json:"dob"`
	Picture     string     `json:"picture"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Admin       string     `json:"admin"`
	Employee    string     `json:"employee"`
	Supervisor  string     `json:"supervisor"`
	Accesslevel string     `json:"accesslevel"`
	Usercode    string     `json:"usercode"`
	gorm.Model
}
type Auth struct {
	//User User `gorm:"foreignKey:UserID; not null"`
	UserID      uint   `json:"userid"`
	UName       string `json:"uname"`
	Token       string `gorm:"size:500;not null" json:"token"`
	Usercode    string `json:"usercode"`
	Email       string `json:"email"`
	Admin       string `json:"admin"`
	Role        string `json:"role"`
	Employee    string `json:"employee"`
	Supervisor  string `json:"supervisor"`
	Accesslevel string `json:"accesslevel"`
	Level       string `json:"level"`
	Picture     string `json:"picture"`
	gorm.Model
}
type LoginUser struct {
	Email string `gorm:"not null"`
	Password string `gorm:"not null"`
}
//Token struct declaration
type Token struct {
	UserID      uint   `json:"user_id"`
	Usercode    string `json:"usercode"`
	UName       string `json:"uname"`
	Email       string `json:"email"`
	Admin       string `json:"admin"`
	Role        string `json:"role"`
	Employee    string `json:"employee"`
	Supervisor  string `json:"supervisor"`
	Accesslevel string `json:"accesslevel"`
	*jwt.StandardClaims
}
func (user User)ValidateEmail(email string) (matchedString bool) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&amp;'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	matchedString = re.MatchString(email)
	return matchedString
}
func (user User)ValidatePassword(password string) (bool, *httperors.HttpError) {
	if len(password) < 5 {
		return false, httperors.NewBadRequestError("your password need more characters!")
	} else if len(password) > 32 {
		return false, httperors.NewBadRequestError("your password is way too long!")
	}
	return true, nil
}
func (user User)HashPassword(password string)(string, *httperors.HttpError){
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		if err != nil {
			return "", httperors.NewNotFoundError("type a stronger password!")
		}
		return string(pass),nil 
		
	}
func (user User) Compare(p1,p2 string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(p2), []byte(p1))
	if err != nil {
		return false
	}
	return true
}
func (loginuser LoginUser) Validate() *httperors.HttpError{ 
	if loginuser.Email == "" {
		return httperors.NewNotFoundError("Invalid Email")
	}
	if loginuser.Password == "" {
		return httperors.NewNotFoundError("Invalid password")
	}
	return nil
}
func (user User) Validate() *httperors.HttpError{
	if user.FName == "" {
		return httperors.NewNotFoundError("Invalid first Name")
	}
	if user.LName == "" {
		return httperors.NewNotFoundError("Invalid last name")
	}
	if user.UName == "" {
		return httperors.NewNotFoundError("Invalid username")
	}
	if user.Phone == "" {
		return httperors.NewNotFoundError("Invalid phone number")
	}
	if user.Email == "" {
		return httperors.NewNotFoundError("Invalid Email")
	}
	if user.Address == "" {
		return httperors.NewNotFoundError("Invalid Address")
	}
	if user.Password == "" {
		return httperors.NewNotFoundError("Invalid password")
	}
	// if user.Picture == "" {
	// 	return httperors.NewNotFoundError("Invalid picture")
	// }
	if user.Email == "" {
		return httperors.NewNotFoundError("Invalid picture")
	}
	return nil
}
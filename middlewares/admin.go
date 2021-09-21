package middlewares

import (
	"fmt"
	// "os"
	// "fmt"
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)
type Userkey struct{
	EncryptionKey string `mapstructure:"EncryptionKey"`
}
func Loaduserkey() (userkey Userkey, err error) {
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
//IsAdmin middleware evalutes if the user is admin - super admin
func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	userkey, err := Loaduserkey()
	if err != nil {
        log.Fatal("cannot load config:", err)
    }
	return func(c echo.Context) error {
		headertoken := c.Request().Header.Get("Authorization")
		token := strings.Split(headertoken, " ")[1]
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token)(interface{}, error){
			return []byte(userkey.EncryptionKey), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "unable to parse token")
		}
		// fmt.Println("fc", fc)
		// level := claims["role"].(string)
		admin := claims["admin"].(string)
		supervisor := claims["supervisor"].(string)
		employee := claims["employee"].(string)
		fmt.Println("role", admin)
		if admin != "admin" {
			return echo.NewHTTPError(http.StatusForbidden, "unable to parse token")
		}
		if supervisor != "supervisor"{
			return echo.NewHTTPError(http.StatusForbidden, "unable to parse token")
		}
		if employee != "employee" {
			return echo.NewHTTPError(http.StatusForbidden, "unable to parse token")
		}
		return next(c)
	}
}
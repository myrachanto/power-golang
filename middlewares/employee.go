package middlewares

import (
	// "os"
	// "fmt"
	"log"
	"net/http"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)
//IsEmployee middleware evalutes if the user is admin - super admin
func IsEmployee(next echo.HandlerFunc) echo.HandlerFunc {
	
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
		// level := claims["role"].(string)
		admin := claims["admin"].(string)
		supervisor := claims["supervisor"].(string)
		employee := claims["employee"].(string)
		// fmt.Println("role", isAdmin)
		if admin != "notadmin" {
			return echo.NewHTTPError(http.StatusForbidden, "unable to parse token")
		}
		if supervisor != "notsupervisor"{
			return echo.NewHTTPError(http.StatusForbidden, "unable to parse token")
		}
		if employee != "employee" {
			return echo.NewHTTPError(http.StatusForbidden, "unable to parse token")
		}
		return next(c)
	}
}
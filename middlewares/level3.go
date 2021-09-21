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
//IsAdmin middleware evalutes if the user is admin - super admin
func Level3(next echo.HandlerFunc) echo.HandlerFunc {
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
		level := claims["role"].(string)
		if level != "level3" {
			return echo.NewHTTPError(http.StatusForbidden, "unable to parse token")
		}
		return next(c)
	}
}
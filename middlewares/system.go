package middlewares

import (
	// "fmt"
	// "os"
	// "fmt"
	"log"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

//IsAdmin middleware evalutes if the user is admin - super admin
func System(next echo.HandlerFunc) echo.HandlerFunc {
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
		bizname := claims["bizname"].(string)
		central := claims["central"].(string)
		if bizname == "" {
			return echo.NewHTTPError(http.StatusForbidden, "unable to parse token and get alias")
		}
		//ensuring the context has the db variable to all routes
		c.Set("bizname", bizname)
		c.Set("central", central)
		return next(c)
	}
}
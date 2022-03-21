package auth

import (
	"brostools-api-person/lib/log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

//Authentication with HTTP-header secret
func Secret() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			objReturn := new(ResponseData)
			objErrors := new(ResponseErrors)
			objReturn.Data = make([]string, 0)
			secret := c.Request().Header.Get("secret")
			if secret == "" {
				log.Infof(nil, "Get secret failed: secret is empty")
				objErrors.Code = http.StatusUnauthorized
				objErrors.Message = "Please provide secret credentials"
				objReturn.Errors = append(objReturn.Errors, objErrors)
				return c.JSON(http.StatusUnauthorized, objReturn)
			}

			if secret != os.Getenv("API_USE_SECRET_BROSTOOLS_API_JWT") {
				log.Infof(nil, "Wrong secret credentials: "+secret)
				objErrors.Code = http.StatusUnauthorized
				objErrors.Message = "Wrong secret credentials"
				objReturn.Errors = append(objReturn.Errors, objErrors)
				return c.JSON(http.StatusUnauthorized, objReturn)
			}

			return next(c)
		}
	}
}

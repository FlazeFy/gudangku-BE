package middlewares

import (
	"fmt"
	"gudangku/modules/auth/models"
	"gudangku/modules/auth/repositories"
	"gudangku/packages/helpers/auth"
	"gudangku/packages/helpers/generator"
	"gudangku/packages/helpers/response"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func CustomJWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte("secret"),
		})(next)(c)

		if err != nil {
			if errJWT, res := err.(*jwt.ValidationError); res {
				if errJWT.Errors&jwt.ValidationErrorExpired != 0 {
					var res response.Response
					res.Status = http.StatusUnauthorized
					res.Message = "Your access is expired, please try sign in again"
					return c.JSON(http.StatusUnauthorized, res)
				}
			}

			var res response.Response
			res.Status = http.StatusUnauthorized
			res.Message = "Token invalid. Please check again"
			return c.JSON(http.StatusUnauthorized, res)
		}

		return nil
	}
}

func CheckLogin(c echo.Context, body models.UserLogin) (response.Response, error) {
	var res response.Response
	timeExpStr, err := strconv.Atoi(auth.GetJWTConfiguration("exp"))
	if err != nil {
		res.Status = http.StatusInternalServerError
		return res, err
	}

	duration := time.Duration(timeExpStr) * time.Second

	id, err, ctx := repositories.PostUserAuth(body.Username, body.Password)
	// Response
	if err == nil && id == "" {
		res.Status = http.StatusUnprocessableEntity
		res.Message = ctx
		return res, err
	}

	if err != nil {
		res.Status = http.StatusUnauthorized
		res.Message = ctx
		return res, err
	}

	if id == "" {
		res.Status = http.StatusUnauthorized
		res.Message = ctx
		return res, err
	}

	// Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = body.Username
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * duration).Unix()
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = err.Error()
	}

	// DB Token
	var obj models.UserToken
	obj.ContextType = "user" // For now
	obj.ContextId = id
	obj.Token = fmt.Sprintf("%s", t)

	errAccs := repositories.PostAccessToken(obj)
	if errAccs != nil {
		res.Status = http.StatusInternalServerError
		res.Message = ctx
		return res, err
	}

	// Response
	res.Status = http.StatusOK
	res.Message = generator.GenerateCommandMsg("login", "", 1)
	res.Data = map[string]string{
		"token": t,
	}

	return res, nil
}

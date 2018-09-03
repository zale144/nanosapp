package service

import (
	"fmt"
	"net/http"
	"time"
	"log"
	"github.com/zale144/nanosapp/services/web/commons"
	"github.com/dchest/authcookie"
	"github.com/labstack/echo"
	"strings"
	"github.com/zale144/nanosapp/services/web/client"
	"github.com/dgrijalva/jwt-go"
	"github.com/zale144/instagram-bot/services/api/model"
)

type AccountService interface {
	Login(c echo.Context) error
}

// Login handles login requests by requesting 'session'
// service to log into Instagram and save the session to cache.
// It also requests 'api' service to create a JWT token,
// for 'api' authorization
func Login(c echo.Context) error {

	username, password, ok := c.Request().BasicAuth()

	// No Authentication header
	if ok != true {
		err := fmt.Errorf("bad auth credentials")
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}
	if username == "" || password == "" {
		return echo.ErrUnauthorized
	}
	account, err := client.AccountClient{}.Get(username)
	if err != nil {
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}
	if account == nil {
		err := fmt.Errorf("the account with provided username does not exist")
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}

	passwordHash := commons.CryptPrivate(password, commons.CRYPT_SETTING)

	if passwordHash != account.Password {
		err := fmt.Errorf("wrong password for your account")
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}

	token, err := LoginApi(username)
	if err != nil {
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}
	// get the session cookie
	cookie := &http.Cookie{
		Name:  commons.CookieName,
		Value: authcookie.NewSinceNow(username, 24*time.Hour, []byte(commons.SECRET)),
		Path:  "/",
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func LoginApi(username string) (string, error) {
	claims := &model.JwtCustomClaims{
		Name: username,
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(model.SECRET))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return t, nil
}

// Logout handles logout requests. It expires the cookie and
// logs the user out of Instagram by calling the 'session' service.
func Logout(c echo.Context) error {
	// expire the cookie
	cookie := &http.Cookie{
		Name:    commons.CookieName,
		Expires: time.Now(),
		Path:    "/",
	}
	c.SetCookie(cookie)

	return c.Redirect(http.StatusSeeOther, "/login")
}
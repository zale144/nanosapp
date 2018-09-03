package service

import (
	"fmt"
	"net/http"
	"time"
	"log"

	"github.com/zale144/nanosapp/services/web/commons"
	"github.com/dchest/authcookie"
	"github.com/labstack/echo"
	"github.com/zale144/nanosapp/services/web/client"
	"github.com/dgrijalva/jwt-go"
	"github.com/zale144/instagram-bot/services/api/model"
)

type AccountService struct {}

// Login handles login requests
func (as AccountService) Login(c echo.Context) error {

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

	token, err := loginApi(username)
	if err != nil {
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}
	// get the session cookie
	cookie := &http.Cookie{
		Name:  commons.CookieName,
		Value: authcookie.NewSinceNow(username, 24 * time.Hour, []byte(commons.SECRET)),
		Path:  "/",
	}
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

func loginApi(username string) (string, error) {
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

// Register registers a new account
func (as AccountService) Register(c echo.Context) error {

	username, password := c.Param("username"), c.Param("password")

	if username == "" {
		err := fmt.Errorf("Username is required")
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}
	if password == "" {
		err := fmt.Errorf("Password is required")
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}
	password = commons.CryptPrivate(password, commons.CRYPT_SETTING)

	accountResponse, err := client.AccountClient{}.Add(username, password)
	if err != nil {
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}
	if accountResponse == nil {
		err := fmt.Errorf("Error while registering account")
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}

	cookie := &http.Cookie{
		Name:  model.CookieName,
		Value: authcookie.NewSinceNow(accountResponse.Account, 24*time.Hour, []byte(model.SECRET)),
		Path:  "/",
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusCreated, "Created")
}

// Logout handles logout requests. It expires the cookie and
// logs the user out of Instagram by calling the 'session' service.
func (as AccountService) Logout(c echo.Context) error {
	// expire the cookie
	cookie := &http.Cookie{
		Name:    commons.CookieName,
		Expires: time.Now(),
		Path:    "/",
	}
	c.SetCookie(cookie)

	return c.Redirect(http.StatusSeeOther, "/login")
}
package service

import (
	"fmt"
	"log"
	"time"
	"net/http"

	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"github.com/dchest/authcookie"
	"github.com/zale144/nanosapp/services/web/client"
	"github.com/zale144/nanosapp/services/web/commons"
	"github.com/zale144/nanosapp/services/web/model"
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
	if err != nil || account == nil {
		err := fmt.Errorf("the account with provided username does not exist")
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}

	if !commons.PortableHashCheck(password, account.Password) {
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

// loginApi creates a signed JWT token for
// accessing the api endpoints
func loginApi(username string) (string, error) {
	claims := &commons.JwtCustomClaims{
		Name: username,
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(commons.SECRET))
	if err != nil {
		log.Println(err)
		return "", err
	}
	return t, nil
}

// Register handles requests to register a new account
func (as AccountService) Register(c echo.Context) error {

	acc := new(model.Account)           //initialize  struct Account
	if err := c.Bind(acc); err != nil { //get and bind data from request to struct Account
		err := fmt.Errorf("Invalid JSON payload")
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}

	if acc.Username == "" {
		err := fmt.Errorf("Username is required")
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}
	if acc.Password == "" {
		err := fmt.Errorf("Password is required")
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}
	if acc.Password != acc.ConfirmPassword {
		err := fmt.Errorf("Passwords don't match")
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}
	// encrypt the password
	acc.Password = commons.CryptPrivate(acc.Password, commons.CRYPT_SETTING)
	// use the account microservice client to add a new account to it's db
	accountResponse, err := client.AccountClient{}.Add(acc.Username, acc.Password)
	if err != nil {
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}
	if accountResponse == nil {
		err := fmt.Errorf("Error while registering account")
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}

	// login to api
	token, err := loginApi(acc.Username)
	if err != nil {
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}

	// login to web
	cookie := &http.Cookie{
		Name:  commons.CookieName,
		Value: authcookie.NewSinceNow(accountResponse.Username, 24 * time.Hour, []byte(commons.SECRET)),
		Path:  "/",
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

// Logout handles logout requests. It expires the cookie
// and redirects the user to the login page
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
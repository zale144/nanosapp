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
	_, err := client.AccountClient{}.Get(username, password)
	if err != nil {
		c.Error(echo.NewHTTPError(http.StatusBadRequest, err.Error()))
		return err
	}

	token, err := client.Api{}.Login(username)
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

	user, err := GetUsernameFromCookie(&c)
	if err == nil {
		_, err := client.AccountClient{}.Logout(user)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println(err)
	}
	return c.Redirect(http.StatusSeeOther, "/login")
}

// GetUsernameFromCookie gets the username from the cookie
func GetUsernameFromCookie(cp *echo.Context) (string, error) {
	c := *cp
	headers := c.Request().Header
	cookieStr := headers.Get("cookie")
	if cookieStr == "" {
		err := fmt.Errorf("empty cookie")
		return "", err
	}
	value := strings.Replace(cookieStr, commons.CookieName+"=", "", -1)
	username := authcookie.Login(value, []byte(commons.SECRET))
	if username == "" {
		err := fmt.Errorf("no user authenticated")
		return "", err
	}
	return username, nil
}
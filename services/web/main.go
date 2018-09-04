package main

import (
	"os"
	"io"
	"fmt"
	"log"
	"net/http"
	"html/template"

	"github.com/labstack/echo"
	"github.com/dchest/authcookie"
	"github.com/labstack/echo/middleware"
	"github.com/zale144/nanosapp/services/web/commons"
	"github.com/zale144/nanosapp/services/web/service"
	"github.com/zale144/nanosapp/services/web/client"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
)

func main() {
	e := echo.New()

	t := &wTemplate{
		templates: template.Must(template.ParseGlob("public/templates/*.html")),
	}
	e.Static("/static", "public/static")
	e.Renderer = t

	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// ***************** public ***************************
	e.GET("/login", func(c echo.Context) error {
		return c.File("public/static/html/login.html")
	})
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "/admin/home")
	})
	e.POST("/login", service.AccountService{}.Login)
	e.GET("/logout", service.AccountService{}.Logout)
	e.POST("/register", service.AccountService{}.Register)

	// ***************** private ***************************
	a := e.Group("/admin")
	a.Use(authMiddleware)
	a.GET("/home", func(c echo.Context) error {
		data := map[string]interface{}{
			"ApiURL":   commons.ApiURL,
		}
		return c.Render(http.StatusOK, "home", data)
	})

	commons.ApiURL = os.Getenv("API_HOST")

	api := e.Group("/api/v1")
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &commons.JwtCustomClaims{},
		SigningKey: []byte(commons.SECRET),
	}
	api.Use(middleware.JWTWithConfig(config))

	api.GET("/ad-campaigns", service.AdCampaignService{}.GetAll)

	go reqService()
	e.Logger.Fatal(e.Start(":8081"))
}

// reqService registers the 'web' microservice
func reqService()  {
	commons.Service = k8s.NewService(
		micro.Name("web"),
		micro.Version("latest"),
	)
	commons.Service.Init()

	if err := commons.Service.Run(); err != nil {
		log.Fatal(err)
	}
}

type wTemplate struct {
	templates *template.Template
}

func (t *wTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// authMiddleware is used to check if user is logged in
func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(commons.CookieName)
		if err == nil {
			login := authcookie.Login(cookie.Value, []byte(commons.SECRET))
			if login == "" {
				return c.Redirect(http.StatusTemporaryRedirect, "/login")
			}
			acc, err := client.AccountClient{}.Get(login)
			if err != nil {
				service.AccountService{}.Logout(c)
				return c.Redirect(http.StatusTemporaryRedirect, "/login")
			}
			fmt.Println(acc)
			c.Request().Header.Set(commons.HEADER_AUTH_USER_ID, login)
			return next(c)
		}
		log.Println(err)
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}

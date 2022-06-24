package endpoints

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pablodz/sherlockgo/internal/endpoints/sites"
	"github.com/pablodz/sherlockgo/internal/endpoints/username"
	"gorm.io/gorm"
)

func HandleRequest(db *gorm.DB) {
	e := echo.New()

	/* Add here the middlewares */
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	/* Add here the routes or endpoints */
	e.GET("/", GETsimpleResponse())
	e.GET("/sites", sites.GETListSites(db))
	e.GET("/username/:username", username.GETByUsername(db))
	e.Logger.Fatal(e.Start("0.0.0.0")) // :8080
}

func GETsimpleResponse() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, I'm Sherlock, but faster!")
	}
}

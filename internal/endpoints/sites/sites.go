package sites

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/pablodz/sherlockgo/internal/models"
)

func GETListSites(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var paths []models.Sites
		db.Find(&paths)
		return c.JSON(http.StatusOK, paths)
	}
}

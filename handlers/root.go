package handlers

import (
	"fmt"
	"net/http"

	"github.com/kimihito/tohatebu/models"
	"github.com/kimihito/tohatebu/services"
	"github.com/labstack/echo/v4"
)

// Root hoge
func Root(c echo.Context) (err error) {
	e := new(models.Entry)
	if err := c.Bind(e); err != nil {
		return err
	}

	b, err := services.AddBookmark(e)

	if err != nil {
		return err
	}

	fmt.Println(b)

	return c.JSON(http.StatusOK, b)
}

package handler

import (
	"net/http"

	"brostools-api-person/lib/log"

	"github.com/labstack/echo/v4"
)

type IndexHandler interface {
	HandleIndex() echo.HandlerFunc
}

type indexHandler struct{}

func NewIndexHandler() IndexHandler {
	return &indexHandler{}
}

func (ih *indexHandler) HandleIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		log.Infof(ctx, "IndexHandler.HandleIndex")
		type ResIndex struct {
			Errors   []string `json:"errors"`
			Messages string   `json:"messages"`
		}

		resIndex := &ResIndex{
			Errors:   []string{},
			Messages: "",
		}

		resIndex.Messages = "Hello, everyone. brostools-api-person!"

		return c.JSON(http.StatusOK, resIndex)
	}
}

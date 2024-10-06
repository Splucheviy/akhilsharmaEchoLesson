package handlers

import (
	"net/http"
	"strconv"

	"github.com/Splucheviy/akhilsharmaEchoLesson/cmd/api/service"
	"github.com/labstack/echo/v4"
)

func PostIndexHandler(c echo.Context) error {
	data, err := service.GetAll()
	if err != nil {
		return c.String(http.StatusBadGateway, "Unable to process data")
	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid id")
	}
	data, err := service.GetById(idx)
	if err != nil {
		return c.String(http.StatusNotFound, "Post not found")
	}

	res := make(map[string]any)
	res["status"] = "OK"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}

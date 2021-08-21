package controller

import (
	"log"

	"github.com/Drakkar-Software/Errors-Server/api/dao"
	"github.com/Drakkar-Software/Errors-Server/api/model"

	"net/http"

	"github.com/labstack/echo/v4"
)


// PostError registers a bot as started (creates a new bot if necessary)
func PostError(c echo.Context) error {
	inputError := new(model.Error)
	_ = c.Bind(inputError)
	id, err := dao.SaveError(inputError)
	if err != nil {
		log.Println(err, inputError.Timestamp)
		return c.JSON(http.StatusBadRequest, id)
	}
	return c.JSON(http.StatusOK, id)
}

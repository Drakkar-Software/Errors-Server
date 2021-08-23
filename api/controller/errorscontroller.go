package controller

import (
	"fmt"
	"log"

	"github.com/Drakkar-Software/Errors-Server/api/dao"
	"github.com/Drakkar-Software/Errors-Server/api/model"

	"net/http"

	"github.com/labstack/echo/v4"
)


// PostError registers a bot as started (creates a new bot if necessary)
func PostError(c echo.Context) error {
	inputErrors := new([]model.Error)
	_ = c.Bind(inputErrors)
	ids := make([]string, len(*inputErrors), len(*inputErrors))
	var err error = nil
	for index, inputError := range *inputErrors {
		id, err := dao.SaveError(inputError)
		ids[index] = fmt.Sprintf("%s", id)
		if err != nil {
			log.Println(err, inputError.Error.Title, inputError.Error.Timestamp)
		}
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, ids)
	}
	return c.JSON(http.StatusOK, ids)
}

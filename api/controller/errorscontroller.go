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
	inputErrorData := new([]model.ErrorData)
	_ = c.Bind(inputErrorData)
	ids := make([]string, len(*inputErrorData), len(*inputErrorData))
	var err error = nil
	for index, inputErrorDatum := range *inputErrorData {
		inputError := model.Error{
			Error: inputErrorDatum,
		}
		id, err := dao.SaveError(inputError)
		ids[index] = fmt.Sprintf("%s", id)
		if err != nil {
			log.Println(err, inputError.Error.Title, inputError.Error.FirstTimestamp)
		}
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, ids)
	}
	return c.JSON(http.StatusOK, ids)
}

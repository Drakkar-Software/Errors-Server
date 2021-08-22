package dao

import (
	"context"
	errorPackage "github.com/Drakkar-Software/Errors-Server/api/model"
	"log"

	"github.com/Drakkar-Software/Errors-Server/database"
)


// SaveError saves an error into the database
func SaveError(uploadedError *errorPackage.Error) (interface{}, error) {
	newID, err := database.Database.Collection.InsertOne(context.Background(), uploadedError)
	if err != nil {
		return nil, err
	}
	log.Println("Registered new Error \"", uploadedError.Error.Title, "\" with id:", newID.InsertedID)
	return newID.InsertedID, nil

}

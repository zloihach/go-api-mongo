package query

import (
	"go-api-mongo/modules/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type GoAppDb struct {
	App config.GoAppTools
	DB  mongo.Client
}

//func NewGoAppDB(app config.GoAppTools, db mongo.Client) database.DBRepo{
//	return
//}

package infrastruct

import (
	"gopkg.in/mgo.v2"
)

func NewSqliteHandler(connString string) *MongoHandler {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	mongoHandler := new(MongoHandler)
	sqliteHandler.Session = session
	return mongoHandler
}

type MongoHandler struct {
	Session *mgo.Session
}

func (handler *MongoHandler) GetAll() ([]*map[string]interface{}, error) {

}

func (handler *MongoHandler) GetByTitle(title string) (map[string]interface{}, error) {

}

func (handler *MongoHandler) Save(article map[string]interface{}) error {

}

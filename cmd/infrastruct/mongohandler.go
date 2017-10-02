package infrastruct

import (
	"gopkg.in/mgo.v2"
)

type MongoHandler struct {
	Session *mgo.Session
}

func (handler *MongoHandler) GetAll() ([]*map[string]interface{}, error) {

}

func (handler *MongoHandler) GetByTitle(title string) (map[string]interface{}, error) {

}

func (handler *MongoHandler) Save(article map[string]interface{}) error {

}

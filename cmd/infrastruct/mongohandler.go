package infrastruct

import (
	"encoding/json"
	"log"

	"github.com/4406arthur/clean-micro-article-service/pkg/entity"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func NewMongoHandler(connString string) *MongoHandler {
	session, err := mgo.Dial(connString)
	if err != nil {
		panic(err)
	}
	mongoHandler := new(MongoHandler)
	mongoHandler.Session = session
	return mongoHandler
}

type MongoHandler struct {
	Session *mgo.Session
}

func (handler *MongoHandler) GetAll() ([]map[string]interface{}, error) {
	session := handler.Session.Copy()
	defer session.Close()

	c := session.DB("blog").C("articles")

	var collection []map[string]interface{}
	err := c.Find(bson.M{}).All(&collection)
	if err != nil {
		log.Println("Failed get all articles: ", err)
		return nil, err
	}

	return collection, nil
}

func (handler *MongoHandler) GetByTitle(title string) (map[string]interface{}, error) {
	session := handler.Session.Copy()
	defer session.Close()

	c := session.DB("blog").C("articles")

	var collection map[string]interface{}
	err := c.Find(bson.M{"title": title}).One(&collection)
	if err != nil {
		log.Println("Failed get article by title: ", err)
		return nil, err
	}

	return collection, nil
}

func (handler *MongoHandler) Save(doc interface{}) error {
	session := handler.Session.Copy()
	defer session.Close()

	c := session.DB("blog").C("articles")

	var article entity.Article

	buf, err := json.Marshal(doc)
	json.Unmarshal(buf, &article)

	if err != nil {
		log.Fatal(err)
		return err
	}

	err = c.Insert(article)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

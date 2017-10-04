package infrastruct

import (
	"gopkg.in/mgo.v2"
	"log"
	"github.com/4406arthur/clean-micro-article-service/cmd/interfaces"
	
)

func NewSqliteHandler(connString string) *MongoHandler {
	session, err := mgo.Dial(connString)
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
	session := handler.Session.Copy()
	defer session.Close()

	c := session.DB("blog").C("articles")

	var article
	err := c.Find(bson.M{"title": title}).One(&article)
	if err != nil {
		ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		log.Println("Failed find book: ", err)
		return
	}

	if article.title == "" {
		ErrorWithJSON(w, "article not found", http.StatusNotFound)
		return
	}

	respBody, err := json.MarshalIndent(article, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return respBody
}

func (handler *MongoHandler) Save(article map[string]interface{}) error {

}

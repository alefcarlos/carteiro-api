package repo

import mgo "gopkg.in/mgo.v2"

//MongoSession armazena a sessao de conexao com o Mongo
var MongoSession *mgo.Session

//NewMongoSession faz a conexao com o Mongo
func NewMongoSession() (err error) {
	err = nil
	MongoSession, err = mgo.Dial("mongodb://localhost:27017/")
	return
}

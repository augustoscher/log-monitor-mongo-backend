package dao

import (
	"log"

	"bitbucket.org/augustoscher/logs-monitor-api/config"
	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database
var c = config.Config{}

//Connect estabelece conexão com banco
func Connect() {
	c.Read()

	session, err := mgo.Dial(c.Server)
	if err != nil {
		log.Printf("Parametros: %+v", c)
		log.Fatal("Erro de conexão com banco: ", err)
	}
	db = session.DB(c.Database)
}

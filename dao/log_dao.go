package dao

import (
	"fmt"
	"log"

	"bitbucket.org/augustoscher/logs-monitor-api/model"
	"gopkg.in/mgo.v2/bson"
)

//LogDAO represents dao logs
type LogDAO struct {
}

const (
	//COLLECTION name of database document
	COLLECTION = "logmessage"
)

//FindAll return all logs
func (dao *LogDAO) FindAll() ([]model.LogMessage, error) {
	var logs []model.LogMessage
	err := db.C(COLLECTION).Find(bson.M{}).All(&logs)
	return logs, err
}

//FindAllGrouped return all logs - grouped by codigo_integracao
func (dao *LogDAO) FindAllGrouped() ([]model.LogGroup, error) {
	query := []bson.M{
		{"$group": bson.M{"_id": "$nome_integracao", "quantidade": bson.M{"$sum": 1}}},
		{"$sort": bson.M{"codigo_integracao": 1}},
	}

	pipeline := db.C(COLLECTION).Pipe(query)
	var logs []model.LogGroup
	err := pipeline.All(&logs)
	if err != nil {
		log.Print(err)
	}
	fmt.Println("result:", logs)
	return logs, err
}

//FindByID busca por id
func (dao *LogDAO) FindByID(id string) (model.LogMessage, error) {
	var logs model.LogMessage
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&logs)
	return logs, err
}

// Insert adiciona novo
func (dao *LogDAO) Insert(log model.LogMessage) error {
	err := db.C(COLLECTION).Insert(&log)
	return err
}

// Delete deleta um registro
func (dao *LogDAO) Delete(log model.LogMessage) error {
	err := db.C(COLLECTION).Remove(&log)
	return err
}

package controller

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/augustoscher/logs-monitor-api/dao"
	"bitbucket.org/augustoscher/logs-monitor-api/model"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var logDAO = dao.LogDAO{}

//AllLogsEndPoint return all logs
func AllLogsEndPoint(w http.ResponseWriter, r *http.Request) {
	logs, err := logDAO.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, logs)
}

//FindLogEndpoint find a log by id
func FindLogEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log, err := logDAO.FindByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Log ID")
		return
	}
	respondWithJSON(w, http.StatusOK, log)
}

//CreateLogEndPoint create new log
func CreateLogEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var log model.LogMessage
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	log.ID = bson.NewObjectId()
	if err := logDAO.Insert(log); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, log)
}

//DeleteLogEndPoint delete log by id
func DeleteLogEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var log model.LogMessage
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := logDAO.Delete(log); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

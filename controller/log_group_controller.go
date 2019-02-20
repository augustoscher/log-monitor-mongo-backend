package controller

import (
	"net/http"

	"bitbucket.org/augustoscher/logs-monitor-api/dao"
)

var logGroupDAO = dao.LogDAO{}

//AllLogsGroupEndPoint return grouped logs
func AllLogsGroupEndPoint(w http.ResponseWriter, r *http.Request) {
	logs, err := logGroupDAO.FindAllGrouped()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, logs)
}

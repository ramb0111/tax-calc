package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ramb0111/tax-calc/src"
	"github.com/ramb0111/tax-calc/src/services"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.ParseInt(vars["userId"], 10, 64)
	if err != nil {
		src.RespondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	decoder := json.NewDecoder(r.Body)
	var pl src.Payload
	err = decoder.Decode(&pl)
	if err != nil {
		src.RespondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	if _, err = services.Insert(userId, &pl); err != nil {
		src.RespondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	src.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, err := strconv.ParseInt(vars["userId"], 10, 64)
	if err != nil {
		src.RespondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	result := services.GetAll(userId)
	src.RespondWithJSON(w, http.StatusOK, result)
}

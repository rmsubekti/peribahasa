package controllers

import (
	"encoding/json"
	"net/http"
	"peribahasa/app/models"
	"peribahasa/app/utils"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateAsal controller
var CreateAsal = func(w http.ResponseWriter, r *http.Request) {
	asal := &models.Asal{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(asal)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request payload"))
		return
	}

	if e := asal.Create(); e != nil {
		utils.Respond(w, utils.Message(false, e.Error()))
		return
	}

	resp := utils.Message(true, "Success")
	resp["data"] = asal
	utils.Respond(w, resp)
}

// GetAsal Controller
var GetAsal = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c := &models.Asal{}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	if err := c.Get(id); err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	resp := utils.Message(true, "Success")
	resp["data"] = c
	utils.Respond(w, resp)
}

// UpdateAsal Controller
var UpdateAsal = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	asal := &models.Asal{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(asal)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request payload"))
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	if err := asal.Update(id); err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	asal.ID = uint(id)
	resp := utils.Message(true, "Success")
	resp["data"] = asal
	utils.Respond(w, resp)
}

// DeleteAsal Controller
var DeleteAsal = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	asal := &models.Asal{}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	if err := asal.Delete(id); err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	resp := utils.Message(true, "Deleted")
	utils.Respond(w, resp)
}

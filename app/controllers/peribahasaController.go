package controllers

import (
	"encoding/json"
	"net/http"
	"peribahasa/app/models"
	"peribahasa/app/utils"
	"strconv"

	"github.com/gorilla/mux"
)

// CreatePeribahasa controller
var CreatePeribahasa = func(w http.ResponseWriter, r *http.Request) {
	peribahasa := &models.Peribahasa{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(peribahasa)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request payload"))
		return
	}

	if e := peribahasa.Create(); e != nil {
		utils.Respond(w, utils.Message(false, e.Error()))
		return
	}

	resp := utils.Message(true, "Success")
	resp["data"] = peribahasa
	utils.Respond(w, resp)
}

// GetPeribahasa Controller
var GetPeribahasa = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c := &models.Peribahasa{}

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

}// GetRandomPeribahasa Controller
var GetRandomPeribahasa = func(w http.ResponseWriter, r *http.Request) {
	p := &models.Peribahasa{}

	if err := p.Get(-1); err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	resp := utils.Message(true, "Success")
	resp["data"] = p
	utils.Respond(w, resp)
}

// UpdatePeribahasa Controller
var UpdatePeribahasa = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	peribahasa := &models.Peribahasa{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(peribahasa)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request payload"))
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	if err := peribahasa.Update(id); err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	peribahasa.ID = uint(id)
	resp := utils.Message(true, "Success")
	resp["data"] = peribahasa
	utils.Respond(w, resp)
}

// DeletePeribahasa Controller
var DeletePeribahasa = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	peribahasa := &models.Peribahasa{}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	if err := peribahasa.Delete(id); err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	resp := utils.Message(true, "Deleted")
	utils.Respond(w, resp)
}

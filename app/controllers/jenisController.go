package controllers

import (
	"encoding/json"
	"net/http"
	"peribahasa/app/models"
	"peribahasa/app/utils"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateJenis controller
var CreateJenis = func(w http.ResponseWriter, r *http.Request) {
	category := &models.Jenis{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(category)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request payload"))
		return
	}

	if e := category.Create(); e != nil {
		utils.Respond(w, utils.Message(false, e.Error()))
		return
	}

	resp := utils.Message(true, "Success")
	resp["data"] = category
	utils.Respond(w, resp)
}

// GetJenis Controller
var GetJenis = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	c := &models.Jenis{}

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

// UpdateJenis Controller
var UpdateJenis = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := &models.Jenis{}

	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(category)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request payload"))
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	if err := category.Update(id); err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	category.ID = uint(id)
	resp := utils.Message(true, "Success")
	resp["data"] = category
	utils.Respond(w, resp)
}

// DeleteJenis Controller
var DeleteJenis = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := &models.Jenis{}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	if err := category.Delete(id); err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}

	resp := utils.Message(true, "Deleted")
	utils.Respond(w, resp)
}

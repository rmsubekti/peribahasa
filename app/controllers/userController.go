package controllers

import (
	"encoding/json"
	"net/http"
	"peribahasa/app/models"
	"peribahasa/app/utils"
)

// Register new user
var Register = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "invalid request:"+err.Error()))
		return
	}

	if err := user.Create(); err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}
	respond := utils.Message(true, "User Created")
	respond["data"] = user
	utils.Respond(w, respond)
}

// Login user
var Login = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}
	if err := user.Login(user.Email, user.Password); err != nil {
		utils.Respond(w, utils.Message(false, err.Error()))
		return
	}
	respond := utils.Message(true, "Logged in")
	respond["data"] = user
	utils.Respond(w, respond)
}

// // Me logged in user
// var Me = func(w http.ResponseWriter, r *http.Request) {
// 	resp := models.GetUser(app.GetUserID(r))
// 	utils.Respond(w, resp)
// }

package handlers

import (
	"net/http"
	"peribahasa/app/models"
	"peribahasa/web/utils"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// EditAsalPage controller
var EditAsalPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = utils.ParseTemplates("web/templates/admin/asal/edit.html")
	vars := mux.Vars(r)
	a := &models.Asal{}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := a.Get(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		a.Bahasa = r.FormValue("bahasa")
		if err := a.Update(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/asal", http.StatusSeeOther)
		return
	}

	if err = tmpl.ExecuteTemplate(w, "base", map[string]interface{}{"data": a, "status": false}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// AddNewAsalPage controller
var AddNewAsalPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = utils.ParseTemplates("web/templates/admin/asal/new.html")

	if r.Method == http.MethodPost {
		asal := models.Asal{
			Bahasa: r.FormValue("bahasa"),
		}
		if err := asal.Create(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/asal", http.StatusSeeOther)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteAsalPage controller
var DeleteAsalPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = utils.ParseTemplates("web/templates/admin/asal/delete.html")
	vars := mux.Vars(r)
	a := &models.Asal{}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		if strings.ToLower(r.FormValue("confirm")) == "ya" {
			if err := a.Delete(id); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/admin/asal", http.StatusSeeOther)
		return
	}

	if err := a.Get(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "base", map[string]interface{}{"data": a}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// AsalIndexPage controller
var AsalIndexPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = utils.ParseTemplates("web/templates/admin/asal/index.html")
	var lAsal models.ListAsal

	if err := lAsal.List(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", map[string]interface{}{"data": lAsal}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

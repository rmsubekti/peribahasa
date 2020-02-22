package handlers

import (
	"net/http"
	"peribahasa/app/models"
	"peribahasa/web/utils"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// EditJenisPage controller
var EditJenisPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = utils.ParseTemplates("web/templates/admin/jenis/edit.html")
	vars := mux.Vars(r)
	j := &models.Jenis{}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := j.Get(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		j.Nama = r.FormValue("nama")
		if err := j.Update(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/jenis", http.StatusSeeOther)
		return
	}

	if err = tmpl.ExecuteTemplate(w, "base", map[string]interface{}{"data": j, "status": false}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// AddNewJenisPage controller
var AddNewJenisPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = utils.ParseTemplates("web/templates/admin/jenis/new.html")

	if r.Method == http.MethodPost {
		jenis := models.Jenis{
			Nama: r.FormValue("nama"),
		}
		if err := jenis.Create(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/jenis", http.StatusSeeOther)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeleteJenisPage controller
var DeleteJenisPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = utils.ParseTemplates("web/templates/admin/jenis/delete.html")
	vars := mux.Vars(r)
	j := &models.Jenis{}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		if strings.ToLower(r.FormValue("confirm")) == "ya" {
			if err := j.Delete(id); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/admin/jenis", http.StatusSeeOther)
		return
	}

	if err := j.Get(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "base", map[string]interface{}{"data": j}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// JenisIndexPage controller
var JenisIndexPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = utils.ParseTemplates("web/templates/admin/jenis/index.html")
	var lJenis models.ListJenis

	if err := lJenis.List(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", map[string]interface{}{"data": lJenis}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

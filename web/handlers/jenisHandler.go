package handlers

import (
	"html/template"
	"net/http"
	"peribahasa/app/models"
	"peribahasa/web/utils"
	"strconv"

	"github.com/gorilla/mux"
)

// JenisPage controller
var JenisPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"web/templates/partial/header.html",
		"web/templates/partial/footer.html",
		"web/templates/admin/jenis/index.html",
		"web/templates/base.html",
	))

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

	if r.Method == http.MethodGet {
		if err = tmpl.ExecuteTemplate(w, "base", map[string]interface{}{"data": j}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// EditJenisPage controller
var EditJenisPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"web/templates/partial/header.html",
		"web/templates/partial/footer.html",
		"web/templates/admin/jenis/edit.html",
		"web/templates/base.html",
	))

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
		jenis := models.Jenis{
			Nama: r.FormValue("nama"),
		}

		if err := jenis.Update(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jenis.ID = uint(id)
		if err = tmpl.ExecuteTemplate(w, "base", map[string]interface{}{"data": jenis, "status": true}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	if r.Method == http.MethodGet {
		if err = tmpl.ExecuteTemplate(w, "base", map[string]interface{}{"data": j, "status": false}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// AddNewJenisPage controller
var AddNewJenisPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles(
		"web/templates/partial/header.html",
		"web/templates/partial/footer.html",
		"web/templates/admin/jenis/new.html",
		"web/templates/base.html",
	))

	if r.Method == http.MethodGet {
		if err := tmpl.ExecuteTemplate(w, "base", map[string]interface{}{"data": nil, "status": false}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	jenis := models.Jenis{
		Nama: r.FormValue("name"),
	}

	if err := jenis.Create(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", map[string]interface{}{"data": jenis, "status": true}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// JenisListPage controller
var JenisListPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = utils.ParseTemplates("web/templates/admin/jenis/list.html")
	var lJenis models.ListJenis
	if r.Method == http.MethodPost {
		jenis := models.Jenis{
			Nama: r.FormValue("nama"),
		}

		if err := jenis.Create(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/jenis", http.StatusSeeOther)
	}
	if err := lJenis.List(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", map[string]interface{}{"data": lJenis}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

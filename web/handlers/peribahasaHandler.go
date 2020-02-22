package handlers

import (
	"net/http"
	"peribahasa/app/models"
	"peribahasa/web/utils"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// EditPeribahasaPage controller
var EditPeribahasaPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = utils.ParseTemplates("web/templates/admin/peribahasa/edit.html")
	vars := mux.Vars(r)
	p := &models.Peribahasa{}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		p.TeksAsli = r.FormValue("asli")
		p.Arti = r.FormValue("arti")

		asalID, aerr := strconv.Atoi(r.FormValue("asal"))
		if aerr != nil {
			http.Error(w, aerr.Error(), http.StatusInternalServerError)
			return
		}
		p.IDAsal = uint(asalID)

		jenisID, jerr := strconv.Atoi(r.FormValue("jenis"))
		if jerr != nil {
			http.Error(w, jerr.Error(), http.StatusInternalServerError)
			return
		}
		p.IDJenis = uint(jenisID)

		if err := p.Update(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/peribahasa", http.StatusSeeOther)
		return
	}

	if err := p.Get(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var jenis models.ListJenis
	if err := jenis.List(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var asal models.ListAsal
	if err := asal.List(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.ExecuteTemplate(w, "base", map[string]interface{}{
		"peribahasa": p,
		"asal":       asal,
		"jenis":      jenis,
		"status":     false,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// AddNewPeribahasaPage controller
var AddNewPeribahasaPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = utils.ParseTemplates("web/templates/admin/peribahasa/new.html")

	if r.Method == http.MethodPost {
		asalID, aerr := strconv.Atoi(r.FormValue("asal"))
		if aerr != nil {
			http.Error(w, aerr.Error(), http.StatusInternalServerError)
			return
		}

		jenisID, jerr := strconv.Atoi(r.FormValue("jenis"))
		if jerr != nil {
			http.Error(w, jerr.Error(), http.StatusInternalServerError)
			return
		}

		peribahasa := models.Peribahasa{
			TeksAsli: r.FormValue("asli"),
			Arti:     r.FormValue("arti"),
			IDAsal:   uint(asalID),
			IDJenis:  uint(jenisID),
		}
		if err := peribahasa.Create(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/admin/peribahasa", http.StatusSeeOther)
		return
	}

	var jenis models.ListJenis
	if err := jenis.List(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var asal models.ListAsal
	if err := asal.List(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "base", map[string]interface{}{
		"asal":  asal,
		"jenis": jenis,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DeletePeribahasaPage controller
var DeletePeribahasaPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = utils.ParseTemplates("web/templates/admin/peribahasa/delete.html")
	vars := mux.Vars(r)
	p := &models.Peribahasa{}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		if strings.ToLower(r.FormValue("confirm")) == "ya" {
			if err := p.Delete(id); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		http.Redirect(w, r, "/admin/peribahasa", http.StatusSeeOther)
		return
	}

	if err := p.Get(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.ExecuteTemplate(w, "base", map[string]interface{}{"data": p}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// PeribahasaIndexPage controller
var PeribahasaIndexPage = func(w http.ResponseWriter, r *http.Request) {
	var tmpl = utils.ParseTemplates("web/templates/admin/peribahasa/index.html")
	var lPeribahasa models.ListPeribahasa

	if err := lPeribahasa.List(0, 100); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "base", map[string]interface{}{"data": lPeribahasa}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

package utils

import (
	"html/template"
)

//ParseTemplates : parse all template
func ParseTemplates(path string) *template.Template {
	return template.Must(template.ParseFiles(
		"web/templates/partial/head.html",
		"web/templates/partial/header.html",
		"web/templates/partial/footer.html",
		path,
		"web/templates/base.html",
	))
}

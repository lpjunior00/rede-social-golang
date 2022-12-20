package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// Insere os templates da pasta views dentro da variavel de template
func LoadTemplates() {
	//Adiciono os templates principais
	templates = template.Must(template.ParseGlob("views/*.html"))
	//Adiciona os templates reutilizaveis (notar que a variavel é templates, porque referencia a de cima)
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// Renderiza uma pagina html na tela
func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}

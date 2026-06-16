package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// CarregarTamplates insere os templates html na variavel templates
func CarregarTamplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecutarTemplate renderiza uma pagina html na tela
func ExecutarTemplate(w http.ResponseWriter, tamplate string, dados interface{}) {
	templates.ExecuteTemplate(w, tamplate, dados)
}

package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func semantic(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "semantic.html", nil)
}

func portfolio(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "portfolio.html", nil)
}

func flexbox(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "flexbox.html", nil)
}

func gohtml(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusNoContent)
		return
	}

	fName := r.FormValue("firstName")
	lName := r.FormValue("lastName")

	d := struct {
		First string
		Last  string
	}{
		First: fName,
		Last:  lName,
	}
	tpl.ExecuteTemplate(w, "nextpage.html", d)
}
func main() {

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", handler)
	http.HandleFunc("/next", gohtml)
	http.HandleFunc("/semantic", semantic)
	http.HandleFunc("/portfolio", portfolio)
	http.HandleFunc("/flexbox", flexbox)
	http.ListenAndServe(":8000", nil)
}

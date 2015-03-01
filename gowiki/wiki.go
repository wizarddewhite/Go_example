package main

import (
	"fmt"
	"html/template"
	"net/http"
)

import (
	"github.com/wizarddewhite/Go_example/gowiki/wikipage"
)

func renderTemplate(w http.ResponseWriter, tmpl string, p *wikipage.Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	fmt.Println("View page: ", title)
	p, err := wikipage.LoadPage(title)
	if err != nil {
		fmt.Println("No such page: %s, redirect to /edit/", title)
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := wikipage.LoadPage(title)
	if err != nil {
		p = &wikipage.Page{Title: title}
	}
	if len(p.Title) == 0 {
		fmt.Fprint(w, "<h1>Please specify your page name</h1>")
	} else {
		fmt.Println("Edit page: ", p.Title)
		renderTemplate(w, "edit", p)
	}
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.ListenAndServe(":8080", nil)
}

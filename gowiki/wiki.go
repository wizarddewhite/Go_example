package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
)

import (
	"github.com/wizarddewhite/Go_example/gowiki/wikipage"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *wikipage.Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
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
	title, err := getTitle(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
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

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		fmt.Println(err)
		return
	}
	body := r.FormValue("body")
	p := &wikipage.Page{Title: title, Body: []byte(body)}
	fmt.Println("Save page: ", title)
	err = p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}

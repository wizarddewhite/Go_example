package main

import (
	"fmt"
	"html/template"
	"net/http"
)

import (
	"github.com/wizarddewhite/Go_example/gowiki/wikipage"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	fmt.Println("Get page: ", title)
	p, err := wikipage.LoadPage(title)
	if err != nil {
		fmt.Fprintf(w, "No such page: %s", title)
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
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
		t, _ := template.ParseFiles("edit.html")
		t.Execute(w, p)
	}
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.ListenAndServe(":8080", nil)
}

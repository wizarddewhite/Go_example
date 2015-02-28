package main

import (
	"fmt"
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

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.ListenAndServe(":8080", nil)
}

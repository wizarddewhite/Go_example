package main

import (
	"fmt"
	"github.com/wizarddewhite/Go_example/gowiki/wikipage"
)

func main() {
	p1 := &wikipage.Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.Save()
	p2, _ := wikipage.LoadPage("TestPage")
	fmt.Println(string(p2.Body))
}

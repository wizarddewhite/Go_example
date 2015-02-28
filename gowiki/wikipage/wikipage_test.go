package wikipage

import "testing"

func TestLoadPage(t *testing.T) {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.Save()
	_, err := LoadPage("TestPage")

	if err != nil {
		t.Errorf("Failed to load page")
	}
}

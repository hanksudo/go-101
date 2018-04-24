package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
)

const (
	templatePath = "tmpl"
	dataPath     = "data"
)

var templates = template.Must(template.New("main").Funcs(
	template.FuncMap{"unescaped": func(x string) interface{} {
		return template.HTML(validLink.ReplaceAllString(x, "<a href=\"$1\">[$1]</a>"))
	}},
).ParseGlob(filepath.Join(templatePath, "*.html")))
var validPath = regexp.MustCompile(`^/(edit|view|save)/(\w+)$`)
var validLink = regexp.MustCompile("\\[([a-zA-Z0-9]+)\\]")

// Page - basic page data structure
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := filepath.Join(dataPath, p.Title+".txt")
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := filepath.Join(dataPath, title+".txt")
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

func main() {
	// Test save/load wiki page
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page.")}
	// p1.save()
	// p2, _ := loadPage("TestPage")
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8000", nil))
}

package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

// RenderTemplate is function for render template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser:", err)
	}
}

// CreateTemplateCache creates template cache as map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}
	fmt.Println(myCache)

	pages, err := filepath.Glob("./static/templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./static/templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./static/templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}

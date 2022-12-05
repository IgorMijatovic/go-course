package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl");
	err := parsedTemplate.Execute(w, nil);
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error 

	// check to see if we already have template in cache
	_, inMap := tc[t]
	if !inMap {
		log.Println("creating template and adding to cache")
		// need to create template
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have template in the cache
		log.Println("Using cached version")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
			log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse template
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	// add template to cache
	tc[t] = tmpl

	return nil
}
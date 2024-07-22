package render

import (
	"bytes"
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/howters/gopack/pkg/config"
	"github.com/howters/gopack/pkg/models"
)
var app *config.AppConfig

func NewTemplates (a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{

	return td
}

func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if(app.UseCache) {
		tc = app.TemplateCache
	} else {
		tc,_ = CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := tc[html]

	if(!ok) {
		log.Fatal("Could not get template")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	err := t.Execute(buf,td)
	if(err != nil) {
		log.Println(err)
	}
	
	//render the template

	_, err = buf.WriteTo(w)
	if (err != nil ) {
		log.Println(err)
	}

	// parsedTemplate, _ := template.ParseFiles("./templates/" + html, "./templates/base.layout.html")
	// err := parsedTemplate.Execute(w, nil)
	// if err != nil {
	// 	fmt.Println("error : ", err)
	// 	return
	// }
}	

func CreateTemplateCache() (map[string] *template.Template, error) {
	myCache := map[string] *template.Template{}
	//get all of the files named *.page.html from templates/

	pages, err := filepath.Glob("./templates/*.page.html")

	if(err != nil) {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if(err != nil) {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		
		if(err != nil) {
			return myCache, err
		}

		if(len(matches) > 0) {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if(err != nil) {
				return myCache, err
			}
		}

		myCache[name] = ts
	}	
	return myCache, nil
}


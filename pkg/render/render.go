package render

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/aungkokoye/go_web/model"
	"github.com/aungkokoye/go_web/pkg/config"
)

const templatePath = "./template/"

var app *config.AppConfig

func RenderTemplate(w http.ResponseWriter, tmplName string, td *model.TempaleData) {

	var tc map[string]*template.Template
	var err error

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
		if err != nil {
			fmt.Println(err)
		}
	}

	//get requested from cache
	t, ok := tc[tmplName]
	if !ok {
		log.Fatal("Cannot find template in App Cache!")
	}

	//render the template
	err = t.Execute(w, td)
	if err != nil {
		fmt.Println(err)
	}

}

func NewAppConfig(a *config.AppConfig) {
	app = a
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	//get all fields named *.page.tmpl from ./template
	pages, err := filepath.Glob(templatePath + "*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	//loop through all files ending .page.tmpl

	for _, page := range pages {

		// get only file name not the whole path
		name := filepath.Base(page)

		// save template name in Tempale
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		// get all layout files
		layouts, err := filepath.Glob(templatePath + "*.layout.tmpl")

		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			//bind layout file to template
			ts, err = ts.ParseGlob(templatePath + "*.layout.tmpl")

			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, nil

}

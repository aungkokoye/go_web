package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aungkokoye/go_web/pkg/config"
	"github.com/aungkokoye/go_web/pkg/handler"
	"github.com/aungkokoye/go_web/pkg/render"
)

const portNumber = ":8400"

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't create application cache.")
	}

	// store tempalte cache in App config
	app.TemplateCache = tc
	app.UseCache = false

	// set App Config for render module
	render.NewAppConfig(&app)

	// set App Config for handler module
	repo := handler.NewRepo(&app)
	handler.NewHandler(repo)

	http.HandleFunc("/", handler.Repo.Home)
	http.HandleFunc("/about", handler.Repo.About)
	http.HandleFunc("/test", handler.Repo.Home)

	fmt.Printf("Apllication is running on%s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}

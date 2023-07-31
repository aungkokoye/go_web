package handler

import (
	"net/http"

	"github.com/aungkokoye/go_web/model"
	"github.com/aungkokoye/go_web/pkg/config"
	"github.com/aungkokoye/go_web/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (re *Repository) About(w http.ResponseWriter, r *http.Request) {
	rb, err := r.GetBody(r.Body)
	var stringMap = make(map[string]string)
	stringMap["test"] = "hello there!"
	render.RenderTemplate(w, "about.page.tmpl", &model.TempaleData{
		StringMap: stringMap,
	})
}

func (re *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &model.TempaleData{})
}

func (re *Repository) Test(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "test.page.tmpl", &model.TempaleData{})

}

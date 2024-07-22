package handler

import (
	"net/http"

	"github.com/howters/gopack/pkg/config"
	"github.com/howters/gopack/pkg/models"
	"github.com/howters/gopack/pkg/render"
)


var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo (a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr;
	m.App.Session.Put(r.Context(),"remote_ip", remoteIp)


	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello,again"
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}


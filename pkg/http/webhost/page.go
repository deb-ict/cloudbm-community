package webhost

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title       string
	ShowSideBar bool
	Model       interface{}
	Layout      string
	Template    string
	NavItem     string
	NavChild    string
	template    *template.Template
}

func (p *Page) Handle(w http.ResponseWriter, r *http.Request) {
	if len(p.Template) == 0 {
		return
	}

	if p.template == nil {
		if len(p.Layout) > 0 {
			p.template = template.Must(template.ParseFiles(
				"./web/layout/public.html",
				"./web/pages/home.html",
			))
		} else {
			p.template = template.Must(template.ParseFiles(
				"./web/pages/home.html",
			))
		}
	}
	p.template.Execute(w, p)
}

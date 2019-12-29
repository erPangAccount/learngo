package common

import (
	"crawler/frontend/model"
	"html/template"
	"io"
)

type SearchResultTempalteView struct {
	template *template.Template
}

func CreateSearchResultTemplateView(filename string) SearchResultTempalteView {
	return SearchResultTempalteView{
		template: template.Must(template.ParseFiles(filename)),
	}
}

func (s SearchResultTempalteView) Render(w io.Writer, data model.SearchResult) error {
	return s.template.Execute(w, data)
}

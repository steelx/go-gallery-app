package views

import (
	"html/template"
)

// NewView initiates View struct
func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/default.tpl.html",
		"views/layouts/header.tpl.html",
		"views/layouts/footer.tpl.html",
	)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout: layout,
	}
}

// View type
type View struct {
	Template *template.Template
	Layout   string
}

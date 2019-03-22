package views

import (
	"html/template"
)

// NewView initiates View struct
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.tpl.html")

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

// View type
type View struct {
	Template *template.Template
	Layout   string
}

package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	layoutDir   = "views/layouts/"
	templateExt = ".tpl.html"
)

// NewView initiates View struct
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

// View type
type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	err := v.Template.ExecuteTemplate(w, v.Layout, data)
	return err
}

// layoutFiles returns all template files from the destination folder
func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "*" + templateExt)
	if err != nil {
		panic(err)
	}
	return files
}

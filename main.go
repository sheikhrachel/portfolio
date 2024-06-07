package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/sheikhrachel/portfolio/components"
)

const (
	// port is the port number for the server
	port = ":8080"
	// errTemplate is the error message for executing a template
	errTemplate = "error executing template: %v"
	// errServerStart is the error message for starting the server
	errServerStart = "error starting server: %v"

	// dirTemplates is the directory containing templates
	dirTemplates = "templates"
	// dirPartials is the directory containing partial templates
	dirPartials = "templates/partials"
)

// templates is a list of template files
var templates = []string{
	filepath.Join(dirTemplates, "base.html"),
	filepath.Join(dirPartials, "head.html"),
	filepath.Join(dirPartials, "switch.html"),
	filepath.Join(dirPartials, "intro.html"),
	filepath.Join(dirPartials, "background.html"),
	filepath.Join(dirPartials, "education.html"),
	filepath.Join(dirPartials, "experience.html"),
	filepath.Join(dirPartials, "publications.html"),
	filepath.Join(dirPartials, "talks.html"),
	filepath.Join(dirPartials, "footer.html"),
}

type PageData struct {
	IntroData  components.IntroData
	FooterData components.FooterData
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexHandler)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf(errServerStart, err)
	}
}

// indexHandler handles requests for the index page
func indexHandler(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = tmpl.ExecuteTemplate(w, "base.html", PageData{
		IntroData: components.IntroData{
			Name:  "Rachel Sheikh",
			Email: "sheikhrachel97@gmail.com",
		},
		FooterData: components.FooterData{
			CurrentYear: time.Now().Year(),
			SocialLinks: components.GetSocialLinks(),
		},
	}); err != nil {
		log.Printf(errTemplate, err)
	}
}

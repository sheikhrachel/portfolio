package components

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// IntroData represents the data for the intro section of the page
type IntroData struct {
	// Name is the name of the user
	Name string
	// Email is the email address of the user
	Email string
}

// RenderIntro renders the intro section
func RenderIntro(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles(filepath.Join(dirPartials, "intro.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = tmpl.ExecuteTemplate(w, "intro.html", IntroData{
		Name:  "Rachel Sheikh",
		Email: "sheikhrachel97@gmail.com",
	}); err != nil {
		log.Printf(errTemplate, err)
	}
}

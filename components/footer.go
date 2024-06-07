package components

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"
)

// SocialLink represents a social media link
type SocialLink struct {
	// Name is the name of the social media platform
	Name string
	// URL is the URL to the social media profile
	URL string
	// Icon is the path to the icon for the social media platform
	Icon string
}

// FooterData represents the data for the footer section
type FooterData struct {
	// CurrentYear is the current year
	CurrentYear int
	// SocialLinks is a list of social media links
	SocialLinks []SocialLink
}

// RenderFooter renders the footer
func RenderFooter(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles(filepath.Join(dirPartials, "footer.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = tmpl.ExecuteTemplate(w, "footer.html", FooterData{
		CurrentYear: time.Now().Year(),
		SocialLinks: GetSocialLinks(),
	}); err != nil {
		log.Printf(errTemplate, err)
	}
}

// GetSocialLinks returns a list of social media links
func GetSocialLinks() []SocialLink {
	return []SocialLink{
		{"email", "mailto:sheikhrachel97@gmail.com", "../static/img/social/email.svg"},
		{"twitter", "https://twitter.com/rachel_sheikh", "../static/img/social/twitter.svg"},
		{"instagram", "https://instagram.com/rachel.sheikh97", "../static/img/social/instagram.svg"},
		{"github", "https://github.com/sheikhrachel", "../static/img/social/github.svg"},
		{"linkedin", "https://www.linkedin.com/in/rachelsheikh", "../static/img/social/linkedin.svg"},
	}
}

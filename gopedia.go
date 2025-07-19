// Package gopedia provides functions to request the Wikimedia API.
package gopedia

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/text/language"
)

var en, _ = language.English.Base()

// A Project represents a Wikimedia Project such as Wikipedia or Wikimedia Commons.
type Project struct {
	name string
	Lang language.Base // Lang is the language that will be requested from the API.
}

// Name returns the name of the project.
func (project Project) Name() string {
	return project.name
}

// URL returns the base endpoint for the wiki's API.
func (project Project) URL(api string) string {
	url := fmt.Sprintf("https://api.wikimedia.org/%s/v1/%s", api, project.Name())
	if project.Lang.String() != "und" {
		url += "/" + project.Lang.String()
	}
	return url
}

// request requests the endpoint from the base endpoint and marshals a JSON response into v.
func (project Project) request(api string, endpoint string, v any) error {
	url, err := url.JoinPath(project.URL(api), endpoint)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	url = strings.ReplaceAll(url, "%3F", "?")
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		return errors.New(resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &v); err != nil {
		return fmt.Errorf("unmarshal response: %w", err)
	}
	return nil
}

var (
	Wikipedia   = Project{name: "wikipedia", Lang: en}   // Wikipedia - https://www.wikipedia.org/
	Wiktionary  = Project{name: "wiktionary", Lang: en}  // Wiktionary - https://www.wiktionary.org/
	Wikiquote   = Project{name: "wikiquote", Lang: en}   // Wikiquote - https://www.wikiquote.org/
	Wikivoyage  = Project{name: "wikivoyage", Lang: en}  // Wikivoyage - https://www.wikivoyage.org/
	Wikinews    = Project{name: "wikinews", Lang: en}    // Wikinews - https://www.wikinews.org/
	Wikibooks   = Project{name: "wikibooks", Lang: en}   // Wikibooks - https://www.wikibooks.org/
	Wikisource  = Project{name: "wikisource", Lang: en}  // Wikisource - https://wikisource.org/
	Wikiversity = Project{name: "wikiversity", Lang: en} // Wikiversity - https://www.wikiversity.org/
	Commons     = Project{name: "commons"}               // Wikimedia Commons - https://commons.wikimedia.org/
	Wikispecies = Project{name: "wikispecies"}           // Wikispecies - https://species.wikimedia.org/
)

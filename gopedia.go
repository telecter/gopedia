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

// newProject creates a new Project with the provided name and the default language.
func newProject(name string) Project {
	lang, _ := language.English.Base()
	return Project{
		name: name,
		Lang: lang,
	}
}

// A Project represents a Wikimedia Project such as Wikipedia or Wikimedia Commons.
type Project struct {
	name string
	Lang language.Base
}

// Name returns the name of the project.
func (project Project) Name() string {
	return project.name
}

// URL returns the base endpoint for the wiki's API.
func (project Project) URL() string {
	url := "https://api.wikimedia.org/core/v1/" + project.name
	if project != Commons && project != Wikispecies {
		url += "/" + project.Lang.String()
	}
	return url
}

// request requests the endpoint from the base endpoint and marshals a JSON response into v.
func (project Project) request(endpoint string, v any) error {
	url, err := url.JoinPath(project.URL(), endpoint)
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
	Wikipedia   = newProject("wikipedia")
	Wiktionary  = newProject("wiktionary")
	Wikiquote   = newProject("wikiquote")
	Wikivoyage  = newProject("wikivoyage")
	Wikinews    = newProject("wikinews")
	Wikibooks   = newProject("wikibooks")
	Wikisource  = newProject("wikisource")
	Wikiversity = newProject("wikiversity")
	Commons     = newProject("commons")
	Wikispecies = newProject("wikispecies")
)

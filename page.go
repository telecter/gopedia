package gopedia

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// PageMode represents the type of data to return from a page query.
type PageMode int

const (
	PageModeBare   PageMode = iota // Standard query. Includes license and revision metadata.
	PageModeHTML                   // Returns page HTML and standard metadata.
	PageModeSource                 // Returns page source and standard metadata.
)

// A PageLanguage represents an available language translation for a page.
type PageLanguage struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Key   string `json:"key"`
	Title string `json:"title"`
}

// A Page is a Wikipedia article and metadata.
type Page struct {
	Project Project

	ID     int    `json:"id"`    // Page identifier
	Key    string `json:"key"`   // Page title in URL-friendly format
	Title  string `json:"title"` // Page title in reading-friendly format
	Latest struct {
		ID        int       `json:"id"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"latest"` // Information about the latest revision
	/*
		Page content format:

		wikitext (default): MediaWiki-flavored markup, used for most wiki pages

		css: CSS

		javascript: JavaScript

		json: JSON

		text: Plain text
	*/
	ContentModel string `json:"content_model"`
	License      struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	} `json:"license"` // Information about the wiki's license

	HTML   string `json:"html"`   // Latest page content in HTML. This field is only filled when using PageModeHTML.
	Source string `json:"source"` // Latest page content in the format specified by the content_model field. This field is only filled when using PageModeSource.
}

// GetLanguages returns all available languages for the page.
func (page Page) GetLanguages() ([]PageLanguage, error) {
	var data []PageLanguage
	path := fmt.Sprintf("/page/%s/links/language", page.Key)
	if err := page.Project.request(path, &data); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	return data, nil
}

// FetchFiles returns all files on the page.
func (page Page) FetchFiles() ([]File, error) {
	var data map[string][]File
	path := fmt.Sprintf("/page/%s/links/media", page.Key)
	if err := page.Project.request(path, &data); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	files, ok := data["files"]
	if !ok {
		return nil, errors.New("malformed response")
	}
	return files, nil
}

// GetHistory returns the revision history of the page, filtered by the specified filter.
//
// olderThan returns the 20 revisions older than the revision with its ID.
// newerThan returns the 20 revisions newer than the revision with its ID.
//
// They are mutually exclusive.
// If olderThan or newerThan is -1, it is disregarded. Passing -1 for both gives the latest revisions.
func (page Page) GetHistory(olderThan, newerThan int, filter RevisionFilter) ([]PageRevision, error) {
	path, err := url.Parse(fmt.Sprintf("/page/%s/history", page.Key))
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	switch filter {
	case RevisionFilterReverted:
		params.Add("filter", "reverted")
	case RevisionFilterAnonymous:
		params.Add("filter", "anonymous")
	case RevisionFilterBot:
		params.Add("filter", "bot")
	case RevisionFilterMinor:
		params.Add("filter", "minor")
	}
	if olderThan > -1 && newerThan > -1 {
		return nil, fmt.Errorf("olderThan and newerThan cannot both be specified")
	}
	if olderThan > -1 {
		params.Add("older_than", strconv.Itoa(olderThan))
	}
	if newerThan > -1 {
		params.Add("newer_than", strconv.Itoa(newerThan))
	}
	path.RawQuery = params.Encode()

	var data struct {
		Revisions []PageRevision `json:"revisions"`
	}
	if err := page.Project.request(path.String(), &data); err != nil {
		return nil, fmt.Errorf("request: %w", err)
	}
	return data.Revisions, nil
}

// FetchPage requests a page using the specified mode.
func (project Project) FetchPage(title string, mode PageMode) (Page, error) {
	var path string
	switch mode {
	case PageModeBare:
		path = "/page/%s/bare"
	case PageModeHTML:
		path = "/page/%s/with_html"
	case PageModeSource:
		path = "/page/%s"
	}
	path = fmt.Sprintf(path, title)

	var data Page
	if err := project.request(path, &data); err != nil {
		return Page{}, fmt.Errorf("request: %w", err)
	}
	data.Project = project

	return data, nil
}

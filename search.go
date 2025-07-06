package gopedia

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

// A SearchType is a method of searching a wiki.
type SearchType int

const (
	SearchTypeContent SearchType = iota // Searches wiki pages for the given search terms, and returns matching pages.
	SearchTypeTitles                    // Searches wiki page titles, and returns pages with titles that begin with the provided search terms.
)

// A SearchResult represents a wiki page matching the requested search.
type SearchResult struct {
	ID           int    `json:"id"`
	Key          string `json:"key"`
	Title        string `json:"title"`
	Excerpt      string `json:"excerpt"`
	MatchedTitle any    `json:"matched_title"`
	Description  string `json:"description"`
	Thumbnail    struct {
		Mimetype string `json:"mimetype"`
		Size     any    `json:"size"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
		Duration any    `json:"duration"`
		URL      string `json:"url"`
	} `json:"thumbnail"`
}

// Search searches the wiki with the specified query and method, returning the number of pages specified by limit.
//
// If limit is -1, the default value (50) will be used.
func (project Project) Search(query string, by SearchType, limit int) ([]SearchResult, error) {
	var path string
	switch by {
	case SearchTypeContent:
		path = "/search/page"
	case SearchTypeTitles:
		path = "/search/title"
	}
	endpoint, err := url.Parse(path)
	if err != nil {
		panic(err)
	}

	params := url.Values{
		"q": {query},
	}
	if limit > -1 {
		params.Add("limit", strconv.Itoa(limit))
	}
	endpoint.RawQuery = params.Encode()

	var data map[string][]SearchResult

	if err := project.request(endpoint.String(), &data); err != nil {
		return nil, fmt.Errorf("request search results: %w", err)
	}

	results, ok := data["pages"]
	if !ok {
		return nil, errors.New("malformed search results")
	}

	return results, nil
}

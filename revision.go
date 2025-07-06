package gopedia

import (
	"fmt"
	"time"
)

// A RevisionFilter limits what page revisions are shown.
type RevisionFilter int

const (
	RevisionFilterNone      RevisionFilter = iota // No revision filtering
	RevisionFilterReverted                        // Only show revisions that revert an earlier edit
	RevisionFilterAnonymous                       // Only show revisions made by anonymous users
	RevisionFilterBot                             // Only show revisions made by bots
	RevisionFilterMinor                           // Only show revisions marked as minor edits
)

// A PageRevision represents a change to wiki page.
type PageRevision struct {
	ID   int `json:"id"` // Revision identifier
	Page struct {
		ID    int    `json:"id"`    // Page identifier
		Title string `json:"title"` // Page title in reading-friendly format
	} `json:"page"` // Information about the page of the revision. This field is only filled when it is obtained via FetchRevision.
	Size      int       `json:"size"`      // Size of the revision in bytes
	Minor     bool      `json:"minor"`     // Revision is marked as a minor edit.
	Timestamp time.Time `json:"timestamp"` // Time of the edit
	User      struct {
		ID   int     `json:"id"`   // Username, or originating IP address for anonymous users
		Name *string `json:"name"` // User identifier, or nil for anonymous users
	} `json:"user"` // Information about the user who made the edit
	Comment *string `json:"comment"` // Comment or edit summary written by the editor. For revisions without a comment, nil or an empty string is returned.
	Delta   int     `json:"delta"`   // Number of bytes changed, positive or negative, between a revision and the preceding revision (example: -20). If the preceding revision is unavailable, nil is returned.
}

// FetchRevision retrieves the specified page version based on its ID.
func (project Project) FetchRevision(id int) (PageRevision, error) {
	path := fmt.Sprintf("/revision/%d/bare", id)
	var data PageRevision
	if err := project.request(path, &data); err != nil {
		return PageRevision{}, fmt.Errorf("request: %w", err)
	}
	return data, nil
}

package gopedia

import (
	"fmt"
	"time"
)

// A File represents a file uploaded to a wiki.
type File struct {
	Title              string `json:"title"`                // File title
	FileDescriptionURL string `json:"file_description_url"` // URL for the page describing the file, including license information and other metadata
	Latest             struct {
		Timestamp time.Time `json:"timestamp"` // Last modified timestamp
		User      struct {
			ID   int    `json:"id"`   // User identifier
			Name string `json:"name"` // Username
		} `json:"user"` // User who uploaded the file
	} `json:"latest"` // Object containing information about the latest revision to the file
	Preferred FileArtifact `json:"preferred"` // Information about the file's preferred preview format
	Original  FileArtifact `json:"original"`  // Information about the file's original format
	Thumbnail FileArtifact `json:"thumbnail"` // Information about the file's thumbnail format. This field is only filled when obtained via FetchFile.
}

// A FileArtifact represents the properties of a file.
type FileArtifact struct {
	Mediatype string   `json:"mediatype"` // File type, one of: BITMAP, DRAWING, AUDIO, VIDEO, MULTIMEDIA, UNKNOWN, OFFICE, TEXT, EXECUTABLE, ARCHIVE, or 3D.
	Size      *int     `json:"size"`      // File size in bytes or nil if not available
	Width     *int     `json:"width"`     // Maximum recommended image width in pixels or nil if not available
	Height    *int     `json:"height"`    // Maximum recommended image height in pixels or nil if not available
	Duration  *float32 `json:"duration"`  // Length of the video, audio, or multimedia file or nil for other media types
	URL       string   `json:"url"`       // URL to download the file
}

// FetchFile requests a file based on its title.
func (project Project) FetchFile(title string) (File, error) {
	path := fmt.Sprintf("/file/%s", title)
	var data File
	if err := project.request(path, &data); err != nil {
		return File{}, fmt.Errorf("request: %w", err)
	}
	return data, nil
}

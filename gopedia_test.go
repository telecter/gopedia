package gopedia_test

import (
	"gopedia"
	"testing"
	"time"
)

func TestFetchPageBare(t *testing.T) {
	page, err := gopedia.Wikipedia.FetchPage("Earth", gopedia.PageModeBare)
	if err != nil {
		t.Fatalf("failed to fetch page: %s", err)
	}
	if page.Title == "" {
		t.Errorf("page.Title is empty; wanted non-empty")
	}
	t.Logf("page.Title = %#v", page.Title)
}

func TestFetchPageHTML(t *testing.T) {
	page, err := gopedia.Wikipedia.FetchPage("Earth", gopedia.PageModeHTML)
	if err != nil {
		t.Fatalf("failed to fetch page: %s", err)
	}
	if page.HTML == "" {
		t.Errorf("page.HTML is empty; wanted non-empty")
	}
}

func TestFetchPageSource(t *testing.T) {
	page, err := gopedia.Wikipedia.FetchPage("Earth", gopedia.PageModeSource)
	if err != nil {
		t.Fatalf("failed to fetch page: %s", err)
	}
	if page.Source == "" {
		t.Errorf("page.Source is empty; wanted non-empty")
	}
}

func TestGetPageLanguages(t *testing.T) {
	page, err := gopedia.Wikipedia.FetchPage("Earth", gopedia.PageModeBare)
	if err != nil {
		t.Fatalf("failed to fetch page: %s", err)
	}
	langs, err := page.GetLanguages()
	if err != nil {
		t.Fatalf("failed to fetch page languages: %s", err)
	}
	if len(langs) < 1 {
		t.Errorf("page.GetLanguages returned empty slice; wanted non-empty")
	}
	t.Logf("langs[0].Key = %#v", langs[0].Key)
}

func TestGetHistory(t *testing.T) {
	page, err := gopedia.Wikipedia.FetchPage("Earth", gopedia.PageModeBare)
	if err != nil {
		t.Fatalf("failed to fetch page: %s", err)
	}
	revisions, err := page.GetHistory(-1, -1, gopedia.RevisionFilterNone)
	if err != nil {
		t.Fatalf("failed to fetch page history: %s", err)
	}
	if len(revisions) < 1 {
		t.Errorf("page.GetHistory returned empty slice; wanted non-empty")
	}
	t.Logf("revisions[0].Timestamp = %#v", revisions[0].Timestamp.Format(time.UnixDate))
}

func TestFetchFiles(t *testing.T) {
	page, err := gopedia.Wikipedia.FetchPage("Earth", gopedia.PageModeBare)
	if err != nil {
		t.Fatalf("failed to fetch page: %s", err)
	}
	files, err := page.FetchFiles()
	if err != nil {
		t.Fatalf("failed to fetch page files: %s", err)
	}
	if len(files) < 1 {
		t.Errorf("page.FetchFiles returned empty slice; wanted non-empty")
	}
	t.Logf("files[0].Title = %#v", files[0].Title)
}

func TestSearch(t *testing.T) {
	cases := []struct {
		mode gopedia.SearchType
		name string
	}{
		{mode: gopedia.SearchTypeContent, name: "By Content"},
		{mode: gopedia.SearchTypeTitles, name: "By Titles"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			results, err := gopedia.Wikipedia.Search("Hello", c.mode, -1)
			if err != nil {
				t.Fatalf("failed to search: %s", err)
			}
			if len(results) < 1 {
				t.Errorf("Search returned empty slice; wanted non-empty")
			}
			t.Logf("results[0].Title = %#v", results[0].Title)
		})
	}
}

func TestFetchRevision(t *testing.T) {
	revision, err := gopedia.Wikipedia.FetchRevision(1297707661)
	if err != nil {
		t.Fatalf("failed to fetch revision: %s", err)
	}
	if revision.ID == 0 {
		t.Errorf("revision.ID is 0; wanted non-zero")
	}
	t.Logf("revision.ID = %#v", revision.ID)
}

func TestFetchFile(t *testing.T) {
	file, err := gopedia.Commons.FetchFile("File:Dark_chocolate_bar.jpg")
	if err != nil {
		t.Fatalf("failed to fetch file: %s", err)
	}
	if file.Title == "" {
		t.Errorf("file.Title is empty; wanted non-empty")
	}
	t.Logf("file.Title = %#v", file.Title)
}

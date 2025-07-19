# gopedia

A Go wrapper for the Wikimedia API.

## Usage

This package aims to (hopefully) provide nearly all Wikimedia API endpoints. This is the basic format:

```
gopedia.<Project Name>.<Method>()
```

### Search

You can search for pages based on content or titles:

```go
results, err := gopedia.Wikipedia.Search("chocolate", gopedia.SearchTypeTitles, 20)
```

This searches for articles with "chocolate" in their titles, returning a max of 20.  
**Search Type:**  
`SearchTypeTitles` - search by article title  
`SearchTypeContent` - search by article content

### Pages

To get a page, use the `FetchPage` method:

```go
page, err := gopedia.Wikipedia.FetchPage("Gopher", gopedia.PageModeBare)
```

**Page Modes:**  
`PageModeBare` - only return page metadata  
`PageModeHTML` - return metadata + HTML content  
`PageModeSource` - return metadata + source content (usually wikitext)

### Page History

You can get page revisions with the `GetHistory` method of a page:

```go
revisions, err := page.GetHistory(-1, -1, gopedia.RevisionFilterNone)
```

-1 and -1 are ways to sort what history you get.
They are the `olderThan` and `newerThan` values, meaning the 20 revisions older or newer than the one with the specified ID will be shown.

`olderThen` and `newerThen` are mutually exclusive and setting them to -1 disables them.  
 If both are set to -1, the latest 20 revisions are shown.

**Revision Filters:**  
`RevisionFilterNone` - no revision filtering  
`RevisionFilterReverted` - Only show revisions that revert an earlier edit  
`RevisionFilterAnonymous` - Only show revisions made by anonymous users  
`RevisionFilterBot` - Only show revisions made by bots  
`RevisionFilterMinor` - Only show revisions marked as minor edits

### Files

The `FetchFiles` method of a page returns all files on that page.
To fetch an individual file by name, use the `FetchFile` method of a project:

```go
file, err := gopedia.Commons.FetchFile("File:Dark_chocolate_bar.jpg")
```

### Revision

To fetch a revision by ID, use the `FetchRevision` method of a project:

```go
revision, err := gopedia.Wikipedia.FetchRevision(1297707661)
```

### Misc.

To fetch all available languages for a page and the translated page titles, use the `GetLanguages` method of a page.

## Credit

For go doc comments, I used the descriptions of the endpoints and parameters from the official Wikimedia API documentation to ensure this package is as clearly documented as possible.

The API documentation can be found here: https://api.wikimedia.org/wiki/

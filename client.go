package shiori

import (
	"fmt"
	"net/http"
	"time"
)

// Client provides a request client to execute
// REST API requests with.
type Client struct {
	token    string
	endpoint string

	client *http.Client
}

// NewClient returns a new instance of client with
// the passed authorization token which connects
// to the passed endpoint URL of the siori gateway
// instance.
func NewClient(token string, endpoint string) (c *Client) {
	c = &Client{
		token:    token,
		endpoint: endpoint,
	}

	c.client = &http.Client{
		Timeout: 60 * time.Second,
	}

	return
}

// Links returns a list of links limited by the
// passed limit and offset by the passed offset.
func (c *Client) Links(limit, offset int) (links []*Link, err error) {
	links = make([]*Link, 0)
	path := fmt.Sprintf("links?limt=%d&offset=%d", limit, offset)
	err = c.get(path, &links)
	return
}

// Link returns the info of the pased links ID.
func (c *Client) Link(id int) (link *Link, err error) {
	link = new(Link)
	path := fmt.Sprintf("links/%d", id)
	err = c.get(path, link)
	return
}

// Search executes a search request with the passed
// query string and returns the search result. The
// hits are capped by the passed limit and offset by
// the passed offset.
func (c *Client) Search(query string, limit, offset int) (sr *SearchResult, err error) {
	sr = new(SearchResult)
	path := fmt.Sprintf("links/search?query=%s&limt=%d&offset=%d", query, limit, offset)
	err = c.get(path, sr)
	return
}

// CreateLink creates a new link object and returns it.
func (c *Client) CreateLink(link *Link) (res *Link, err error) {
	res = new(Link)
	err = c.post("links", link, res)
	return
}

// UpdateLink replaces the link object (by ID) passed.
//
// `nil` or empty fields are interpreted as to be reset.
func (c *Client) UpdateLink(link *Link) (res *Link, err error) {
	res = new(Link)
	path := fmt.Sprintf("links/%d", link.Id)
	err = c.put(path, link, res)
	return
}

// RemoveLink removes a link by passed ID of the link.
func (c *Client) RemoveLink(id int) (err error) {
	path := fmt.Sprintf("links/%d", id)
	err = c.delete(path, nil)
	return
}

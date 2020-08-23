package shiori

// Link request and response model.
type Link struct {
	Description string   `json:"description"`
	Id          int      `json:"id,omitempty"`
	IdStr       string   `json:"id_str,omitempty"`
	Tags        []string `json:"tags"`
	Url         string   `json:"url"`
}

// SearchResult response model.
type SearchResult struct {
	Hits             []*Link `json:"hits"`
	ProcessingTimeMs int     `json:"processing_time_ms"`
	Query            string  `json:"query"`
}

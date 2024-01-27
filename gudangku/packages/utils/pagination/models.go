package pagination

type (
	PaginationResponse struct {
		FirstPageURL string `json:"first_page_url"`
		From         int    `json:"from"`
		LastPage     int    `json:"last_page"`
		LastPageURL  string `json:"last_page_url"`
		Links        []Link `json:"links"`
		NextPageURL  string `json:"next_page_url"`
		Path         string `json:"path"`
		PrevPageURL  string `json:"prev_page_url"`
		To           int    `json:"to"`
	}
	Link struct {
		URL    string `json:"url"`
		Label  string `json:"label"`
		Active bool   `json:"active"`
	}
)

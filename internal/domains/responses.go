package domains

type Reponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
	Data   any    `json:"data,omitempty"`
}

type QuoteResponse struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

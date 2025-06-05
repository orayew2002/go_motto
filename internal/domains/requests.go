package domains

type CreateQuoteRequest struct {
	Author string `json:"author" required:"true"`
	Quote  string `json:"quote" required:"true"`
}

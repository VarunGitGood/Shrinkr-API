package types

// Redis
type LinkDTO struct {
	ShortURL    string `json:"shortURL"`
	LongURL     string `json:"longURL"`
	Description string `json:"description"`
	Expiration int    `json:"expiration"`
}

func (l *LinkDTO) Validate() error {
	if l.ShortURL == "" || l.LongURL == "" || l.Description == "" {
		return &ValidationError{Message: "Missing required fields"}
	}
	return nil
}
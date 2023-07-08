package database

// MongoDB
type User struct {
	Username string `bson:"username"`
	Joined   string `bson:"joined"`
}

type LinkInfo struct {
	Key         string `bson:"key"`
	Clicks      int    `bson:"clicks,omitempty"`
	Passcode    string `bson:"passcode,omitempty"`
	LongURL     string `bson:"longURL"`
	Description string `bson:"description"`
	Created     string `bson:"created"`
	CreatedBy   string `bson:"createdBy"`
}

// Redis
type LinkDTO struct {
	ShortURL    string `json:"shortURL"`
	LongURL     string `json:"longURL"`
	Description string `json:"description"`
}

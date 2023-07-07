package database

type User struct {
	Username string `bson:"username"`
	Joined   string `bson:"joined"`
}

type LinkInfo struct {
	Key  string `bson:"key"`
	Clicks int `bson:"clicks"`
	Created string `bson:"created"`
}

type Link struct {
	ShortURL    string `json:"shortURL"`
	LongURL     string `json:"longURL"`
	Description string `json:"description"`
}

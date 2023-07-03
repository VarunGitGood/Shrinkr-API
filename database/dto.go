package database

type User struct {
	Username string
	Password string
}

type Link struct {
	ShortURL    string
	LongURL     string
	Description string
}

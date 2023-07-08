package types

type User struct {
	Username string `bson:"username"`
	Joined   string `bson:"joined"`
}

type LinkInfo struct {
	Key         string `bson:"key"`
	Clicks      int    `bson:"clicks,omitempty"`
	Passcode    string `bson:"passcode,omitempty"`
	Description string `bson:"description,omitempty"`
	LongURL     string `bson:"longURL"`
	Created     string `bson:"created"`
	CreatedBy   string `bson:"createdBy"`
}
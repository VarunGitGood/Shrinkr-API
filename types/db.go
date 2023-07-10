package types

type User struct {
	Username string `bson:"username"`
	Joined   string `bson:"joined"`
	Links    int    `bson:"links"`
}

type LinkInfo struct {
	Key         string `bson:"key"`
	Clicks      int    `bson:"clicks,omitempty"`
	Passcode    string `bson:"passcode,omitempty"`
	Expiration  int    `bson:"expiration,omitempty"`
	Description string `bson:"description"`
	LongURL     string `bson:"longURL"`
	Created     string `bson:"created"`
	CreatedBy   string `bson:"createdBy"`
}

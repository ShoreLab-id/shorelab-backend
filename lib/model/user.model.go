package model

type User struct {
	ID       string `bson:"_id,omitempty"`
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password []byte `bson:"password"`
}

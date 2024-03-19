package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID  primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Title   string             `json:"title,omitempty" bson:"title,omitempty"`
	Content string             `json:"content,omitempty" bson:"content,omitempty"`
}

type Comment struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	PostID  primitive.ObjectID `json:"post_id,omitempty" bson:"post_id,omitempty"`
	Content string             `json:"content,omitempty" bson:"content,omitempty"`
}

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

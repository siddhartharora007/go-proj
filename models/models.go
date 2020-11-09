package models

import "go.mongodb.org/mongo-driver/bson/primitive"



type Article struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Id   string             `json:"id,omitempty" bson:"id,omitempty"`
	Title  string             `json:"Title" bson:"Title,omitempty"`
	SubTitle  string             `json:"STitle" bson:"STitle,omitempty"`
	Content  string             `json:"Content" bson:"Content,omitempty"`

}
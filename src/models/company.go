package models

import "github.com/beevik/guid"

type Company struct {
	Id *guid.Guid `json:"id,omitempty" bson:"id,omitempty"`
	Name string `json:"name" bson:"name"`
}

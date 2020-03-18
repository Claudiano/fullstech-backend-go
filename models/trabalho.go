package models

type Trabalho struct {
	Id    string  `json:"_id,omitempty" bson:"_id,omitempty"`
	Nome  string  `json:"nome,omitempty" bson:"nome,omitempty"`
	Valor float32 `json:"valor,omitempty" bson:"valor,omitempty"`
}

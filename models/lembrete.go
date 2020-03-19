package models

type Lembrete struct {
	Id       string `json:"_id,omitempty" bson:"_id,omitempty"`
	Conteudo string `json:"conteudo,omitempty" bson:"conteudo,omitempty"`
}

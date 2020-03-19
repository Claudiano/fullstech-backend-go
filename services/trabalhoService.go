package services

import (
	"fullstech-backend-go/models"
	"fullstech-backend-go/repositories"
)

type TrabalhoService struct{}

var repository repositories.TrabalhoRepository

func (TrabalhoService) GetAllTrabalhos() []models.Trabalho {
	return repository.FindAllTrabalhos()
}

func (TrabalhoService) SaveTrabalho(trabalho models.Trabalho) interface{} {
	return repository.Save(trabalho)
}

func (TrabalhoService) DeleteTrabalho(id string) interface{} {
	return repository.Delete(id)
}

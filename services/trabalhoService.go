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

func (TrabalhoService) SaveTrabalho(trabalho models.Trabalho) string {
	return repository.Save(trabalho)
}

func (TrabalhoService) DeleteTrabalho(id int64) string {
	return repository.Delete(id)
}

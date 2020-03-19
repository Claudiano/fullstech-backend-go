package services

import (
	"fullstech-backend-go/models"
	"fullstech-backend-go/repositories"
)

type LembreteService struct{}

var repositoryLembrete repositories.LembreteRepository

func (LembreteService) GetAllLembretes() []models.Lembrete {
	return repositoryLembrete.FindAllLembretes()
}

func (LembreteService) SaveLembrete(lembrete models.Lembrete) interface{} {
	return repositoryLembrete.Save(lembrete)
}

func (LembreteService) DeleteLembrete(id string) interface{} {
	return repositoryLembrete.Delete(id)
}

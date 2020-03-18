package services

import "fullstech-backend-go/repositories"

type LembreteService struct{}

var repositoryLembrete repositories.LembreteRepository

func (LembreteService) GetAllLembretes() string {
	return repositoryLembrete.FindAllLembretes()
}

func (LembreteService) SaveTrabalho() string {
	return repositoryLembrete.Save()
}

func (LembreteService) DeleteTrabalho(id int64) string {
	return repositoryLembrete.Delete(id)
}

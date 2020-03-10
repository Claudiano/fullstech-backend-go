package services

import "fullstech-backend-go/repositories"

type TrabalhoService struct{}

var repository repositories.TrabalhoRepository

func (TrabalhoService) GetAllTrabalhos() string {
	return repository.FindAllTrabalhos()
}

func (TrabalhoService) SaveTrabalho() string {
	return repository.Update()
}

func (TrabalhoService) DeleteTrabalho(id int64) string {
	return repository.Delete(id)
}

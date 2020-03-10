package repositories

import "strconv"

type TrabalhoRepository struct{}

func (TrabalhoRepository) FindAllTrabalhos() string {
	return "retorna todos os trabalhos"
}

func (TrabalhoRepository) Save() string {
	return "retorna todos os trabalhos"
}

func (TrabalhoRepository) Delete(id int64) string {
	return "exclui um trabalho, com  id = " + strconv.FormatInt(id, 10)
}

package repositories

import "strconv"

type LembreteRepository struct{}

func (LembreteRepository) FindAllLembretes() string {
	return "retorna todos os lembretes"
}

func (LembreteRepository) Save() string {
	return "Salvar um lembrete"
}

func (LembreteRepository) Delete(id int64) string {
	return "exclui um lembrete, com  id = " + strconv.FormatInt(id, 10)
}

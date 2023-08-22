package character

import "github.com/Seiya-Tagami/favorite-character-api/domain/entity"

type Repository interface {
	Insert(character *entity.Character) error
	SelectALL(characters *[]entity.Character) error
	SelectById(character *entity.Character, id int) error
	// UpdateById(character *entity.Character) (*entity.Character, error)
	DeleteById(id int) error
}
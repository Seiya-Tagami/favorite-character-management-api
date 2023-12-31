package character

import (
	"fmt"

	"github.com/Seiya-Tagami/favorite-character-management-api/domain/entity"
	"github.com/Seiya-Tagami/favorite-character-management-api/domain/repository/character"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) character.Repository {
	return &Repository{db}
}

func (r *Repository) Insert(character *entity.Character) error {
	fmt.Println("hi")
	if err := r.db.Create(character).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) SelectALL(characters *[]entity.Character) error {
	if err := r.db.Find(characters).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) SelectById(character *entity.Character, id int) error {
	if err := r.db.First(character, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateById(character *entity.Character, id int) error {
	result := r.db.Model(character).Where("id = ?", id).Updates(character)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (r *Repository) DeleteById(id int) error {
	result := r.db.Where("id= ? ", id).Delete(&entity.Character{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

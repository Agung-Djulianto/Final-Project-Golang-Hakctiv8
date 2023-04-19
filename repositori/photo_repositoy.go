package repositori

import (
	"Project-Akhir/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{
		db: db,
	}
}

func (pr *PhotoRepository) Save(photo models.Photo) (models.Photo, error) {
	tx := pr.db.Create(&photo)

	return photo, tx.Error
}

func (pr *PhotoRepository) Get() ([]models.Photo, error) {
	photo := make([]models.Photo, 0)

	tx := pr.db.Find(&photo)
	return photo, tx.Error
}

func (pr *PhotoRepository) GetById(id string) (models.Photo, error) {
	photo := models.Photo{}

	tx := pr.db.First(&photo, "id = ?", id)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return models.Photo{}, models.ErrorNotFound
	}
	return photo, tx.Error
}

func (pr *PhotoRepository) Update(update models.Photo, id string) (models.Photo, error) {
	tx := pr.db.Clauses(clause.Returning{
		Columns: []clause.Column{
			{Name: "id"},
			{Name: "user_id"},
			{Name: "created_at"},
			{Name: "updated_at"},
		},
	}).Where("id = ?", id).Updates(&update)

	return update, tx.Error
}

func (pr *PhotoRepository) Delete(id string) error {
	photo := models.Photo{
		ID: id,
	}

	tx := pr.db.Select("Comments").Delete(&photo)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

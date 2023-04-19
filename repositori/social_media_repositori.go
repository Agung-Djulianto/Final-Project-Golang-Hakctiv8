package repositori

import (
	"Project-Akhir/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ISocialMediaRepository interface {
	Get() ([]models.SocialMedia, error)
	GetOne(id string) (models.SocialMedia, error)
	Save(socialMedia models.SocialMedia) (models.SocialMedia, error)
	Update(updateSocialMedia models.SocialMedia, id string) (models.SocialMedia, error)
	Delete(id string) error
}

type SocialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{
		db: db,
	}
}

func (smr *SocialMediaRepository) Save(socialMedia models.SocialMedia) (models.SocialMedia, error) {
	tx := smr.db.Create(&socialMedia)

	return socialMedia, tx.Error
}

func (smr *SocialMediaRepository) Get() ([]models.SocialMedia, error) {
	sosmed := make([]models.SocialMedia, 0)

	tx := smr.db.Find(&sosmed)
	return sosmed, tx.Error
}

func (smr *SocialMediaRepository) GetById(id string) (models.SocialMedia, error) {
	sosmed := models.SocialMedia{}

	tx := smr.db.First(&sosmed, "id = ?", id)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return models.SocialMedia{}, models.ErrorNotFound
	}
	return sosmed, tx.Error
}

func (smr *SocialMediaRepository) Update(UpdateSosmed models.SocialMedia, id string) (models.SocialMedia, error) {
	tx := smr.db.Clauses(clause.Returning{
		Columns: []clause.Column{
			{Name: "id"},
			{Name: "user_id"},
			{Name: "created_at"},
			{Name: "updated_at"},
		},
	}).Where("id = ?", id).Updates(&UpdateSosmed)
	return UpdateSosmed, tx.Error
}

func (smr *SocialMediaRepository) Delete(id string) error {
	sosmed := models.SocialMedia{}

	tx := smr.db.Delete(&sosmed, "id = ?", id)

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

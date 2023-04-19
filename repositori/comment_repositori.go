package repositori

import (
	"Project-Akhir/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (cr *CommentRepository) Save(comment models.Comment) (models.Comment, error) {
	tx := cr.db.Save(&comment)
	return comment, tx.Error
}

func (cr *CommentRepository) GetAll() ([]models.Comment, error) {
	comment := make([]models.Comment, 0)

	tx := cr.db.Find(&comment)
	return comment, tx.Error
}

func (cr *CommentRepository) GetById(id string) (models.Comment, error) {
	commnet := models.Comment{}

	tx := cr.db.First(&commnet, "id = ?", id)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return models.Comment{}, models.ErrorNotFound
	}
	return commnet, tx.Error
}

func (cr *CommentRepository) Update(updateComment models.Comment, id string) (models.Comment, error) {
	tx := cr.db.
		Clauses(clause.Returning{
			Columns: []clause.Column{
				{Name: "id"},
				{Name: "user_id"},
				{Name: "photo_id"},
				{Name: "created_at"},
				{Name: "updated_at"},
			},
		},
		).
		Where("id = ?", id).
		Updates(&updateComment)
	return updateComment, tx.Error
}

func (cr *CommentRepository) Delete(id string) error {
	comment := models.Comment{}

	tx := cr.db.Delete(&comment, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

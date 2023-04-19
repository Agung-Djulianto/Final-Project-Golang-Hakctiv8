package repositori

import (
	"Project-Akhir/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Save(newUser models.User) (models.User, error)
	GetByUsername(username string) (models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUSerRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}

}
func (ur *UserRepository) Save(newUser models.User) (models.User, error) {
	tx := ur.db.Create(&newUser)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return newUser, nil
}

func (ur *UserRepository) GetByUsername(username string) (models.User, error) {
	user := models.User{}
	tx := ur.db.First(&user, "user_name = ?", username)

	return user, tx.Error
}

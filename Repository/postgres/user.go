package postgres

import (
	"errors"
	"time"

	model "github.com/MelvinKim/Design-Reddit-API/Model"
	"github.com/MelvinKim/Design-Reddit-API/repository"
	"github.com/jinzhu/gorm"
)

type DBUserRepository struct {
	DB *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) repository.UserRepository {
	return &DBUserRepository{db}
}

func (r *DBUserRepository) Create(user *model.User) (*model.User, error) {
	err := r.DB.Debug().Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *DBUserRepository) Get(id int) (*model.User, error) {
	var user model.User
	err := r.DB.Debug().First(&user, id).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (r *DBUserRepository) Update(id int, user *model.User) (*model.User, error) {

	db := r.DB.Debug().Model(model.User{}).Where("id = ?", id).Take(model.User{}).UpdateColumns(
		map[string]interface{}{
			"password":   user.Password,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return nil, db.Error
	}

	err := db.Debug().Model(model.User{}).Where("id = ?", id).Take(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *DBUserRepository) FindAll() ([]*model.User, error) {
	users := []*model.User{}
	err := r.DB.Debug().Model(model.User{}).Limit(100).Find(users).Error
	if err != nil {
		return nil, err
	}
	return users, err
}

func (r *DBUserRepository) Delete(id int) (int64, error) {
	db := r.DB.Debug().Model(model.User{}).Where("id = ?", id).Take(&model.User{}).Delete(&model.User{})
	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}

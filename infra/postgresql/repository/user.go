package repository

import (
	"base-gin-golang/domain/entity"
	"base-gin-golang/domain/repository"
	"base-gin-golang/infra/postgresql"
	"base-gin-golang/infra/postgresql/model"
	dataPkg "base-gin-golang/pkg/data"
	"context"

	"github.com/google/wire"
)

type userRepository struct {
	db          *postgresql.Database
	dataService dataPkg.Service
}

var UserProviderSet = wire.NewSet(NewUserRepository)

func NewUserRepository(db *postgresql.Database, dataService dataPkg.Service) repository.UserRepository {
	return &userRepository{
		db:          db,
		dataService: dataService,
	}
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := &model.User{}
	err := r.db.WithContext(ctx).Model(&model.User{}).
		Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	result := &entity.User{}
	err = r.dataService.Copy(result, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*entity.User, error) {
	user := &model.User{}
	err := r.db.WithContext(ctx).Model(&model.User{}).
		Where("id = ?", id).
		Error
	if err != nil {
		return nil, err
	}
	result := &entity.User{}
	err = r.dataService.Copy(result, user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

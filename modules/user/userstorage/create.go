package userstorage

import (
	"context"
	"go_restaurant/common"
	"go_restaurant/modules/user/usermodel"
)

func (s *sqlStorage) CreateUser(ctx context.Context, createUserData *usermodel.UserCreate) error {
	db := s.db.Begin()

	if err := db.Table(usermodel.UserCreate{}.TableName()).Create(&createUserData).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}

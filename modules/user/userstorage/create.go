package userstorage

import (
	"context"
	"fmt"
	"go_restaurant/common"
	"go_restaurant/modules/user/usermodel"
	"log"
)

func (s *sqlStorage) CreateUser(ctx context.Context, createUserData *usermodel.UserCreate) error {
	db := s.db.Begin()

	log.Println("create user data", createUserData)

	if err := db.Table(usermodel.UserCreate{}.TableName()).Create(&createUserData).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	fmt.Println("after create user data", createUserData)

	return nil
}

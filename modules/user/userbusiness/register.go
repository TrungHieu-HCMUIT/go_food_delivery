package userbusiness

import (
	"context"
	"go_restaurant/common"
	"go_restaurant/modules/user/usermodel"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, relations ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, createUserData *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	storage RegisterStorage
	hasher  Hasher
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		storage: registerStorage,
		hasher:  hasher,
	}
}

func (business *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, err := business.storage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return common.ErrEntityExisted(usermodel.EntityName, err)
	}

	salt := common.GenSalt(50)

	data.Password = business.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	data.Status = 1

	if err := business.storage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}

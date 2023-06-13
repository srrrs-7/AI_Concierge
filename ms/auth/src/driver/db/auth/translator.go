package auth

import (
	"auth/driver/db/model"
	"auth/pkg/entity"
)

func ConvertAuthEntity(authModel *model.Auth) (*entity.Auth, error) {
	return &entity.Auth{
		UserID:   authModel.UserID,
		Password: authModel.Password,
		Email:    authModel.Email,
		Scope:    authModel.Scope,
	}, nil
}

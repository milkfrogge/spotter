package storage

import (
	"SpotterBackend/src/internal/domain/user/model"
	"context"
)

type Storage interface {
	CreateByEmail(ctx context.Context, user model.CreateByEmailDTO) (int64, error)
	CreateByPhoneNumber(ctx context.Context, user model.CreateByPhoneNumberDTO) (int64, error)
	FindOne(ctx context.Context, id int) (model.User, error)
	Update(ctx context.Context, user model.UpdateUserDTO)
	Delete(ctx context.Context, id int) error
}

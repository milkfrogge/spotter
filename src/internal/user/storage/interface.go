package storage

import (
	"SpotterBackend/src/internal/user/model"
	"context"
)

type Storage interface {
	CreateByEmail(ctx context.Context, user model.CreateByEmailDTO) (int64, error)
	CreateByPhoneNumber(ctx context.Context, user model.CreateByPhoneNumberDTO)
	FindOne(ctx context.Context, id int)
	Update(ctx context.Context, user model.User)
	Delete(ctx context.Context, id int)
}

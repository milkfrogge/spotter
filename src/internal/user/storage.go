package user

import "context"

type Storage interface {
	CreateByEmail(ctx context.Context, user CreateByEmailDTO)
	CreateByPhoneNumber(ctx context.Context, user CreateByPhoneNumberDTO)
	FindOne(ctx context.Context, id int)
	Update(ctx context.Context, user User)
	Delete(ctx context.Context, id int)
}

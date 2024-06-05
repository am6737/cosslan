package user

import "context"

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	Get(ctx context.Context, id string) (*User, error)
	Delete(ctx context.Context, id string) error
	Find(ctx context.Context, filter *UserFilter) ([]*User, error)
	Update(ctx context.Context, user *User) error
}

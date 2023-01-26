package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type authentificator struct {
	*redis.Client
}

type Authentificator interface {
	CheckUserByID(context.Context, int) bool
	UserValidate(context.Context, int) error
}

func NewAuthentificator(r *redis.Client) *authentificator {
	return &authentificator{r}
}

type Store struct {
	UserID string
	Valid  bool
}

func (a *authentificator) UserValidate(ctx context.Context, id int) error {
	u := Store{
		UserID: fmt.Sprint(id),
		Valid:  true,
	}
	return a.Set(ctx, u.UserID, u.Valid, time.Hour*8).Err()
}

func (a *authentificator) CheckUserByID(ctx context.Context, id int) bool {
	u := Store{
		UserID: fmt.Sprint(id),
		Valid:  true,
	}
	res := a.Get(ctx, u.UserID)
	if err := res.Err(); err != nil {
		return false
	}

	val, _ := res.Bool()
	return val
}

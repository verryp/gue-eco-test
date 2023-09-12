package repository

import (
	"context"

	"github.com/verryp/gue-eco-test/internal/auth/model"
)

type activityLogRepo struct {
	*Option
}

func NewActivityLogRepo(opt *Option) ActivityLog {
	return &activityLogRepo{opt}
}

func (repo *activityLogRepo) Create(ctx context.Context, log *model.ActivityLog) error {
	return repo.DB.Insert(log)
}

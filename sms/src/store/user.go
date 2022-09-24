package store

import (
	"context"

	"sms/src/models"

	"gorm.io/gorm"
)

type UserStore struct {
	*gorm.DB
}

func NewUseStore(db *gorm.DB) *UserStore {
	return &UserStore{db}
}

func (s *UserStore) Save(ctx context.Context, r *models.User) error {
	return s.DB.WithContext(ctx).Save(r).Error
}

func (s *UserStore) GetByReqId(ctx context.Context, reqId string) (*models.User, error) {
	var ans models.User
	err := s.WithContext(ctx).Where("req_id=?", reqId).First(&ans).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, nil
	}

	return &ans, err
}

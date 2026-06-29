package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
)

func (b *Buyer) GetProfile(ctx context.Context, buyerID uuid.UUID) (*domain.User, error) {
	user, err := b.UserRepo.FindByID(ctx, buyerID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}
	user.Password = ""
	return user, nil
}

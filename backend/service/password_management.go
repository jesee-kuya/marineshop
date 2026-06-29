package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jesee-kuya/marineshop/domain"
	"github.com/jesee-kuya/marineshop/utils"
)

func (s *Auth) ChangePassword(ctx context.Context, userID uuid.UUID, req *domain.ChangePasswordRequest) error {
	user, err := s.UserRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}
	if user == nil {
		return domain.ErrUserNotFound
	}

	if !utils.CheckPassword(req.OldPassword, user.Password) {
		return domain.ErrInvalidCredentials
	}

	hashed, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	return s.UserRepo.UpdatePassword(ctx, userID, hashed)
}

func (s *Auth) ResetPassword(ctx context.Context, req *domain.ResetPasswordRequest) error {
	user, err := s.UserRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if user == nil {
		return domain.ErrInvalidCredentials
	}

	if !utils.CheckPassword(req.CurrentPassword, user.Password) {
		return domain.ErrInvalidCredentials
	}

	hashed, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return err
	}

	return s.UserRepo.UpdatePassword(ctx, user.ID, hashed)
}

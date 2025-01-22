package biz

import (
	"finance-tracker/internal/mods/auth/dal"
	"finance-tracker/internal/mods/auth/schema"
	"finance-tracker/pkg/config"
	"finance-tracker/pkg/crypto/hash"
	"finance-tracker/pkg/errors"
	"finance-tracker/pkg/jwt"
)

type Login struct {
	UserDAL *dal.User
	UserBIZ User
}

func (a *Login) Login(formItem *schema.LoginForm) (*schema.LoginToken, error) {
	user, err := a.UserDAL.GetByUsername(formItem.Trim().Username)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, errors.BadRequest(config.ErrInvalidUsernameOrPassword, "%s", "Invalid credentials")
	} else if user.Status != schema.UserStatusActivated {
		return nil, errors.BadRequest("", "User status is not activated, please contact the administrator")
	}

	// check password
	if err := hash.CompareHashAndPassword(user.Password, formItem.Password); err != nil {
		return nil, errors.BadRequest(config.ErrInvalidUsernameOrPassword, "%s", "2 Invalid credentials")
	}

	// Load configuration
	cfg := config.LoadConfigs()

	// Generate JWT token
	jwtManager := jwt.NewJWTManager(cfg.JWTConfig.JWTSecretKey, cfg.JWTConfig.TokenExpiry)
	token, err := jwtManager.GenerateToken(uint(user.ID), user.Email)
	if err != nil {
		return nil, errors.InternalServerError("", "Failed to generate token")
	}

	loginToken := &schema.LoginToken{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresAt:   int64(cfg.JWTConfig.TokenExpiry),
	}

	return loginToken, nil
}

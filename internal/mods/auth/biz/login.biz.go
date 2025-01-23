package biz

import (
	"finance-tracker/internal/mods/auth/dal"
	"finance-tracker/internal/mods/auth/schema"
	"finance-tracker/pkg/config"
	"finance-tracker/pkg/crypto/hash"
	"finance-tracker/pkg/errors"
	"finance-tracker/pkg/jwt"
	rds "finance-tracker/pkg/redis"
	"time"

	gjwt "github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type Login struct {
	UserDAL     *dal.User
	UserBIZ     User
	RedisClient *redis.Client // Tambahkan Redis client
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

func (a *Login) Logout(tokenString string) error {
	cfg := config.LoadConfigs()

	// Parse the JWT token to extract the expiration time
	token, err := gjwt.Parse(tokenString, func(token *gjwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*gjwt.SigningMethodHMAC); !ok {
			return nil, errors.BadRequest("", "invalid signing method")
		}
		return []byte(cfg.JWTConfig.JWTSecretKey), nil
	})
	if err != nil || !token.Valid {
		return errors.BadRequest("", "invalid token")
	}

	claims, ok := token.Claims.(gjwt.MapClaims)
	if !ok || !token.Valid {
		return errors.BadRequest("", "invalid token claims")
	}

	expiration, ok := claims["exp"].(float64)
	if !ok {
		return errors.BadRequest("", "missing exp field in token")
	}

	// Calculate the remaining expiration time
	expirationTime := time.Until(time.Unix(int64(expiration), 0))

	return rds.BlacklistToken(a.RedisClient, tokenString, expirationTime)
}

func (a *Login) UpdatePassword(id int, updateItem *schema.UpdateLoginPassword) error {
	user, err := a.UserDAL.Get(id)

	if err != nil {
		return err
	} else if user == nil {
		return errors.NotFound("", "User not found")
	}

	// check old password
	if err := hash.CompareHashAndPassword(user.Password, updateItem.OldPassword); err != nil {
		return errors.BadRequest("", "Incorrect old password")
	}

	// update password
	newPassword, err := hash.GeneratePassword(updateItem.NewPassword)
	if err != nil {
		return err
	}

	return a.UserDAL.UpdatePasswordByID(user.ID, newPassword)
}

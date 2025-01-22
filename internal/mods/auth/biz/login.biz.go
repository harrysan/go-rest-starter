package biz

import (
	"finance-tracker/internal/mods/auth/dal"
	"finance-tracker/internal/mods/auth/schema"
	"finance-tracker/pkg/config"
	"finance-tracker/pkg/crypto/hash"
	"finance-tracker/pkg/errors"
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
		return nil, errors.BadRequest(config.ErrInvalidUsernameOrPassword, "%s", "Incorrect username "+formItem.Username+" or password")
	} else if user.Status != schema.UserStatusActivated {
		return nil, errors.BadRequest("", "User status is not activated, please contact the administrator")
	}

	// check password
	if err := hash.CompareHashAndPassword(user.Password, formItem.Password); err != nil {
		// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formItem.Password)); err != nil {
		// fmt.Println(bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formItem.Password)))
		return nil, errors.BadRequest(config.ErrInvalidUsernameOrPassword, "%s", "2 Incorrect username or password, Password "+user.Password+" Form password "+formItem.Password)
	}

	token := &schema.LoginToken{
		AccessToken: "mock-secret-token",
		TokenType:   "Bearer",
		ExpiresAt:   1000,
	}

	return token, nil
}

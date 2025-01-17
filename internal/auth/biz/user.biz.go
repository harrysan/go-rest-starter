package biz

import (
	"finance-tracker/internal/auth/dal"
	"finance-tracker/internal/auth/schema"
	"finance-tracker/pkg/config"
	"finance-tracker/pkg/errors"
	"time"
)

type User struct {
	UserDAL *dal.User
}

func (a *User) Query() (*schema.UserQueryResult, error) {
	result, err := a.UserDAL.Query()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *User) Get(id int) (*schema.User, error) {
	user, err := a.UserDAL.Get(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *User) Create(formItem *schema.UserForm) (*schema.User, error) {
	existsUsername, err := a.UserDAL.ExistsByUsername(formItem.Username)
	if err != nil {
		return nil, err
	} else if existsUsername {
		return nil, err
	}

	if formItem.Password == "" {
		formItem.Password = config.RootConfig.UserConfig.DefaultLoginPwd
	}
	user := &schema.User{
		CreatedAt: time.Now(),
	}
	if err := formItem.FillTo(user); err != nil {
		return nil, err
	}

	err = a.UserDAL.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *User) Update(id int, formItem *schema.UserForm) error {
	user, err := a.UserDAL.Get(id)
	if err != nil {
		return err
	} else if user == nil {
		return errors.NotFound("", "User not found")
	} else if user.Username != formItem.Username {
		existsUsername, err := a.UserDAL.ExistsByUsername(formItem.Username)
		if err != nil {
			return err
		} else if existsUsername {
			return errors.BadRequest("", "Username already exists")
		}
	}

	if err := formItem.FillTo(user); err != nil {
		return err
	}
	user.UpdatedAt = time.Now()

	if err := a.UserDAL.Update(user); err != nil {
		return err
	}

	return nil
}

func (a *User) Delete(id int) error {
	exists, err := a.UserDAL.Exists(id)
	if err != nil {
		return err
	} else if !exists {
		return errors.NotFound("", "User not found")
	}

	if err := a.UserDAL.Delete(id); err != nil {
		return err
	}

	return nil
}

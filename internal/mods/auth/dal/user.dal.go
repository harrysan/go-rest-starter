package dal

import (
	"finance-tracker/internal/mods/auth/schema"

	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUserDal(db *gorm.DB) *User {
	return &User{DB: db}
}

func (a *User) Query() (*schema.UserQueryResult, error) {
	var users schema.Users

	if err := a.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	queryResult := &schema.UserQueryResult{
		Data: users,
	}

	return queryResult, nil
}

func (a *User) Get(id int) (*schema.User, error) {
	item := new(schema.User)

	if err := a.DB.First(&item, id).Error; err != nil {
		return nil, err
	}

	return item, nil
}

func (a *User) Exists(id int) (bool, error) {
	var user schema.User
	result := a.DB.First(&user, id)

	if result.Error.Error() == "record not found" {
		return false, nil
	} else {
		return false, result.Error
	}
}

func (a *User) ExistsByUsername(username string) (bool, error) {
	var user schema.User
	result := a.DB.First(&user, "username=?", username)

	if result.Error.Error() == "record not found" {
		return false, nil
	} else {
		return false, result.Error
	}
}

func (a *User) Create(user *schema.User) error {
	return a.DB.Create(user).Error
}

func (a *User) Update(item *schema.User) error {
	exists, err := a.Exists(item.ID)
	if err != nil {
		return err
	}

	if exists {
		result := a.DB.Updates(item)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (a *User) Delete(id int) error {
	result := a.DB.Where("id=?", id).Delete(new(schema.User))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

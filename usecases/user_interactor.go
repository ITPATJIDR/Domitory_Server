package usecases

import (
	"Domitory_Server/database"
	"Domitory_Server/domain"
)

type UserInteractor struct {
	DB database.GormDB
}

func (ui UserInteractor) NewUser(u *domain.User) error {
	err := ui.DB.Create(&u)
	return err
}

func (ui UserInteractor) FindByID(id uint) (*domain.User, error) {
	var user *domain.User
	err := ui.DB.FindByID(&user, id)
	if user.ID != 0 {
		return user, err
	} else {
		return nil, err
	}
}

func (ui UserInteractor) FindByUsername(username string) (bool, error) {
	var user *domain.User
	err := ui.DB.FindByUserName(&user, username)
	if err != nil {
		return false, err
	}

	return user.ID != 0, err
}

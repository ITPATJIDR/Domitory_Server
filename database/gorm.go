package database

import (
	"Domitory_Server/domain"
	"errors"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormDB struct {
	db *gorm.DB
}

func NewGormStore() (*GormDB, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("DB_CONN")), &gorm.Config{})
	if err != nil {
		panic("Can't Connect to Database ")
	}

	db.AutoMigrate(&domain.User{}, &domain.Domitory{})

	return &GormDB{db: db}, nil
}

func (g *GormDB) Create(obj interface{}) error {
	return g.db.Create(obj).Error
}

func (g *GormDB) FindByID(obj interface{}, id uint) error {
	if err := g.db.Select("ID", "Username").Find(obj, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return nil
}

func (g *GormDB) FindByUserName(obj interface{}, username string) error {
	if err := g.db.Where("Username =  ?", username).Find(obj).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return nil
}

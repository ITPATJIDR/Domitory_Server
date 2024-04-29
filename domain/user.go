package domain

import "time"

type User struct {
	ID               uint      `gorm:"primarykey"`
	Username         string    `json:"username" binding:"required"`
	Password         string    `json:"password" binding:"required"`
	FirstName        string    `json:"first_name" binding:"required"`
	LastName         string    `json:"last_name" binding:"required"`
	Sex              string    `json:"sex" binding:"required"`
	Phone            string    `json:"phone" binding:"required"`
	NearByUniversity string    `json:"near_by_University" binding:"required"`
	Address          string    `json:"address" binding:"required"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime"`
}

type UserID struct {
	ID uint `json:"id" binding:"required"`
}

package models

import (
	"AuthenticationModule/utils"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey"`
	Email     string         `json:"email,omitempty" gorm:"type:VARCHAR(150);UNIQUE;NOT NULL" validate:"required,email"`
	Password  string         `json:"password" gorm:"type:VARCHAR(255);NOT NULL" validate:"required"`
	CreatedAt string         `json:"-" gorm:"type:datetime"`
	UpdatedAt *string        `json:"-" gorm:"type:datetime;NULL"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"type:datetime;NULL"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		return
	}
	return
}

func (user *User) Show(userID uint) error {
	return Db.Model(&User{}).
		Where("id = ?", userID).
		First(&user).Error
}

func (user *User) ShowByCredentials(email string) error {
	return Db.Model(&User{}).
		Where("email = ?", email).
		First(&user).Error
}

func (user *User) Store() error {
	return Db.Model(&User{}).
		Omit("id", "created_at", "updated_at", "deleted_at").
		Create(&user).Error
}

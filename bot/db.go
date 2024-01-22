package bot

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	QQGuid string
}

func (b *MinecraftBot) AddUser(msg, guid string) {
	user := User{Name: msg, QQGuid: guid}
	b.Db.Create(&user)
}

func (b *MinecraftBot) FirstUser(qQGuid string) User {
	var user User
	result := b.Db.Where("qq_guid = ?", qQGuid).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user
		} else {
			return user
		}
	} else {
		return user
	}
}

func (b *MinecraftBot) UpdateUser(msg string) {
	user := User{Name: msg}
	b.Db.Model(&user).Update("Age", 20)
}

func (b *MinecraftBot) DeleteUser(qQGuid string) User {
	var user User
	result := b.Db.Where("qq_guid = ?", qQGuid).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user
		} else {
			return user
		}
	} else {
		result = b.Db.Delete(&user)
		if result.Error != nil {
			return user
		}
		return user
	}
}

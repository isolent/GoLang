package database

import (
	"golang/assignment3/models"
	"github.com/pkg/errors"
)

func IsExistUserByName(name string) bool {
	var count = 0
	var user models.User

	GetDB().First(&user, models.User{Name: name}).Count(&count)
	return count == 1 && user.ID > 0
}	

func Add (bean interface{}) error{
	if !GetDB().NewRecord(bean){
		return errors.New("unable to create")
	}
	
	return GetDB().Create(bean).Error
}

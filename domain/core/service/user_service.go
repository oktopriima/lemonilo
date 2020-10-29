/*
 * Name : Okto Prima Jaya
 * Github : https://github.com/oktopriima
 * Email : octoprima93@gmail.com
 * Created At : 29/10/2020, 19:52
 * Copyright (c) 2020
 */

package service

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/oktopriima/lemonilo/domain/core/model"
)

type UserService interface {
	Create(user *model.User, tx *gorm.DB) (*model.User, error)
	Find(ID int) (*model.User, error)
	FindOneBy(criteria map[string]interface{}) (*model.User, error)
	FindBy(criteria map[string]interface{}) ([]*model.User, error)
	Update(user *model.User, tx *gorm.DB) error
	Delete(user *model.User, tx *gorm.DB) error
}

type userService struct {
	db *gorm.DB
}

func (u *userService) FindOneBy(criteria map[string]interface{}) (*model.User, error) {
	data := new(model.User)
	if err := u.db.Where(criteria).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (u *userService) Create(user *model.User, tx *gorm.DB) (*model.User, error) {
	db := tx.Create(&user)
	m := new(model.User)

	if err := db.Error; err != nil {
		return nil, err
	}

	byteData, err := json.Marshal(db.Value)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(byteData, &m); err != nil {
		return nil, err
	}

	return m, nil
}

func (u *userService) Find(ID int) (*model.User, error) {
	data := new(model.User)
	data.ID = ID
	if err := u.db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (u *userService) FindBy(criteria map[string]interface{}) ([]*model.User, error) {
	var data []*model.User
	if err := u.db.Where(criteria).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (u *userService) Update(user *model.User, tx *gorm.DB) error {
	if err := tx.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userService) Delete(user *model.User, tx *gorm.DB) error {
	if err := tx.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

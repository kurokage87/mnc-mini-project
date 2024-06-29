package service

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"mnc/config"
	"mnc/helper"
	"mnc/model"
	"time"
)

func Register(req model.RequestRegister) (data model.DataResponseRegister, err error) {
	db := config.DB
	user := model.User{
		UserId:      uuid.New(),
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		Pin:         req.Pin,
		CreatedDate: time.Now(),
	}
	if err = db.Create(&user).Error; err != nil {
		err = errors.New("phone number already register")
		return
	}
	data = model.DataResponseRegister{
		UserId:      user.UserId,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Pin:         user.Pin,
		CreatedDate: user.CreatedDate.Format("2006-1-2 15:04:05"),
	}
	return
}

func Login(req model.RequestLogin) (data model.Token, err error) {
	db := config.DB

	var user model.User
	if err = db.Where("phone_number = ? AND pin = ?", req.PhoneNumber, req.Pin).First(&user).Error; err != nil {
		err = errors.New("phone number and pin doesn't match")
		return
	}

	data.AccessToken, err = helper.GenerateToken(user.UserId, time.Hour*1) // Access token valid for 1 hour
	if err != nil {
		err = errors.New("failed generate access token")
		return
	}

	data.RefreshToken, err = helper.GenerateToken(user.UserId, time.Hour*24*7) // Refresh token valid for 7 days
	if err != nil {
		err = errors.New("failed generate refresh token")
		return
	}
	return
}

func Update(user_id string, req model.RequestUpdateProfile) (data model.DataResponseUpdateProfile, err error) {
	db := config.DB

	now := time.Now()
	dataUpdate := map[string]interface{}{
		"FirstName":  req.FirstName,
		"LastName":   req.LastName,
		"Address":    req.Address,
		"UpdateDate": now,
	}

	if err = db.Model(&model.User{}).Where("user_id = ?", user_id).Updates(dataUpdate).Error; err != nil {
		err = errors.New(fmt.Sprintf("failed update data for user_id %s and name %s", user_id, req.FirstName))
		return
	}

	data = model.DataResponseUpdateProfile{
		UserId:     user_id,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Address:    req.Address,
		UpdateDate: now.Format("2006-1-2 15:04:05"),
	}
	return
}

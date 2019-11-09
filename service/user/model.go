/**
 *******************************************slade********************************************
 * Copyright (c)  slade
 * Created by myGin.
 * User: 605724193@qq.com
 * Date: 2019/08/07
 * Time: 11:18
 ********************************************************************************************
 */

package user

import (
	"github.com/jinzhu/gorm"

	"myGin/service"
)

type User struct {
	UserId int64  `json:"user_id" gorm:"primary_key"`
	Name   string `json:"name" binding:"required"`
	Phone  string `json:"phone"`

	service.BaseModel
}

func (*User) BeforeCreate(scope *gorm.Scope) (err error) {
	_ = service.BeforeCreate(scope)

	return
}

func (*User) BeforeUpdate(scope *gorm.Scope) (err error) {
	_ = service.BeforeUpdate(scope)

	return
}

/**
 *******************************************slade********************************************
 * Copyright (c)  slade
 * Created by my-gin.
 * User: 605724193@qq.com
 * Date: 2019/08/07
 * Time: 11:18
 ********************************************************************************************
 */

package service

import (
	"time"

	"github.com/jinzhu/gorm"
)

type BaseModel struct {
	CreateUser string `json:"create_user"`
	CreateAt   int    `json:"create_at"`
	UpdateAt   int    `json:"update_at"`
}

func BeforeCreate(scope *gorm.Scope) (err error) {
	nowTimeStamp := int(time.Now().Unix())
	scope.SetColumn("CreateAt", nowTimeStamp)
	scope.SetColumn("UpdateAt", nowTimeStamp)

	return
}

func BeforeUpdate(scope *gorm.Scope) (err error) {
	nowTimeStamp := int(time.Now().Unix())
	scope.SetColumn("UpdateAt", nowTimeStamp)

	return
}

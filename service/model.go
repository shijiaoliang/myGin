/**
 *******************************************slade********************************************
 * Copyright (c)  slade
 * Created by myGin.
 * User: 605724193@qq.com
 * Date: 2019/08/07
 * Time: 11:18
 ********************************************************************************************
 */

package service

import "time"

type BaseModel struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

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

type BaseDto struct {
	CreateUser string `json:"create_user"`

	CreateAt int `json:"create_at"`
	UpdateAt int `json:"update_at"`

	StartCreateAt int `json:"start_create_at"`
	EndCreateAt   int `json:"end_create_at"`

	StartUpdateAt int `json:"start_update_at"`
	EndUpdateAt   int `json:"end_update_at"`

	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Sort    string `json:"sort"`
}

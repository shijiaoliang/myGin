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

type BaseDto struct {
	ID uint64 `json:"id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	StartCreatedAt time.Time `json:"start_created_at"`
	EndCreatedAt   time.Time `json:"end_created_at"`

	StartUpdatedAt time.Time `json:"start_updated_at"`
	EndUpdatedAt   time.Time `json:"end_updated_at"`

	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Sort    string `json:"sort"`
}

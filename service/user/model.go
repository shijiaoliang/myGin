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
	"myGin/service"
)

type User struct {
	service.BaseModel

	Name string `json:"name" binding:"required"`
	Age  int    `json:"age"`
}

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

type Dto struct {
	service.BaseDto

	Name string `json:"name"`
	Age  int    `json:"age"`
}

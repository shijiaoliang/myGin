/**
 *******************************************slade********************************************
 * Copyright (c)  slade
 * Created by my-gin.
 * User: 605724193@qq.com
 * Date: 2019/08/07
 * Time: 11:18
 ********************************************************************************************
 */

package util

import (
	"gopkg.in/go-playground/validator.v9"
)

// 全局验证实例
var Validate *validator.Validate

func init()  {
	Validate = validator.New()
}

/**
 *******************************************slade********************************************
 * Copyright (c)  slade
 * Created by myGin.
 * User: 605724193@qq.com
 * Date: 2019/08/07
 * Time: 11:18
 ********************************************************************************************
 */

package router

import (
	"github.com/gin-gonic/gin"

	"myGin/util"
)

func InitIndexRouter(r *gin.Engine) {
	indexRouter := r.Group("/index")

	//[/index/index]
	indexRouter.GET("index", func(c *gin.Context) {
		util.ResSuccess(c, "Hello Kitty ^_^", "")
	})
}

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
	"github.com/pingcap/errors"

	"myGin/db"
	"myGin/service"
	"myGin/service/user"
	"myGin/util"
)

func InitUserRouter(r *gin.Engine) {
	userRouter := r.Group("/user")

	//[/user/add]
	userRouter.Any("add", func(c *gin.Context) {
		var u user.User

		if err := c.ShouldBindJSON(&u); err != nil {
			err = errors.WithStack(service.ErrParamError)
			util.CheckErr(err, c)

			return
		}

		db.Mysql.Create(&u)

		util.ResSuccess(c, u, "")
	})
}

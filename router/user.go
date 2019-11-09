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

	//[/user/list]
	userRouter.POST("list", func(c *gin.Context) {
		var dto user.User

		if err := c.ShouldBindJSON(&dto); err != nil {
			err = errors.WithStack(service.ErrParamError)
			util.CheckErr(err, c)

			return
		}

		query := db.Mysql.Where(&dto)

		//if dto.StartCreatedAt.Second() > 0 {
		//	query = query.Where("created_at >= ?", dto.StartCreatedAt)
		//}
		//if dto.EndCreatedAt.Second() > 0 {
		//	query = query.Where("created_at <= ?", dto.EndCreatedAt)
		//}
		//if dto.StartUpdatedAt.Second() > 0 {
		//	query = query.Where("updated_at >= ?", dto.StartUpdatedAt)
		//}
		//if dto.StartUpdatedAt.Second() > 0 {
		//	query = query.Where("updated_at <= ?", dto.StartUpdatedAt)
		//}

		var users []user.User
		var totalCount int64
		query.Find(&users).Count(&totalCount)

		pagination := util.Pagination{TotalCount:totalCount}

		util.ResSuccessList(c, users, pagination, "")
	})

	//[/user/add]
	userRouter.POST("add", func(c *gin.Context) {
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

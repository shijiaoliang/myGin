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
		var dto user.Dto

		if err := c.ShouldBindJSON(&dto); err != nil {
			err = errors.WithStack(service.ErrParamError)
			util.CheckErr(err, c)

			return
		}

		query := dto.BaseQuery()
		query = query.Table("user")
		if dto.Name != "" {
			query = query.Where("name = ?", dto.Name)
		}
		if dto.Age > 0 {
			query = query.Where("age = ?", dto.Age)
		}

		//users
		var users []user.User

		//分页
		pagination := &util.Pagination{}
		var totalCount int64
		if dto.DoPage {
			query.Count(&totalCount)
			pagination = dto.BasePagination(query, totalCount)

			query = query.Offset(pagination.Page -1)
			query = query.Limit(pagination.PerPage)

			query.Find(&users)
		} else {
			query.Find(&users).Count(&totalCount)
		}


		//res
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

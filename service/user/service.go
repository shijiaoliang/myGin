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
	"github.com/jinzhu/gorm"

	"myGin/db"
	"myGin/util"
	"sync"
)

type userService struct {
	mutex *sync.Mutex
}

var Service = &userService{
	mutex: &sync.Mutex{},
}

func (s *userService) List(dto *Dto) (ret []*User, pagination util.Pagination) {
	query := dto.BaseQuery()
	query = query.Table("user")

	if dto.Name != "" {
		query = query.Where("name = ?", dto.Name)
	}
	if dto.Age > 0 {
		query = query.Where("age = ?", dto.Age)
	}

	//分页
	var totalCount int64
	if dto.DoPage {
		query.Count(&totalCount)
		pagination = dto.BasePagination(query, totalCount)

		query = query.Offset(pagination.Page - 1)
		query = query.Limit(pagination.PerPage)
	}

	query.Find(&ret)

	return
}

func (s *userService) Info(id int64) (ret User, err error) {
	if err := db.Mysql.First(&ret, id).Error; gorm.IsRecordNotFoundError(err) {
		//log.Logger.Error("user info failed: " + err.Error())

		err = ErrUserNotExist
		return ret, err
	}

	return
}

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

import (
	"github.com/jinzhu/gorm"
	"time"

	"myGin/db"
	"myGin/util"
)

type BaseDto struct {
	ID uint64 `json:"id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	StartCreatedAt time.Time `json:"start_created_at"`
	EndCreatedAt   time.Time `json:"end_created_at"`

	StartUpdatedAt time.Time `json:"start_updated_at"`
	EndUpdatedAt   time.Time `json:"end_updated_at"`

	DoPage  bool  `json:"do_page"`
	Page    int64 `json:"page"`
	PerPage int64 `json:"per_page"`

	Sort string `json:"sort"`
}

//BaseQuery
func (d *BaseDto) BaseQuery() *gorm.DB {
	query := db.Mysql

	if d.ID > 0 {
		query = query.Where("id = ?", d.ID)
	}

	if d.StartCreatedAt.Second() > 0 {
		query = query.Where("created_at >= ?", d.StartCreatedAt)
	}
	if d.EndCreatedAt.Second() > 0 {
		query = query.Where("created_at <= ?", d.EndCreatedAt)
	}

	if d.StartUpdatedAt.Second() > 0 {
		query = query.Where("updated_at >= ?", d.StartUpdatedAt)
	}
	if d.StartUpdatedAt.Second() > 0 {
		query = query.Where("updated_at <= ?", d.StartUpdatedAt)
	}

	return query
}

//BasePagination
func (d *BaseDto) BasePagination(query *gorm.DB, totalCount int64) *util.Pagination {
	pagination := &util.Pagination{}

	if d.DoPage {
		page := d.Page
		if page <= 0 {
			page = 1
		}

		perPage := d.PerPage
		if perPage <= 0 {
			perPage = 1
		}

		pagination = &util.Pagination{
			TotalCount: totalCount,
			Page:       page,
			PerPage:    perPage,
		}
	}

	return pagination
}



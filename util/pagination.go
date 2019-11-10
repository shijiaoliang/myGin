/**
 *******************************************slade********************************************
 * Copyright (c)  slade
 * Created by myGin.
 * User: 605724193@qq.com
 * Date: 2019/08/07
 * Time: 11:18
 ********************************************************************************************
 */

package util

//分页
type Pagination struct {
	TotalCount int64 `json:"total_count"`
	Page       int64 `json:"page"`
	PerPage    int64 `json:"per_page"`
}

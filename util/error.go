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
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"my-gin/log"
)

//appError 系统通用error
type AppError struct {
	Err  error
	Code int
}

func (a *AppError) Error() string {
	return a.Err.Error()
}

func (a *AppError) GetCode() int {
	return a.Code
}

func AppErrorNew(err error, code int) *AppError {
	return &AppError{
		err,
		code,
	}
}

func CheckErr(err error, c *gin.Context) {
	if err != nil {
		//记录日志
		log.Logger.Error(fmt.Sprintf("%+v", err))

		//res abort
		msg := ""

		switch err := errors.Cause(err).(type) {
		case *AppError:
			msg = err.Error()
		default:
			msg = "system error"
		}

		c.AbortWithStatusJSON(http.StatusOK, &Res{"", ERROR, msg})
		return
	}
}

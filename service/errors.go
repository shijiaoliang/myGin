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
	"errors"

	"myGin/util"
)

var (
	ErrParamError = util.AppErrorNew(errors.New("params error"), util.ERROR)
)

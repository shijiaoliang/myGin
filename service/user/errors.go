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
	"errors"

	"myGin/util"
)

var (
	ErrUserNotExist = util.AppErrorNew(errors.New("user not exist"), util.ERROR)
)

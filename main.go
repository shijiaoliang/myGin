/**
 *******************************************slade********************************************
 * Copyright (c)  slade
 * Created by myGin.
 * User: 605724193@qq.com
 * Date: 2019/08/07
 * Time: 11:18
 ********************************************************************************************
 */

package main

import (
	"time"

	"github.com/gin-gonic/gin"

	"myGin/config"
	"myGin/db"
	"myGin/log"
	"myGin/middleware"
	"myGin/router"
	"myGin/util"
)

func main() {
	defer db.Mysql.Close()

	r := gin.New()

	//=====config=====
	// config init

	//=====mode=====
	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(config.AppConf.Model)

	//=====log=====
	runtimeDirectory := "runtime"
	if !util.FileExist(runtimeDirectory) {
		_ = util.CreateFile(runtimeDirectory)
	}
	log.ConfigLogger(log.Logger, runtimeDirectory, "app.log", time.Hour*360, time.Hour)

	//=====middleware=====
	r.Use(middleware.Logger(log.Logger))
	r.Use(gin.Recovery())

	//=====route=====
	router.InitRouter(r)

	//=====run=====
	_ = r.Run(config.AppConf.Address)
}

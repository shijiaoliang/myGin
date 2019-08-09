/**
 *******************************************slade********************************************
 * Copyright (c)  slade
 * Created by my-gin.
 * User: 605724193@qq.com
 * Date: 2019/08/07
 * Time: 11:18
 ********************************************************************************************
 */

package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"my-gin/log"
	"my-gin/util"
)

type AppConfig struct {
	Model   string `json:"model"`
	Address string `json:"address"`
}

var (
	AppConf *AppConfig
)

func init() {
	// default config
	AppConf = new(AppConfig)
	AppConf.Model = gin.ReleaseMode
	AppConf.Address = "8989"

	// 如果存在配置json文件 以配置文件为准
	if util.FileExist("conf/app.json") {
		viper.SetConfigName("app")
		viper.AddConfigPath("conf")
		viper.AddConfigPath(".")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		viper.Unmarshal(&AppConf)
	}

	log.Logger.Info("config loaded")
}

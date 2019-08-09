package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"

	"my-gin/log"
)

var Mysql *gorm.DB

func init() {
	//e.g: root:123456@tcp(192.168.8.100:3306)/ceping?charset=utf8&parseTime=True&loc=Local
	dsn := viper.GetString("mysql.user") + ":" + viper.GetString("mysql.password") + "@tcp(" +
		viper.GetString("mysql.host") + ":" + viper.GetString("mysql.port") + ")/" +
		viper.GetString("mysql.db") + "?charset=" +
		viper.GetString("mysql.charset") + "&parseTime=True&loc=Local"

	var err error
	Mysql, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("Mysql connect error: %s \n", err))
	}

	//连接池配置
	//闲置的连接数 | 最大打开连接数, 默认值为0表示不限制
	Mysql.DB().SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
	Mysql.DB().SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))

	//全局禁用表名复数
	Mysql.SingularTable(true)

	//日志
	Mysql.SetLogger(log.Logger)
	Mysql.LogMode(true)

	log.Logger.Info("mysql loaded")
}

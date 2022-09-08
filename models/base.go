package models

import (
	"fmt"

	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func Init()  {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		_ = fmt.Sprintf("mysql error: %w", err)
	}

	databaseSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		viper.Get("mysql.username"),
		viper.Get("mysql.password"),
		viper.Get("mysql.host"),
		viper.Get("mysql.port"),
		viper.Get("mysql.database"),
	)

	err = orm.RegisterDataBase("default", "mysql", databaseSource)
	if err != nil {
		_ = fmt.Sprintf("mysql error: %w", err)
	}

}
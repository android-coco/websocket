package db

import (
	"websocket/common"
	"websocket/module"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var goOrmDefaultDB *gorm.DB
var goOrmInlineActionDB *gorm.DB

type logger struct{}

func (logger) Print(v ...interface{}) {
	common.Logger.Info(v)
}
func InitDb() module.Error {
	err := initDefaultDb()
	if err.ErrCode != common.SUCCESS_CODE {
		common.Logger.Error("init mysql1 failed.", err)
		return err
	}

	err = initInlineActionDb()
	if err.ErrCode != common.SUCCESS_CODE {
		common.Logger.Error("init mysql2 failed.", err)
		return err
	}
	common.Logger.Info("db init success.")
	return module.Error{ErrCode: common.SUCCESS_CODE, ErrMsg: nil}
}

func initInlineActionDb() module.Error {
	var err error
	addr := common.Config.Section("mysql").Key("addr2").MustString("")
	user := common.Config.Section("mysql").Key("user2").MustString("")
	passWd := common.Config.Section("mysql").Key("passwd2").MustString("")
	db := common.Config.Section("mysql").Key("db2").MustString("")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		user, passWd, addr, db)
	goOrmInlineActionDB, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		// g orm 会自己 ping 一次 DB
		common.Logger.Error("sql.Ping command failed. err:", err,
			" data_source_name: ", dataSourceName)
		return module.Error{ErrCode: common.ERROR_MYSQL_COMMAND_PING, ErrMsg: err}
	}

	goOrmInlineActionDB.LogMode(true).SetLogger(logger{})
	return module.Error{ErrCode: common.SUCCESS_CODE, ErrMsg: nil}
}

func initDefaultDb() module.Error {
	var err error
	addr := common.Config.Section("mysql").Key("addr").MustString("")
	user := common.Config.Section("mysql").Key("user").MustString("")
	passWd := common.Config.Section("mysql").Key("passwd").MustString("")
	db := common.Config.Section("mysql").Key("db").MustString("")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		user, passWd, addr, db)
	goOrmDefaultDB, err = gorm.Open("mysql", dataSourceName)
	if err != nil {
		// g orm 会自己 ping 一次 DB
		common.Logger.Error("sql.Ping command failed. err:", err,
			" data_source_name: ", dataSourceName)
		return module.Error{ErrCode: common.ERROR_MYSQL_COMMAND_PING, ErrMsg: err}
	}

	goOrmDefaultDB.LogMode(true).SetLogger(logger{})
	return module.Error{ErrCode: common.SUCCESS_CODE, ErrMsg: nil}
}

func GetInlineActionDB() *gorm.DB {
	return goOrmDefaultDB
}

func GetDefaultDB() *gorm.DB {
	return goOrmDefaultDB
}

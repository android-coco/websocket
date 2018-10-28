package common

import (
	"websocket/module"
	"fmt"
	"github.com/cihub/seelog"
)

var Logger seelog.LoggerInterface

func loadAppConfig() {
	logConfigFile := Path + Config.Section("service").Key("log_config_file").MustString(CONFIG_DEFAULT_LOG_CONFIG_FILE)
	logger, err := seelog.LoggerFromConfigAsFile(logConfigFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	UseLogger(logger)
}

func InitLog() (Err module.Error) {
	DisableLog()
	loadAppConfig()
	Logger.Info("log init success.")
	return module.Error{ErrCode:SUCCESS_CODE, ErrMsg:nil}
}

// DisableLog disables all library log output
func DisableLog() {
	Logger = seelog.Disabled
}

// UseLogger uses a specified seelog.LoggerInterface to output library log. // Use this func if you are using Seelog logging system in your app.
func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}

package common

import (
	"eospart_websocket/module"
	"fmt"
	"os"
)

func InitPid() module.Error {
	var err error
	pidFile := Path + Config.Section("service").Key("pid_file").MustString(CONFIG_DEFAULT_PID_FILE)
	var f *os.File
	f, err = os.OpenFile(pidFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		Logger.Error("open pid file failed")
		return module.Error{ErrCode: ERROR_SYSTEM, ErrMsg: err}
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("%d", os.Getpid()))
	Logger.Info("pid file init success")
	return module.Error{ErrCode: SUCCESS_CODE, ErrMsg: nil}
}

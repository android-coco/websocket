package common

import (
	"fmt"
	"github.com/go-ini/ini"
	"os"
	"path/filepath"
	"eospart_websocket/module"
)

var Config *ini.File
var Path string

func InitConfig() module.Error {
	var err error
	Path, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	Config, err = ini.Load(Path + "/../config/config.ini")
	if err != nil {
		fmt.Printf("Init Config Load failed. err:%s\n", err)
		return module.Error{ErrCode: ERROR_SYSTEM, ErrMsg: err}
	}
	return module.Error{ErrCode: SUCCESS_CODE, ErrMsg: nil}
}

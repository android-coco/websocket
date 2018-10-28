package main

import (
	"eospart_websocket/cache"
	"eospart_websocket/common"
	"eospart_websocket/db"
	"eospart_websocket/httpserver"
	"eospart_websocket/module"
	"fmt"
	"google.golang.org/appengine/memcache"
)

func main() {
	var err module.Error

	// 配置文件
	err = common.InitConfig()
	if err.ErrCode != common.SUCCESS_CODE {
		fmt.Println("init config failed. err: ", err)
		return
	}

	// log 日志
	err = common.InitLog()
	if err.ErrCode != common.SUCCESS_CODE {
		common.Logger.Error("init log failed.", err)
		return
	}

	// 环境变量
	err = common.InitEnv()
	if err.ErrCode != common.SUCCESS_CODE {
		common.Logger.Error("init env failed.", err)
		return
	}

	// 进程id
	err = common.InitPid()
	if err.ErrCode != common.SUCCESS_CODE {
		common.Logger.Error("init pid failed.", err)
		return
	}

	// 数据库
	err = db.InitDb()
	if err.ErrCode != common.SUCCESS_CODE {
		common.Logger.Error("init db failed.", err)
		return
	}

	// redis
	err = cache.InitRedis()
	if err.ErrCode != common.SUCCESS_CODE {
		common.Logger.Error("init redis failed.", err)
		return
	}

	// memCache
	err = cache.InitMemCache()
	if err.ErrCode != common.SUCCESS_CODE {
		common.Logger.Error("init memcache failed.", err)
		return
	}
	//cache.GetMemCache().Set(&memcache.Item{Key:"11",Value:[]byte("fasdfa"),Expiration:1})
	// 启动web socket
	err = httpserver.Run()
	if err.ErrCode != common.SUCCESS_CODE {
		common.Logger.Error("init web socket failed.", err)
		return
	}

}

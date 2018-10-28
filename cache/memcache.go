package cache

import (
	"eospart_websocket/common"
	"eospart_websocket/module"
	"errors"
	"github.com/bradfitz/gomemcache/memcache"
	"strings"
)

var memCacheDB *memcache.Client

func InitMemCache() module.Error {
	cacheSvrList := strings.Split(common.CacheServerList, ",")
	memCacheDB = memcache.New(cacheSvrList...)
	if memCacheDB == nil {
		return module.Error{ErrCode: common.ERROR_MEMCACHE_COMMAND_PING, ErrMsg: errors.New("")}
	}
	return module.Error{ErrCode: common.SUCCESS_CODE, ErrMsg: nil}
}

func GetMemCache() *memcache.Client {
	return memCacheDB
}

package httpserver

import (
	"websocket/common"
	"websocket/httpserver/api"
	"websocket/module"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() module.Error {
	isDebugMode := common.Config.Section("gin").Key("debug_mode").MustBool(true)
	//wssCert := common.Path + common.Config.Section("service").Key("wss_cert_file").MustString(common.CONFIG_DEFAULT_WSS_CERT_CONFIG_FILE)
	//wssKey := common.Path + common.Config.Section("service").Key("wss_key_file").MustString(common.CONFIG_DEFAULT_WSS_KEY_CONFIG_FILE)
	if !isDebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(initCorConf()))
	api.InitRoutes(router)
	common.Logger.Info("web socket init success on post" + common.ServerAddr)
	//https
	//err := router.RunTLS(common.ServerAddr,wssCert,wssKey)
	err := router.Run(common.ServerAddr)
	if err != nil {
		common.Logger.Errorf("fail to start web service: %s", err.Error())
		return module.Error{ErrCode: common.ERROR_SYSTEM, ErrMsg: err}
	}
	return module.Error{ErrCode: common.SUCCESS_CODE, ErrMsg: nil}
}

func initCorConf() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"https://eospark.com",
		"http://localhost:8888",
	}
	config.AllowOriginFunc = func(origin string) bool {
		return strings.HasSuffix(origin, ".eospark.com") ||
			strings.HasSuffix(origin, "//eospark.com") ||
			strings.HasSuffix(origin, ".blockabc.com") ||
			strings.HasSuffix(origin, "//blockabc.com")
	}
	return config
}

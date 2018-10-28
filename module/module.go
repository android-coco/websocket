package module

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var UpGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { //允许跨域
		return true
	},
}

type Error struct {
	ErrCode int64
	ErrMsg  error
}

type APIResponse struct {
	ErrNo  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data,omitempty"`
}

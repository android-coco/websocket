package echo

import (
	"websocket/common"
	"websocket/model"
	"websocket/module"
	"websocket/cache"
	"websocket/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

//客户端信息
type Client struct {
	ApiKey      string          //api key
	Ip          string          //客户端ip地址
	Context     []string        //订阅内容
	Conn        *websocket.Conn //客户端连接对象
	MessageChan chan interface{}
	Done        chan struct{}
}

// 记录IP 和api key的关系  后续加入缓存中 [apiKey:ip]
//var ipApiKeyMap sync.Map
// 记录客户端信息  [ip:ws]
//var wsMap  sync.Map
// 记录客户的Chan    [ip,chan interface{}]
//var clientMessageChan sync.Map

// 发送数据源
func sendMsg(currentClient *Client) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-currentClient.Done:
			releaseClient(currentClient)
			return
		case <-ticker.C:
			// TODO 解析数据
			yh := model.User{Name: "you hao", Age: 30, Ip: currentClient.Ip}
			currentClient.MessageChan <- yh
		}
	}
}

// 内联数据
var inlineChan chan []byte

func inlineData() {
	//inlineChan = make(chan []byte)
	//context, _ := gozmq.NewContext()
	//socket, _ := context.NewSocket(gozmq.SUB)
	//defer context.Close()
	//defer socket.Close()
	//socket.Connect("tcp://127.0.0.1:5556")
	//socket.SetSubscribe("")
	//for {
	//	if msg, err := socket.Recv(0); err != nil {
	//		inlineChan <- msg
	//	}
	//}
}

// 链上数据  区块
var blockDataChan chan []byte

func blockData() {
	//TODO 1.获取基本信息
	//TODO curl http://127.0.0.1:8888/v1/chain/get_info  得到最新区块
	//TODO 2. 获取指定区块信息
	//curl http://127.0.0.1:8888/v1/chain/get_block -X POST -d '{"block_num_or_id":"32111"}'
}

//var isClose chan int
// 收数据  监听客户端状态
func readMsg(currentClient *Client, ) {
	for {
		_, _, err := currentClient.Conn.ReadMessage()
		if ce, ok := err.(*websocket.CloseError); ok {
			switch ce.Code {
			case websocket.CloseNormalClosure, websocket.CloseGoingAway, websocket.CloseNoStatusReceived:
				// TODO 清除一些缓存 和客户端信息
				releaseClient(currentClient)
				common.Logger.Infof("客户端SendMsg %v Bay Bay bye err: %v", currentClient, err)
			}
			break
		}

	}
}

//模拟多个用户不同要求
var client int

func Ws(c *gin.Context) {
	ws, err := module.UpGrader.Upgrade(c.Writer, c.Request, nil)
	apiKey := c.GetHeader("apikey")
	if err != nil {
		common.Logger.Error("upgrade:", err)
		return
	}
	if apiKey == "" {
		ws.WriteJSON(module.APIResponse{
			ErrNo:  http.StatusForbidden,
			ErrMsg: http.StatusText(http.StatusForbidden)})
		return
	}
	// ip地址
	var ip = ws.RemoteAddr().String()
	//TODO 判断 apiKey 的合法性 和订阅信息  一个ip 一个api key
	redisDb := cache.GetRedisDb()

	apiKeyCase, err := redisDb.Get(utils.GetMd5([]byte(ip))).Result()

	if err != nil {
		redisDb.Set(utils.GetMd5([]byte(ip)), apiKey, 0)
		client++
	} else if apiKeyCase == apiKey {
		common.Logger.Infof("current api key is connection. api key is: %s", apiKey)
		return
	}
	//if _, ok := ipApiKeyMap.Load(apiKey); !ok {
	//	ipApiKeyMap.Store(ip, apiKey)
	//	client ++
	//} else {
	//	return
	//}

	// 数据Chan
	var messageChan = make(chan interface{})
	// 结束标志Chan
	var done = make(chan struct{})
	// 如果这个客户来过 适用第一次的通道
	//if v, ok := clientMessageChan.Load(ip); ok {
	//	messageChan = v.(chan interface{})
	//} else {
	//	clientMessageChan.Store(ip, messageChan)
	//	client ++
	//}

	// 当前客户端
	var currentClient Client
	if client%2 == 0 {
		//v, ok := wsMap.Load(key)
		//if !ok {
		//
		//	wsMap.Store(key,currentClient)
		//}else{
		//	currentClient = v.(Client)
		//}
		currentClient = Client{
			ApiKey:      apiKey,
			Ip:          ip,
			Conn:        ws,
			Context:     []string{"account"}, // 这个数据要配置,然后客户端传过来
			MessageChan: messageChan,
			Done:        done,
		}
	} else {
		//v, ok := wsMap.Load(key)
		//if !ok {
		//
		//	wsMap.Store(key,currentClient)
		//}else{
		//	currentClient = v.(Client)
		//}
		currentClient = Client{
			ApiKey:      apiKey,
			Ip:          ip,
			Conn:        ws,
			Context:     []string{"all"}, // 这个数据要配置,然后客户端传过来
			MessageChan: messageChan,
			Done:        done,
		}
	}

	common.Logger.Infof("当前链接数量:%d\n 连接客户端：%v", client, currentClient)
	go server(&currentClient)

}

func server(currentClient *Client) {
	go sendMsg(currentClient)
	defer close(currentClient.Done)
	for {
		// 模拟数据源  可以有多个
		yh, ok := <-currentClient.MessageChan
		if !ok {
			break
		}
		//发送数据的时候判断要发送哪些内容
		//1,取出api key
		if currentClient.ApiKey == "" {
			//api key不存在  不发送数据
			break
		}
		context := currentClient.Context
		//TODO 根据context 过滤数据
		//模拟过滤
		if context[0] == "account" {
			yh = model.User{Name: "lyl", Age: 20, Ip: currentClient.Ip}
		}
		err := currentClient.Conn.WriteJSON(yh)
		if err != nil {
			// TODO 某个客户端异常退出
			common.Logger.Errorf("客户端 %v Bay Bay bye err: %v", currentClient, err)
			break
		}
	}
}

func releaseClient(currentClient *Client) {
	//apiKey, _ := ipApiKeyMap.Load(currentClient.Ip)
	//ipApiKeyMap.Delete(apiKey)
	redisDb := cache.GetRedisDb()
	redisDb.Del(utils.GetMd5([]byte(currentClient.ApiKey)))
	//clientMessageChan.Delete(currentClient.Ip)
	//客户端数量--
	client--
	defer close(currentClient.MessageChan)
	defer currentClient.Conn.Close()
	common.Logger.Infof("客户端 %v Bay Bay ", currentClient)
}

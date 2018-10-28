package main

import (
	"eospart_websocket/utils"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

//wss://mainnet.eos.dfuse.io/v1/stream?token=eyJ..YOURTOKENHERE...
//var addr = flag.String("addr", "47.75.218.79:8205", "http service address")
var addr = flag.String("addr", "127.0.0.1:8888", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	//createClient()
	for i := 0; i < 1000000; i++ {
		go createClient()
		time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(100 * time.Hour)
}

func createClient() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	fmt.Printf("connecting to %s", u.String())

	s2 := rand.NewSource(time.Now().UnixNano()) //同前面一样的种子
	r2 := rand.New(s2)
	sj := utils.GetMd5([]byte(strconv.Itoa(r2.Intn(10000))))
	c, _, err := websocket.DefaultDialer.Dial(u.String(), map[string][]string{"apikey": {sj + "e9564ebc3289b7a14551baf8ad5ec60a"}})
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			//err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			//if err != nil {
			//	log.Println("write:", err)
			//	return
			//}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "bay bay"))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}

package clipboard

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"job-log/model/common/response"
	"job-log/model/entity"
	"log"
	"net/http"
	"sync"
)

// 定义客户端连接管理器
type clientManager struct {
	clients    map[*websocket.Conn]bool // 所有连接的客户端
	broadcast  chan []byte              // 广播消息通道
	register   chan *websocket.Conn     // 注册新客户端
	unregister chan *websocket.Conn     // 注销客户端
	sync.Mutex                          // 互斥锁，确保并发安全
}

// 全局客户端管理器
var manager = clientManager{
	clients:    make(map[*websocket.Conn]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *websocket.Conn),
	unregister: make(chan *websocket.Conn),
}

func (manager *clientManager) start() {
	for {
		select {
		case client := <-manager.register:
			// 注册新客户端
			manager.Lock()
			manager.clients[client] = true
			manager.Unlock()
			log.Println("新客户端已连接，当前客户端数量:", len(manager.clients))

		case client := <-manager.unregister:
			// 注销客户端
			manager.Lock()
			if _, ok := manager.clients[client]; ok {
				delete(manager.clients, client)
				client.Close()
			}
			manager.Unlock()
			log.Println("客户端已断开，当前客户端数量:", len(manager.clients))

		case message := <-manager.broadcast:
			// 广播消息给所有客户端
			manager.Lock()
			for client := range manager.clients {
				err := client.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Println("发送消息失败:", err)
					client.Close()
					delete(manager.clients, client)
				}
			}
			manager.Unlock()
		}
	}
}

func init() {
	go manager.start()
}

var wsUpgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源的连接
	},
}

func Ws(c *gin.Context) {
	conn, err := wsUpgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		response.FailWithMessage("ws连接失败", c)
		return
	}
	defer func() {
		manager.unregister <- conn
		conn.Close()
	}()
	manager.register <- conn
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		cs.Update(entity.Clipboard{Content: string(message)})
		manager.broadcast <- message
	}
}

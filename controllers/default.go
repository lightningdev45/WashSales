package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"log"
	"time"
	"net/http"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
  CheckOrigin: func(r *http.Request) bool { return true },
}




type hub struct {
	// Registered connections.
	connections map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection

	roomId string
}

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	uname string
	h *hub
}


func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					close(c.send)
					delete(h.connections, c)
				}
			}
		}
	}
}

var rooms []hub


type MainController struct {
	beego.Controller
}

type WebSocketController struct {
	beego.Controller
}

func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (c *connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
			log.Println(string(message))
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *connection) readPump() {
	defer func() {
		c.h.unregister <- c
		c.ws.Close()
		log.Println("connectionClosed")
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		c.h.broadcast <- message
	}
}


func (this *WebSocketController) Get() {
	//uname := this.GetString("uname")

	uname:="Will"
	roomId:=this.Ctx.Input.Param(":roomId")
	log.Println(roomId)
	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	roomFound:=false
	var newHub hub;
	for _,v:=range rooms{
		if v.roomId==roomId{
			roomFound=true
			newHub=v
			break
		}
	}
	if roomFound{
	}else{
			newHub=hub{
				broadcast:   make(chan []byte),
				register:    make(chan *connection),
				unregister:  make(chan *connection),
				connections: make(map[*connection]bool),
				roomId: roomId,
			}
			rooms=append(rooms,newHub)
			go newHub.run()
	}
	log.Println(newHub)
  log.Println("main")

	c := &connection{send: make(chan []byte, 256), ws: ws,uname:uname,h:&newHub}
	newHub.register <- c
  log.Println("hello")
	go c.writePump()
	c.readPump()
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

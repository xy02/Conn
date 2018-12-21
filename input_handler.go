package conn

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InputHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	clientInfo := q.Get("i")
	versionCode := q.Get("v")

	fmt.Println(clientInfo, versionCode, r.URL.Path)
	fmt.Println("RawPath", r.URL.RawPath)
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		if mt != websocket.BinaryMessage {
			log.Println("not BinaryMessage")
			break
		}
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

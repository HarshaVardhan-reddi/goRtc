package socket

import (
	"net/http"
	"github.com/gorilla/websocket"
)

var websocketUpgrader websocket.Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all origins for development. In production, you should validate this!
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


func UpgradeCurrentRequestToWebrequest(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	cusotmheaders := http.Header{}
	conn, err := websocketUpgrader.Upgrade(w,r,cusotmheaders)
	if(err != nil){
		return nil, err
	}
	return conn, nil
}
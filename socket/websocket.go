package socket

import (
	"net/http"
	"github.com/gorilla/websocket"
)

var websocketUpgrader websocket.Upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}


func UpgradeCurrentRequestToWebrequest(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	cusotmheaders := http.Header{}
	conn, err := websocketUpgrader.Upgrade(w,r,cusotmheaders)
	if(err != nil){
		return nil, err
	}
	return conn, nil
}
package socket

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

func ProcessPackets(webSocketConn *websocket.Conn) {
	for {
		pkt := Packet{}
		_, messageInByte, err := webSocketConn.ReadMessage()
		if err != nil {
			log.Println("error in reading packets:", err)
			break
		}

		log.Println("message in string", string(messageInByte))

		if errInUnmarshal := json.Unmarshal(messageInByte, &pkt); errInUnmarshal != nil {
			log.Println("error in parsing packets:", errInUnmarshal)
			continue
		}

		if errInPacket := pkt.validatePacket(); errInPacket != nil {
			log.Println("invalid packet:", errInPacket)
			continue
		}

		pkt.DispatchPacket()
	}
}

func WritePacket(webSocketConn *websocket.Conn, pkt Packet) error {
	data, err := json.Marshal(pkt)
	if err != nil {
		return err
	}
	return webSocketConn.WriteMessage(websocket.TextMessage, data)
}
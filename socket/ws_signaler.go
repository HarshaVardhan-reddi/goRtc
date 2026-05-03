package socket

import "github.com/gorilla/websocket"

// WSSignaler is a WebSocket implementation of the rtc.Signaler interface.
type WSSignaler struct {
	Conn *websocket.Conn
}

// Send wraps the raw signaling payload into a Packet and sends it via WebSocket.
func (s *WSSignaler) Send(payload []byte) error {
	pkt := Packet{
		Type:       "rtc",
		RawMessage: payload,
	}
	return WritePacket(s.Conn, pkt)
}

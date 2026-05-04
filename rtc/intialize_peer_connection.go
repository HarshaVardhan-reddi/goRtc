package rtc

import (
	"encoding/json"
	"log"

	"github.com/pion/webrtc/v4"
)

var iceServer webrtc.ICEServer = webrtc.ICEServer{URLs: []string{"stun:stun.l.google.com:19302"}}

func InitializePeerConnection(signaler Signaler) (*webrtc.PeerConnection, error) {
	config := webrtc.Configuration{ICEServers: []webrtc.ICEServer{iceServer}}
	rtcConn, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return nil, err
	}
	rtcConn.OnICECandidate(func(i *webrtc.ICECandidate) {
		if i == nil {
			return
		}
		rawCandidate := i.ToJSON()
		payload, err := json.Marshal(rawCandidate)
		if err != nil {
			log.Println(err)
			return
		}
		if err := signaler.Send(payload); err != nil {
			log.Println("Signaling failed:", err)
		}
	})
	rtcConn.OnDataChannel(func(dc *webrtc.DataChannel) {
		log.Printf("New DataChannel: %s", dc.Label())

		dc.OnOpen(func() {
			log.Printf("DataChannel %s is open", dc.Label())
		})

		dc.OnMessage(func(msg webrtc.DataChannelMessage) {
			log.Printf("Message from peer on %s: %s", dc.Label(), string(msg.Data))
		})

		dc.OnClose(func() {
			log.Printf("DataChannel %s is closed", dc.Label())
		})
	})
	return rtcConn, nil
}
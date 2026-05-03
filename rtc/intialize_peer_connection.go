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
	return rtcConn, nil
}
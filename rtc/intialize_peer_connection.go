package rtc

import (
	"encoding/json"
	"log"

	
	"github.com/pion/webrtc/v4"
)

var iceServer webrtc.ICEServer = webrtc.ICEServer{URLs: []string{"stun:stun.l.google.com:19302"}}

func IntializePeerConnection(signal func([]byte)error) (*webrtc.PeerConnection, error) {
	config := webrtc.Configuration{ICEServers: []webrtc.ICEServer{iceServer}}
	rtcConn, err := webrtc.NewPeerConnection(config)
	if(err != nil){
		return nil, err
	}
	rtcConn.OnICECandidate(func(i *webrtc.ICECandidate) {
		if(i == nil){
			return
		}
		rawCandidate := i.ToJSON()
		payload, errForCandiateParse := json.Marshal(rawCandidate)
		if(errForCandiateParse != nil){
			log.Println(errForCandiateParse)
		}
		if errInMessageWriter := signal(payload); errInMessageWriter != nil{
			log.Println("Signaling failed:",errInMessageWriter)
		}
	})
	return rtcConn, nil
}
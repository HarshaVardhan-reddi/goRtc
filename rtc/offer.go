package rtc

import (
	"log"

	"github.com/pion/webrtc/v4"
)

func CreateOffer(webrtcConn *webrtc.PeerConnection) (*webrtc.SessionDescription, error) {
	_, err := webrtcConn.CreateDataChannel("data", nil)
	if(err != nil){
		log.Println("data channel creation failed:", err)
	}
	sdp, errInOffer := webrtcConn.CreateOffer(nil)
	if(errInOffer != nil){
		return nil, errInOffer
	}
	webrtcConn.SetLocalDescription(sdp)
	return &sdp,nil
}

func SetRemoteSdp(webrtcConn *webrtc.PeerConnection, remoteSdp string, sdpType webrtc.SDPType) error {
	desc := webrtc.SessionDescription{
		Type: sdpType,
		SDP: remoteSdp,
	}
	return webrtcConn.SetRemoteDescription(desc)
}

func CreateAnswer(webrtcConn *webrtc.PeerConnection,remoteSdp string) (*webrtc.SessionDescription, error) {
	if err := SetRemoteSdp(webrtcConn, remoteSdp, webrtc.SDPTypeOffer); err != nil{
		log.Println(err)
	}
	sdp, err := webrtcConn.CreateAnswer(nil)
	if(err != nil){
		return nil, err 
	}
	if err := webrtcConn.SetLocalDescription(sdp); err != nil{
		log.Println(err)
	}
	return  &sdp, nil
}
package main

import (
	"fmt"
	"log"

	// "github.com/pion/rtcp"
	"github.com/pion/webrtc/v4"
)

var ICESERVERS webrtc.ICEServer = webrtc.ICEServer{URLs: []string{"stun:stun.l.google.com:19302"}}

func main(){
	fmt.Println("welcome to webrtc series in golang")
	peerconn,err := intializePeerConnection()
	if(err != nil){
		log.Fatal(err)
	}
	sdp, errInOffer := createOffer(peerconn)
	if(errInOffer != nil ){
		log.Fatal(errInOffer)
	}
	fmt.Println("Here is your sdp\n",sdp)
}

func intializePeerConnection() (*webrtc.PeerConnection, error) {
	config := webrtc.Configuration{ICEServers: []webrtc.ICEServer{ICESERVERS}}
	conn, err := webrtc.NewPeerConnection(config)
	if(err != nil){
		return nil, err
	}
	return conn, nil
}
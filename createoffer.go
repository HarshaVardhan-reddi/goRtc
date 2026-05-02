package main

import "github.com/pion/webrtc/v4"

func createOffer(peerconn *webrtc.PeerConnection) (webrtc.SessionDescription,error) {
	peerconn.CreateDataChannel("dataexchg", nil)
	sdp, err := peerconn.CreateOffer(nil)
	if(err != nil){
		return webrtc.SessionDescription{}, err
	}
	return sdp, nil
}
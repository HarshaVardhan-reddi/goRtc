package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v4"
)

type PeerConfig struct{
	Host string `json:"host"`
	Port string `json:"port"`
	IsSecure bool `json:"is_secure"`
}

//go:embed peerconfig.json
var rawPeerConfig string

func singalPeer(webSockConn *websocket.Conn, sdp string, webrtcPeerConn *webrtc.PeerConnection){
	peer := PeerConfig{}
	if err := json.Unmarshal([]byte(rawPeerConfig),&peer); err != nil{
		log.Fatal(err)
	}
	// sending sdp
	for{
		writtter, errInWriter := webSockConn.NextWriter(1)
		if(errInWriter != nil){
			log.Println(errInWriter.Error())
		}
		writtter.Write([]byte(sdp))
		_,reader,err := webSockConn.NextReader()
		if(err != nil){
			log.Println(err.Error())
		}
		var message []byte
		_,errorForMessge := reader.Read(message)
		fmt.Println(errorForMessge)
	}
}
package main

import (
	"fmt"
	"log"
	"net/http"
	"rtc/rtc"
	"rtc/socket"
)

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	fmt.Println("Signaling server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 1. Upgrade the connection
	conn, err := socket.UpgradeCurrentRequestToWebrequest(w, r)
	if err != nil {
		log.Println("Upgrade failed:", err)
		return
	}

	// 2. Create the Signaler Adapter
	wsSignaler := &socket.WSSignaler{Conn: conn}

	// 3. Register the RTC Processor with this specific signaler
	socket.Register("rtc", &rtc.RTCProcessor{
		Signaler: wsSignaler,
	})

	// 4. Start the message loop
	fmt.Println("Client connected, starting packet loop...")
	socket.ProcessPackets(conn)
}

package main

import (
	"fmt"
	"rtc/rtc"       // Adjust based on your module name
	"rtc/socket"    // Adjust based on your module name
)

func main() {
	fmt.Println("welcome to webrtc series in golang")

	// Register the RTC processor
	socket.Register("rtc", &rtc.RTCProcessor{})

	// ... rest of your server setup
}
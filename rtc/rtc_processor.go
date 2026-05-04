package rtc

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/pion/webrtc/v4"
)

// RTCEvent represents the structure of the incoming RTC signaling message
type RTCEvent struct {
	Event string          `json:"event"`      // e.g., "peer_request" (Offer), "answer"
	Data  json.RawMessage `json:"event_data"` // The SDP string or Candidate object
}

var activeConn *webrtc.PeerConnection

type RTCProcessor struct {
	Signaler Signaler
}

func (p *RTCProcessor) Process(payload json.RawMessage) error {
	// 1. Initialize the global connection if it doesn't exist yet
	if activeConn == nil {
		var err error
		fmt.Println("Initializing new global PeerConnection...")
		activeConn, err = InitializePeerConnection(p.Signaler)
		if err != nil {
			return fmt.Errorf("failed to initialize global connection: %w", err)
		}
	}

	var event RTCEvent
	if err := json.Unmarshal(payload, &event); err != nil {
		return fmt.Errorf("failed to unmarshal RTC event: %w", err)
	}

	switch event.Event {
	case "peer_request":
		// The peer sent an Offer, we need to create an Answer
		var sdp string
		if err := json.Unmarshal(event.Data, &sdp); err != nil {
			return err
		}

		fmt.Println("Received Offer, creating Answer...")
		answer, err := CreateAnswer(activeConn, sdp)
		if err != nil {
			return err
		}

		// Wrap the answer in an RTCEvent to send it back
		answerPayload, _ := json.Marshal(answer.SDP)
		respEvent := RTCEvent{
			Event: "answer",
			Data:  answerPayload,
		}
		respBytes, _ := json.Marshal(respEvent)

		// Send it back via the signaler!
		if err := p.Signaler.Send(respBytes); err != nil {
			log.Println("Failed to send answer back:", err)
		}

		log.Println("Sent Answer successfully")

	case "candidate":
		// Peer sent an ICE candidate, we need to add it to our connection
		var candidate webrtc.ICECandidateInit
		if err := json.Unmarshal(event.Data, &candidate); err != nil {
			return fmt.Errorf("failed to unmarshal candidate: %w", err)
		}

		if activeConn != nil {
			fmt.Println("Adding remote ICE Candidate...")
			if err := activeConn.AddICECandidate(candidate); err != nil {
				return fmt.Errorf("failed to add ice candidate: %w", err)
			}
		} else {
			log.Println("Warning: Received candidate but activeConn is nil")
		}

	case "answer":
		// We sent an Offer, and received an Answer back
		var sdp string
		if err := json.Unmarshal(event.Data, &sdp); err != nil {
			return err
		}

		fmt.Println("Received Answer, setting remote description...")
		err := SetRemoteSdp(activeConn, sdp, webrtc.SDPTypeAnswer)
		if err != nil {
			return err
		}
		log.Println("Set Remote Answer successfully")

	default:
		log.Printf("Unknown RTC event: %s", event.Event)
	}

	return nil
}

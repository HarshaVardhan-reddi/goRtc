package rtc

// Signaler defines the abstraction for sending signaling messages.
// This allows the rtc package to remain independent of the transport layer (WebSocket, gRPC, etc.)
type Signaler interface {
	Send(payload []byte) error
}

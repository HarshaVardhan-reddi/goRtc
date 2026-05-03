package socket

import (
	"encoding/json"
	"fmt"
)

var registry map[string]PacketProcessor = make(map[string]PacketProcessor)

func Register(packetType string, processor PacketProcessor) {
	registry[packetType] = processor
}

type Packet struct {
	Id         string          `json:"_id"`
	Type       string          `json:"type"`
	RawMessage json.RawMessage `json:"message"`
}

func (pkt *Packet) validatePacket() error {
	if pkt.Type == "" {
		return fmt.Errorf("type cannot be empty")
	}

	if _, exists := registry[pkt.Type]; !exists {
		return fmt.Errorf("unsupported type %s with Id %s", pkt.Type, pkt.Id)
	}

	return nil
}

func (pkt Packet) DispatchPacket() error {
	processor := registry[pkt.Type]
	return processor.Process(pkt.RawMessage)
}
package socket

import "encoding/json"

type PacketProcessor interface {
	Process(payload json.RawMessage) error
}

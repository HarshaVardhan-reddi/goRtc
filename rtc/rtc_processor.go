package rtc

import (
	"encoding/json"
	"fmt"
)

type RTCProcessor struct {
}

func (p *RTCProcessor) Process(payload json.RawMessage) error {
	fmt.Println("RTC Processor handling payload:", string(payload))
	// Future logic: unmarshal payload into SDP or Candidate
	return nil
}

package response

import (
	"encoding/json"
	"time"

	"github.com/volcengine/VolcEngineRTC_Solution_Demo/internal/pkg/logs"
)

type inform struct {
	MessageType string      `json:"message_type"`
	Event       string      `json:"event"`
	Timestamp   int64       `json:"timestamp"`
	Data        interface{} `json:"data"`
}

func NewInformToClient(event string, data interface{}) string {
	info := inform{
		MessageType: "inform",
		Event:       event,
		Timestamp:   time.Now().UnixNano(),
		Data:        data,
	}

	infoByte, err := json.Marshal(info)
	if err != nil {
		logs.Warnf("json marshal error, input: %v, err: %v", info, err)
	}

	return string(infoByte)
}

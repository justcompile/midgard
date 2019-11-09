package events

type WorkerEventType string

var (
	WorkerEventTypeConnected    WorkerEventType = "CONNECTED"
	WorkerEventTypeDisconnected WorkerEventType = "DISCONNECTED"
)

type WorkerEvent struct {
	Data map[string]interface{} `json:"worker"`
	Type WorkerEventType        `json:"type"`
}

func WorkerConnected(data map[string]interface{}) error {
	return send("worker_connected", &WorkerEvent{
		Data: data,
		Type: WorkerEventTypeConnected,
	})
}

func WorkerDisconnected(data map[string]interface{}) error {
	return send("worker_disconnected", &WorkerEvent{
		Data: data,
		Type: WorkerEventTypeDisconnected,
	})
}

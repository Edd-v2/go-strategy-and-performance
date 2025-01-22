package sender

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type ConnectionManager struct {
	NatsClient *nats.Conn
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{}
}

func (cm *ConnectionManager) Connect(natsURL string) error {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return fmt.Errorf("failed to connect to NATS: %v", err)
	}
	cm.NatsClient = nc
	return nil
}

func (cm *ConnectionManager) Publish(topic, message string) {
	err := cm.NatsClient.Publish(topic, []byte(message))
	if err != nil {
		fmt.Printf("Error publishing to topic '%s': %v\n", topic, err)
	}
}

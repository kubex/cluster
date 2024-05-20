package cluster

import "github.com/nats-io/nats.go"

func Subscribe(topic string, receiver nats.MsgHandler) (*nats.Subscription, error) {
	nc, _ := client()
	return nc.Subscribe(topic, receiver)
}

package cluster

func Publish(topic string, message []byte) error {
	nc, _ := client()
	return nc.Publish(topic, message)
}

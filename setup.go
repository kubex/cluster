package cluster

import (
	"fmt"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"time"
)

const Topic = "my-topic"

var NServer *server.Server

func setup(instance *Instance) (err error) {
	//clusterStr := "nats://host.docker.internal:9622,nats://10.10.10.120:9622"
	opts := &server.Options{
		HTTPPort: instance.httpMonitorPort,
		Port:     instance.natsPort,
		Cluster: server.ClusterOpts{
			Name: instance.clusterName,
			Port: instance.clusterPort,
		},
		Routes: server.RoutesFromStr(instance.natsSeed),
	}

	// Initialize new server with options
	NServer, err = server.NewServer(opts)
	NServer.ConfigureLogger()

	if err != nil {
		return err
	}
	return nil
}

func run(background bool) (err error) {

	if NServer == nil {
		return fmt.Errorf("server not initialized")
	}

	NServer.Start()

	// Wait for server to be ready for connections
	if !NServer.ReadyForConnections(4 * time.Second) {
		return fmt.Errorf("not ready for connection")
	}

	if !background {
		NServer.WaitForShutdown()
	}

	return nil
}

func client() (*nats.Conn, error) {
	return nats.Connect(NServer.ClientURL())
}

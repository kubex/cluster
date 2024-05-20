package cluster

import "github.com/nats-io/nats-server/v2/server"

type Instance struct {
	natsPort        int
	httpMonitorPort int
	natsSeed        string
	clusterPort     int
	clusterName     string
}

type Option func(*Instance) error

func Create(options ...Option) (*Instance, error) {
	inst := &Instance{
		natsPort:        server.RANDOM_PORT,
		httpMonitorPort: server.DEFAULT_HTTP_PORT,
		clusterPort:     4223,
		clusterName:     "NATS",
	}

	for _, option := range options {
		optErr := option(inst)
		if optErr != nil {
			return nil, optErr
		}
	}

	setupErr := setup(inst)
	if setupErr != nil {
		return nil, setupErr
	}

	return inst, nil
}

func Run(background bool) error {
	return run(background)
}

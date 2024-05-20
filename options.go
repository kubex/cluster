package cluster

func WithNatsPort(port int) Option {
	return func(inst *Instance) error {
		inst.natsPort = port
		return nil
	}
}

func WithHTTPMonitorPort(port int) Option {
	return func(inst *Instance) error {
		inst.httpMonitorPort = port
		return nil
	}
}

func WithNatsSeed(seed string) Option {
	return func(inst *Instance) error {
		inst.natsSeed = seed
		return nil
	}
}

func WithClusterPort(port int) Option {
	return func(inst *Instance) error {
		inst.clusterPort = port
		return nil
	}
}

func WithClusterName(name string) Option {
	return func(inst *Instance) error {
		inst.clusterName = name
		return nil
	}
}

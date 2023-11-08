package consullib

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/cast"
)

func RegisterService(consulAddr, serviceName, ip, port string, tag []string) error {
	config := api.DefaultConfig()
	config.Address = consulAddr

	consul, err := api.NewClient(config)
	if err != nil {
		return fmt.Errorf("create consul client err: %w", err)
	}

	registration := &api.AgentServiceRegistration{
		Name:    serviceName,
		ID:      serviceName + "-service-" + port,
		Port:    cast.ToInt(port),
		Address: ip,
		Tags:    tag,
	}

	err = consul.Agent().ServiceRegister(registration)
	if err != nil {
		return fmt.Errorf("register service err:%w", err)
	}

	return nil
}

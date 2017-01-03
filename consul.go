package consul

import (
	"strconv"
	"strings"

	"github.com/hashicorp/consul/api"
)

var Client *ConsulClient

type ConsulClient struct {
	*api.Client
	Tags []string
}

func RegisterPort(serviceName string, portNumber int) error {
	return Client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		Name: serviceName,
		Port: portNumber,
		Tags: Client.Tags,
	})
}

func RegisterListenOn(serviceName string, listenOn string) error {
	listenParts := strings.Split(listenOn, ":")
	listenPort, err := strconv.Atoi(listenParts[len(listenParts)-1])
	if err != nil {
		return err
	}
	return RegisterPort(serviceName, listenPort)
}

func Must() {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}
	Client = &ConsulClient{
		Client: client,
	}
}

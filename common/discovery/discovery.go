package discovery

import "context"

type Registry interface {
	Register(ctx context.Context, instanceID, serviceName, hostPort string) error
	Deregister(ctx context.Context, instanceID, serverName string) error
	Discover(ctx context.Context, serviceName string) ([]string, error)
	HealthCheck(instanceID, serviceName string) error
}

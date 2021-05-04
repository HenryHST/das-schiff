package mock

import (
	"net"

	"gitlab.devops.telekom.de/schiff/engine/schiff-operator.git/pkg/ipam"
)

type Manager struct {
	Callback func(t, identifier, networkView string, subnet *net.IPNet)
}

var _ ipam.Manager = &Manager{}

// func (m *Manager) GetClusterSubnet(clusterIdentifier string, spec ipam.NetworkSpec) (ipam.Subnet, error) {
// 	return Subnet{cid: clusterIdentifier}, nil
// }

func (m *Manager) GetOrAllocateIP(identifier, networkView string, subnet *net.IPNet) (net.IP, error) {
	if m.Callback != nil {
		m.Callback("GetOrAllocate", identifier, networkView, subnet)
	}
	return net.IPv4(10, 0, 0, 0), nil
}

func (m *Manager) ReleaseIP(identifier, networkView string, subnet *net.IPNet) error {
	if m.Callback != nil {
		m.Callback("ReleaseIP", identifier, networkView, subnet)
	}
	return nil
}

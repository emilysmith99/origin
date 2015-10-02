package api

type EventType string

const (
	Added    EventType = "ADDED"
	Deleted  EventType = "DELETED"
	Modified EventType = "MODIFIED"
)

type Subnet struct {
	NodeIP     string
	SubnetCIDR string
}

type SubnetEvent struct {
	Type     EventType
	NodeName string
	Subnet   Subnet
}

type Node struct {
	Name string
	IP   string
}

type NodeEvent struct {
	Type EventType
	Node Node
}

type NetNamespace struct {
	Name  string
	NetID uint
}

type NetNamespaceEvent struct {
	Type  EventType
	Name  string
	NetID uint
}

type NamespaceEvent struct {
	Type EventType
	Name string
}

type ServiceProtocol string

const (
	TCP ServiceProtocol = "TCP"
	UDP ServiceProtocol = "UDP"
)

type ServicePort struct {
	Protocol ServiceProtocol
	Port     uint
}

type Service struct {
	Name      string
	Namespace string
	UID       string
	IP        string
	Ports     []ServicePort
}

type ServiceEvent struct {
	Type    EventType
	Service Service
}

type Pod struct {
	Name        string
	Namespace   string
	ContainerID string
}

package network

import (
	"simple-docker/container"
	"testing"
)

func TestBridgeInit(t *testing.T) {
	d := BridgeNetworkDriver{}
	err := d.Delete(Network{Name: "test-bridge"})
	if err != nil {
		return
	}
	_, err = d.Create("192.168.0.1/24", "test-bridge")
	if err != nil {
		t.Logf("err: %v", err)
	}
}

func TestBridgeConnect(t *testing.T) {
	ep := Endpoint{
		ID: "test container",
	}

	n := Network{
		Name: "test-bridge",
	}

	d := BridgeNetworkDriver{}
	err := d.Connect(&n, &ep)
	t.Logf("err: %v", err)
}

func TestNetworkConnect(t *testing.T) {

	cInfo := &container.Info{
		Id:  "test-container",
		Pid: "15438",
	}

	d := BridgeNetworkDriver{}
	err := d.Delete(Network{Name: "test-bridge"})
	if err != nil {
		return
	}
	n, err := d.Create("192.168.0.1/24", "test-bridge")
	if err != nil {
		t.Logf("err: %v", err)
	}

	err = Init()
	if err != nil {
		return
	}

	networks[n.Name] = n
	err = Connect(n.Name, cInfo)
	t.Logf("err: %v", err)
}

func TestLoad(t *testing.T) {
	n := Network{
		Name: "test-bridge",
	}
	err := n.load("/var/run/simple-docker/network/network/testbridge")
	if err != nil {
		return
	}

	t.Logf("network: %v", n)
}

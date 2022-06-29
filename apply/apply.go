package apply

import (
	"go-learn/runtime"
	"go-learn/runtime/k0s"
	"go-learn/runtime/kubernets"
)

type Processor struct {
	Runtime runtime.Interface
}

func (c *Processor) ApplyClusterFile() error {
	c.Runtime = runtime.NewKubeRuntime()
	metadata, err := c.Runtime.GetMetadata()
	if err != nil {
		return err
	}
	switch metadata {
	case "k8s":
		c.Runtime = kubernets.NewK8sRuntime()
	case "k0s":
		c.Runtime = k0s.NewK0sctlRuntime()
	default:
		c.Runtime = kubernets.NewK8sRuntime()
	}
	return c.Runtime.Init()
}

func NewProcessor() *Processor {
	return &Processor{}
}

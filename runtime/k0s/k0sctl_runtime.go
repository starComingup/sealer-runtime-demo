package k0s

import (
	"fmt"
	"go-learn/runtime"
)

type K0sctlRuntime struct {
	Cluster string
	Action  string
	Config  string
}

func (k *K0sctlRuntime) Init() error {
	fmt.Println("k0s init start...")
	return nil
}

func (k *K0sctlRuntime) Reset() error {
	return nil
}

func (k *K0sctlRuntime) Upgrade() error {
	return nil
}

func (k *K0sctlRuntime) GetMetadata() (string, error) {
	return "", nil
}

func (k *K0sctlRuntime) UpdateCert(cert []string) error {
	return nil
}

func NewK0sctlRuntime() runtime.Interface {
	return &K0sctlRuntime{}
}

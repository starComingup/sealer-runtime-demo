package kubernets

import (
	"fmt"
	"go-learn/runtime"
	"sync"
)

type KubeadmRuntime struct {
	*sync.Mutex
	Cluster string
	Action  string
	Config  string
}

func (k *KubeadmRuntime) Reset() error {
	fmt.Println("will reset the k8s...")
	return nil
}

func (k *KubeadmRuntime) Init() error {
	return k.init()
}

func (k *KubeadmRuntime) Upgrade() error {
	return nil
}

func (k *KubeadmRuntime) GetMetadata() (string, error) {
	return "", nil
}

func (k *KubeadmRuntime) UpdateCert(cert []string) error {
	return nil
}

func NewK8sRuntime() runtime.Interface {
	return &KubeadmRuntime{}
}

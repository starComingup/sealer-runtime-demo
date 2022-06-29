package runtime

type Interface interface {
	Init() error
	Upgrade() error
	Reset() error
	GetMetadata() (string, error)
	UpdateCert(certs []string) error
}

type Metadata struct {
	Version string `json:"version"`
	Arch    string `json:"arch"`
	Variant string `json:"variant"`
	//KubeVersion is a SemVer constraint specifying the version of Kubernetes required.
	KubeVersion string `json:"kubeVersion"`
	NydusFlag   bool   `json:"NydusFlag"`
}

type KubeRuntime struct {
}

func (k *KubeRuntime) Init() error {
	return nil
}

func (k *KubeRuntime) Upgrade() error {
	return nil
}

func (k *KubeRuntime) Reset() error {
	return nil
}

func (k *KubeRuntime) GetMetadata() (string, error) {
	return "k8s", nil
}

func (k *KubeRuntime) UpdateCert(certs []string) error {
	return nil
}

func NewKubeRuntime() *KubeRuntime {
	return &KubeRuntime{}
}

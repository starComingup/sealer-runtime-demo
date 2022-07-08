package kubernets

type Metadata struct {
	Version string `json:"version"`
	Arch    string `json:"arch"`
	Variant string `json:"variant"`
	//ClusterRuntime is a Flag to distinguish a cluster install tool.
	ClusterRuntime ClusterRuntime `json:"ClusterRuntime"`
	//KubeVersion is a SemVer constraint specifying the version of Kubernetes required.
	KubeVersion string `json:"kubeVersion"`
	NydusFlag   bool   `json:"NydusFlag"`
}

type ClusterRuntime string

const (
	K0s ClusterRuntime = "k0s"
	K3s ClusterRuntime = "k3s"
	K8s ClusterRuntime = "k8s"
)

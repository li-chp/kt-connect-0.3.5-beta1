package options

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/alibaba/kt-connect/pkg/kt/util"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
)

// ConnectOptions ...
type ConnectOptions struct {
	Global           bool
	DisablePodIp     bool
	DisableTunDevice bool
	DisableTunRoute  bool
	SocksPort        int
	DnsCacheTtl      int64
	IncludeIps       string
	ExcludeIps       string
	ConnectMode      string
	DnsMode          string
	SharedShadow     bool
	ClusterDomain    string
	SkipCleanup      bool
}

// ExchangeOptions ...
type ExchangeOptions struct {
	Mode            string
	Expose          string
	RecoverWaitTime int
}

// MeshOptions ...
type MeshOptions struct {
	MeshMode            string
	Expose              string
	VersionMark         string
	RouterImage         string
	VirtualServiceName  string
	DestinationRuleName string
}

// MeshDebugOptions ...
type MeshDebugOptions struct {
	Global              bool
	DisablePodIp        bool
	DisableTunDevice    bool
	DisableTunRoute     bool
	SocksPort           int
	DnsCacheTtl         int64
	IncludeIps          string
	ExcludeIps          string
	ConnectMode         string
	DnsMode             string
	SharedShadow        bool
	ClusterDomain       string
	SkipCleanup         bool
	MeshMode            string
	Expose              string
	VersionMark         string
	RouterImage         string
	VirtualServiceName  string
	DestinationRuleName string
}

// PreviewOptions ...
type PreviewOptions struct {
	External bool
	Expose   string
}

// CleanOptions ...
type CleanOptions struct {
	DryRun           bool
	ThresholdInMinus int64
	SweepLocalRoute  bool
}

// RuntimeOptions ...
type RuntimeOptions struct {
	Clientset   kubernetes.Interface
	IstioClient versionedclient.Interface
	// Version ktctl version
	Version string
	// UserHome path of user home, same as ${HOME}
	UserHome string
	// AppHome path of kt config folder, default to ${UserHome}/.ktctl
	AppHome string
	// Component current sub-command (connect, exchange, mesh or preview)
	Component string
	// Shadow pod name
	Shadow string
	// Router pod name
	Router string
	// Mesh version of mesh pod
	Mesh string
	// Mesh version of mesh pod
	MeshDebug string
	// Origin the origin deployment or service name
	Origin string
	// Replicas the origin replicas
	Replicas int32
	// Service exposed service name
	Service string
	// RestConfig kubectl config
	RestConfig           *rest.Config
	VirtualServicePatch  bool
	DestinationRulePatch bool
}

// DaemonOptions cli options
type DaemonOptions struct {
	RuntimeStore        *RuntimeOptions
	PreviewOptions      *PreviewOptions
	ConnectOptions      *ConnectOptions
	ExchangeOptions     *ExchangeOptions
	MeshOptions         *MeshOptions
	MeshDebugOptions    *MeshDebugOptions
	CleanOptions        *CleanOptions
	RunAsWorkerProcess  bool
	KubeConfig          string
	Namespace           string
	ServiceAccount      string
	Debug               bool
	Image               string
	ImagePullSecret     string
	NodeSelector        string
	WithLabels          string
	WithAnnotations     string
	PortForwardWaitTime int
	PodCreationWaitTime int
	UseShadowDeployment bool
	AlwaysUpdateShadow  bool
	SkipTimeDiff        bool
	KubeContext         string
	PodQuota            string
}

var opt *DaemonOptions

// Get fetch options instance
func Get() *DaemonOptions {
	if opt == nil {
		opt = &DaemonOptions{
			Namespace: util.DefaultNamespace,
			RuntimeStore: &RuntimeOptions{
				UserHome: util.UserHome,
				AppHome:  util.KtHome,
			},
			ConnectOptions:   &ConnectOptions{},
			ExchangeOptions:  &ExchangeOptions{},
			MeshOptions:      &MeshOptions{},
			MeshDebugOptions: &MeshDebugOptions{},
			PreviewOptions:   &PreviewOptions{},
			CleanOptions:     &CleanOptions{},
		}
	}
	return opt
}

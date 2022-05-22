package util

const (
	// EnvKubeConfig environment variable for kube config file
	EnvKubeConfig = "KUBECONFIG"

	// KubernetesToolkit name of this tool
	KubernetesToolkit = "et"
	// ComponentConnect connect command
	ComponentConnect = "connect"
	// ComponentExchange exchange command
	ComponentExchange = "exchange"
	// ComponentMesh mesh command
	ComponentMesh = "mesh"
	// ComponentMesh meshDebug command
	ComponentMeshDebug = "meshDebug"
	// ComponentPreview preview command
	ComponentPreview = "preview"

	// ImageKtShadow default shadow image
	ImageKtShadow = "10.10.125.67/et/et-connect-shadow"
	// ImageKtRouter default router image
	ImageKtRouter = "10.10.125.67/et/et-connect-router"
	// ImageKtNavigator default navigator image
	ImageKtNavigator = "10.10.125.67/et/et-connect-navigator"

	// ConnectModeShuttle sshuttle mode
	ConnectModeShuttle = "sshuttle"
	// ConnectModeTun2Socks tun2socks mode
	ConnectModeTun2Socks = "tun2socks"
	// ExchangeModeScale scale mode
	ExchangeModeScale = "scale"
	// ExchangeModeEphemeral ephemeral mode
	ExchangeModeEphemeral = "ephemeral"
	// ExchangeModeSelector selector mode
	ExchangeModeSelector = "selector"
	// MeshModeAuto auto mode
	MeshModeAuto = "auto"
	// MeshModeManual manual mode
	MeshModeManual = "manual"
	// DnsModeLocalDns local dns mode
	DnsModeLocalDns = "localDNS"
	// DnsModePodDns pod dns mode
	DnsModePodDns = "podDNS"
	// DnsModeHosts hosts dns mode
	DnsModeHosts = "hosts"
	// DnsOrderCluster proxy to cluster dns
	DnsOrderCluster = "cluster"
	// DnsOrderUpstream proxy to upstream dns
	DnsOrderUpstream = "upstream"

	// ControlBy label used for mark shadow pod
	ControlBy = "control-by"
	// KtTarget label used for service selecting shadow or route pod
	KtTarget = "et-target"
	// KtRole label used for auto mesh roles
	KtRole = "et-role"
	// KtConfig annotation used for clean up context
	KtConfig = "et-config"
	// KtUser annotation used for record independent username
	KtUser = "et-user"
	// KtSelector annotation used for record service origin selector
	KtSelector = "et-selector"
	// KtRefCount annotation used for count of shared pod / service
	KtRefCount = "et-ref-count"
	// KtLastHeartBeat annotation used for timestamp of last heart beat
	KtLastHeartBeat = "et-last-heart-beat"
	// KtLock annotation used for avoid auto mesh conflict
	KtLock = "et-lock"

	// PostfixRsaKey postfix of local private key name
	PostfixRsaKey = "_id_rsa"
	// RouterBin path to router executable
	RouterBin = "/usr/sbin/router"
	// DumpHostsLock path to lock file for dump hosts
	DumpHostsLock = "hosts.lock"
	// SshBitSize ssh bit size
	SshBitSize = 2048
	// SshAuthKey auth key name
	SshAuthKey = "authorized"
	// SshAuthPrivateKey ssh private key
	SshAuthPrivateKey = "privateKey"
	// DefaultNamespace default namespace
	DefaultNamespace = "default"
	// KtExchangeContainer name of exchange ephemeral container
	KtExchangeContainer = "et-exchange"
	// DefaultContainer default container name
	DefaultContainer = "standalone"
	// StuntmanServiceSuffix suffix of stuntman service name
	StuntmanServiceSuffix = "-et-stuntman"
	// RouterPodSuffix suffix of router pod name
	RouterPodSuffix = "-et-router"
	// ExchangePodInfix exchange pod name
	ExchangePodInfix = "-et-exchange-"
	// MeshPodInfix mesh pod and mesh service name
	MeshPodInfix = "-et-mesh-"
	// RectifierPodPrefix rectifier pod name
	RectifierPodPrefix = "et-rectifier-"
	// RoleConnectShadow shadow role
	RoleConnectShadow = "shadow-connect"
	// RoleExchangeShadow shadow role
	RoleExchangeShadow = "shadow-exchange"
	// RoleMeshShadow shadow role
	RoleMeshShadow = "shadow-mesh"
	// RolePreviewShadow shadow role
	RolePreviewShadow = "shadow-preview"
	// RoleRouter router role
	RoleRouter = "router"
	// TunNameWin tun device name in windows
	TunNameWin = "EtConnectTunnel"
	// TunNameLinux tun device name in linux
	TunNameLinux = "et0"
	// TunNameMac tun device name in MacOS
	TunNameMac = "utun"
	// AlternativeDnsPort alternative port for local dns
	AlternativeDnsPort = 10053
)

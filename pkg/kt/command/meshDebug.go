package command

import (
	"fmt"
	"github.com/alibaba/kt-connect/pkg/kt/command/connect"
	"github.com/alibaba/kt-connect/pkg/kt/command/general"
	"github.com/alibaba/kt-connect/pkg/kt/command/mesh"
	opt "github.com/alibaba/kt-connect/pkg/kt/command/options"
	"github.com/alibaba/kt-connect/pkg/kt/util"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// NewMeshCommand return new mesh command
func NewMeshDebugCommand(action ActionInterface) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "meshDebug",
		Short: "combined connect and mesh in one command",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return general.Prepare()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("name of service to meshDebug is required")
			}
			return action.MeshDebug(args[0])
		},
	}

	cmd.SetUsageTemplate(fmt.Sprintf(general.UsageTemplate, "et meshDebug <service-name> [command options]"))
	cmd.Long = cmd.Short

	cmd.Flags().SortFlags = false
	cmd.InheritedFlags().SortFlags = false
	cmd.Flags().StringVar(&opt.Get().ConnectOptions.ConnectMode, "connectMode", util.ConnectModeTun2Socks, "Connect mode 'tun2socks' or 'sshuttle'")
	cmd.Flags().StringVar(&opt.Get().ConnectOptions.DnsMode, "dnsMode", util.DnsModeLocalDns, "Specify how to resolve service domains, can be 'localDNS', 'podDNS', 'hosts' or 'hosts:<namespaces>', for multiple namespaces use ',' separation")
	cmd.Flags().BoolVar(&opt.Get().ConnectOptions.SharedShadow, "shareShadow", false, "Use shared shadow pod")
	cmd.Flags().StringVar(&opt.Get().ConnectOptions.ClusterDomain, "clusterDomain", "cluster.local", "The cluster domain provided to kubernetes api-server")
	cmd.Flags().BoolVar(&opt.Get().ConnectOptions.DisablePodIp, "disablePodIp", false, "Disable access to pod IP address")
	cmd.Flags().BoolVar(&opt.Get().ConnectOptions.SkipCleanup, "skipCleanup", false, "Do not auto cleanup residual resources in cluster")
	cmd.Flags().StringVar(&opt.Get().ConnectOptions.IncludeIps, "includeIps", "", "Specify extra IP ranges which should be route to cluster, e.g. '172.2.0.0/16', use ',' separated")
	cmd.Flags().StringVar(&opt.Get().ConnectOptions.ExcludeIps, "excludeIps", "", "Do not route specified IPs to cluster, e.g. '192.168.64.2' or '192.168.64.0/24', use ',' separated")
	cmd.Flags().BoolVar(&opt.Get().ConnectOptions.DisableTunDevice, "disableTunDevice", false, "(tun2socks mode only) Create socks5 proxy without tun device")
	cmd.Flags().BoolVar(&opt.Get().ConnectOptions.DisableTunRoute, "disableTunRoute", false, "(tun2socks mode only) Do not auto setup tun device route")
	cmd.Flags().IntVar(&opt.Get().ConnectOptions.SocksPort, "proxyPort", 2223, "(tun2socks mode only) Specify the local port which socks5 proxy should use")
	cmd.Flags().Int64Var(&opt.Get().ConnectOptions.DnsCacheTtl, "dnsCacheTtl", 60, "(local dns mode only) DNS cache refresh interval in seconds")

	cmd.Flags().StringVar(&opt.Get().MeshOptions.Expose, "expose", "", "Ports to expose, use ',' separated, in [port] or [local:remote] format, e.g. 7001,8080:80")
	cmd.Flags().StringVar(&opt.Get().MeshOptions.MeshMode, "meshMode", util.MeshModeAuto, "Mesh method 'auto' or 'manual'")
	cmd.Flags().StringVar(&opt.Get().MeshOptions.VersionMark, "versionMark", "", "Specify the version of mesh service, e.g. '0.0.1' or 'mark:local'")
	cmd.Flags().StringVar(&opt.Get().MeshOptions.RouterImage, "routerImage", fmt.Sprintf("%s:v%s", util.ImageKtRouter, opt.Get().RuntimeStore.Version), "(auto method only) Customize router image")
	cmd.Flags().StringVar(&opt.Get().MeshOptions.VirtualServiceName, "vsName", "", "(manual method only) Specify istio VirtualService name")
	cmd.Flags().StringVar(&opt.Get().MeshOptions.DestinationRuleName, "drName", "", "(manual method only) Specify istio DestinationRule name")
	_ = cmd.MarkFlagRequired("expose")
	return cmd
}

//Mesh exchange kubernetes workload
func (action *Action) MeshDebug(resourceName string) error {
	ch, err := general.SetupProcess(util.ComponentMeshDebug)
	if err != nil {
		return err
	}

	if port := util.FindBrokenLocalPort(opt.Get().MeshOptions.Expose); port != "" {
		return fmt.Errorf("no application is running on port %s", port)
	}

	// Get service to mesh
	svc, err := general.GetServiceByResourceName(resourceName, opt.Get().Namespace)
	if err != nil {
		return err
	}

	if port := util.FindInvalidRemotePort(opt.Get().MeshOptions.Expose, general.GetTargetPorts(svc)); port != "" {
		return fmt.Errorf("target port %s not exists in service %s", port, svc.Name)
	}

	if !opt.Get().ConnectOptions.SkipCleanup {
		go silenceCleanup()
	}

	if opt.Get().ConnectOptions.ConnectMode == util.ConnectModeTun2Socks {
		err = connect.ByTun2Socks()
	} else if opt.Get().ConnectOptions.ConnectMode == util.ConnectModeShuttle {
		err = connect.BySshuttle()
	} else {
		err = fmt.Errorf("invalid connect mode: '%s', supportted mode are %s, %s", opt.Get().ConnectOptions.ConnectMode,
			util.ConnectModeTun2Socks, util.ConnectModeShuttle)
	}
	if err != nil {
		return err
	}
	log.Info().Msg("---------------------------------------------------------------")
	log.Info().Msgf(" All looks good, now you can access to resources in the kubernetes cluster")
	log.Info().Msg("---------------------------------------------------------------")

	if opt.Get().MeshOptions.MeshMode == util.MeshModeManual {
		err = mesh.ManualMesh(svc)
	} else if opt.Get().MeshOptions.MeshMode == util.MeshModeAuto {
		err = mesh.AutoMesh(svc)
	} else {
		err = fmt.Errorf("invalid mesh method '%s', supportted are %s, %s", opt.Get().MeshOptions.MeshMode,
			util.MeshModeAuto, util.MeshModeManual)
	}
	if err != nil {
		return err
	}

	// watch background process, clean the workspace and exit if background process occur exception
	s := <-ch
	log.Info().Msgf("Terminal Signal is %s", s)
	return nil
}

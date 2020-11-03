package api

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kops/pkg/apis/kops"
)

type Cluster struct {
	Name    string
	Channel string
	// Addons
	ConfigBase             string
	CloudProvider          string
	ContainerRuntime       string
	KubernetesVersion      string
	Subnet                 []kops.ClusterSubnetSpec
	Project                string
	MasterPublicName       string
	MasterInternalName     string
	NetworkCIDR            string
	AdditionalNetworkCIDRs []string
	NetworkID              string
	Topology               *kops.TopologySpec
	SecretStore            string
	KeyStore               string
	ConfigStore            string
	DNSZone                string
	// AdditionalSANs
	ClusterDNSDomain      string
	ServiceClusterIPRange string
	PodCIDR               string
	NonMasqueradeCIDR     string
	SSHAccess             []string
	NodePortAccess        []string
	// EgressProxy
	SSHKeyName          *string
	KubernetesAPIAccess []string
	IsolateMasters      *bool
	UpdatePolicy        *string
	ExternalPolicies    *map[string][]string
	AdditionalPolicies  *map[string]string
	// FileAssets
	EtcdCluster []*kops.EtcdClusterSpec
	// Containerd                     *ContainerdConfig
	// Docker                         *DockerConfig
	// KubeDNS                        *KubeDNSConfig
	// KubeAPIServer                  *KubeAPIServerConfig
	// KubeControllerManager          *KubeControllerManagerConfig
	// ExternalCloudControllerManager *CloudControllerManagerConfig
	// KubeScheduler                  *KubeSchedulerConfig
	// KubeProxy                      *KubeProxyConfig
	// Kubelet                        *KubeletConfigSpec
	// MasterKubelet                  *KubeletConfigSpec
	// CloudConfig                    *CloudConfiguration
	// ExternalDNS                    *ExternalDNSConfig
	Networking *kops.NetworkingSpec
	API        *kops.AccessSpec
	// Authentication
	// Authorization
	// NodeAuthorization
	CloudLabels map[string]string
	// Hooks
	// Assets
	// IAM
	EncryptionConfig    *bool
	DisableSubnetTags   bool
	UseHostCertificates *bool
	SysctlParameters    []string
	RollingUpdate       *kops.RollingUpdate
	InstanceGroup       []*InstanceGroup
}

func FromKopsCluster(cluster *kops.Cluster, instanceGroups ...*kops.InstanceGroup) *Cluster {
	return &Cluster{
		Name:                   cluster.ObjectMeta.Name,
		Channel:                cluster.Spec.Channel,
		ConfigBase:             cluster.Spec.ConfigBase,
		CloudProvider:          cluster.Spec.CloudProvider,
		ContainerRuntime:       cluster.Spec.ContainerRuntime,
		KubernetesVersion:      cluster.Spec.KubernetesVersion,
		Subnet:                 cluster.Spec.Subnets,
		Project:                cluster.Spec.Project,
		MasterPublicName:       cluster.Spec.MasterPublicName,
		MasterInternalName:     cluster.Spec.MasterInternalName,
		NetworkCIDR:            cluster.Spec.NetworkCIDR,
		AdditionalNetworkCIDRs: cluster.Spec.AdditionalNetworkCIDRs,
		NetworkID:              cluster.Spec.NetworkID,
		Topology:               cluster.Spec.Topology,
		SecretStore:            cluster.Spec.SecretStore,
		KeyStore:               cluster.Spec.KeyStore,
		ConfigStore:            cluster.Spec.ConfigStore,
		DNSZone:                cluster.Spec.DNSZone,
		// AdditionalSANs
		ClusterDNSDomain:      cluster.Spec.ClusterDNSDomain,
		ServiceClusterIPRange: cluster.Spec.ServiceClusterIPRange,
		PodCIDR:               cluster.Spec.PodCIDR,
		NonMasqueradeCIDR:     cluster.Spec.NonMasqueradeCIDR,
		SSHAccess:             cluster.Spec.SSHAccess,
		NodePortAccess:        cluster.Spec.NodePortAccess,
		// EgressProxy
		SSHKeyName:          cluster.Spec.SSHKeyName,
		KubernetesAPIAccess: cluster.Spec.KubernetesAPIAccess,
		IsolateMasters:      cluster.Spec.IsolateMasters,
		UpdatePolicy:        cluster.Spec.UpdatePolicy,
		ExternalPolicies:    cluster.Spec.ExternalPolicies,
		AdditionalPolicies:  cluster.Spec.AdditionalPolicies,
		// FileAssets
		EtcdCluster: cluster.Spec.EtcdClusters,
		// Containerd                     *ContainerdConfig
		// Docker                         *DockerConfig
		// KubeDNS                        *KubeDNSConfig
		// KubeAPIServer                  *KubeAPIServerConfig
		// KubeControllerManager          *KubeControllerManagerConfig
		// ExternalCloudControllerManager *CloudControllerManagerConfig
		// KubeScheduler                  *KubeSchedulerConfig
		// KubeProxy                      *KubeProxyConfig
		// Kubelet                        *KubeletConfigSpec
		// MasterKubelet                  *KubeletConfigSpec
		// CloudConfig                    *CloudConfiguration
		// ExternalDNS                    *ExternalDNSConfig
		Networking: cluster.Spec.Networking,
		API:        cluster.Spec.API,
		// Authentication
		// Authorization
		// NodeAuthorization
		CloudLabels: cluster.Spec.CloudLabels,
		// Hooks
		// Assets
		// IAM
		EncryptionConfig:    cluster.Spec.EncryptionConfig,
		DisableSubnetTags:   cluster.Spec.DisableSubnetTags,
		UseHostCertificates: cluster.Spec.UseHostCertificates,
		SysctlParameters:    cluster.Spec.SysctlParameters,
		RollingUpdate:       cluster.Spec.RollingUpdate,
		InstanceGroup: func(in ...*kops.InstanceGroup) []*InstanceGroup {
			var out []*InstanceGroup
			for _, in := range in {
				out = append(out, FromKopsInstanceGroup(in))
			}
			return out
		}(instanceGroups...),
	}
}

func ToKopsCluster(cluster *Cluster) (*kops.Cluster, []*kops.InstanceGroup) {
	c := kops.Cluster{
		ObjectMeta: metav1.ObjectMeta{
			Name: cluster.Name,
		},
		Spec: kops.ClusterSpec{
			Channel:                cluster.Channel,
			ConfigBase:             cluster.ConfigBase,
			CloudProvider:          cluster.CloudProvider,
			ContainerRuntime:       cluster.ContainerRuntime,
			KubernetesVersion:      cluster.KubernetesVersion,
			Subnets:                cluster.Subnet,
			Project:                cluster.Project,
			MasterPublicName:       cluster.MasterPublicName,
			MasterInternalName:     cluster.MasterInternalName,
			NetworkCIDR:            cluster.NetworkCIDR,
			AdditionalNetworkCIDRs: cluster.AdditionalNetworkCIDRs,
			NetworkID:              cluster.NetworkID,
			Topology:               cluster.Topology,
			SecretStore:            cluster.SecretStore,
			KeyStore:               cluster.KeyStore,
			ConfigStore:            cluster.ConfigStore,
			DNSZone:                cluster.DNSZone,
			// AdditionalSANs
			ClusterDNSDomain:      cluster.ClusterDNSDomain,
			ServiceClusterIPRange: cluster.ServiceClusterIPRange,
			PodCIDR:               cluster.PodCIDR,
			NonMasqueradeCIDR:     cluster.NonMasqueradeCIDR,
			SSHAccess:             cluster.SSHAccess,
			NodePortAccess:        cluster.NodePortAccess,
			// EgressProxy
			SSHKeyName:          cluster.SSHKeyName,
			KubernetesAPIAccess: cluster.KubernetesAPIAccess,
			IsolateMasters:      cluster.IsolateMasters,
			UpdatePolicy:        cluster.UpdatePolicy,
			ExternalPolicies:    cluster.ExternalPolicies,
			AdditionalPolicies:  cluster.AdditionalPolicies,
			// FileAssets
			EtcdClusters: cluster.EtcdCluster,
			// Containerd                     *ContainerdConfig
			// Docker                         *DockerConfig
			// KubeDNS                        *KubeDNSConfig
			// KubeAPIServer                  *KubeAPIServerConfig
			// KubeControllerManager          *KubeControllerManagerConfig
			// ExternalCloudControllerManager *CloudControllerManagerConfig
			// KubeScheduler                  *KubeSchedulerConfig
			// KubeProxy                      *KubeProxyConfig
			// Kubelet                        *KubeletConfigSpec
			// MasterKubelet                  *KubeletConfigSpec
			// CloudConfig                    *CloudConfiguration
			// ExternalDNS                    *ExternalDNSConfig
			Networking: cluster.Networking,
			API:        cluster.API,
			// Authentication
			// Authorization
			// NodeAuthorization
			CloudLabels: cluster.CloudLabels,
			// Hooks
			// Assets
			// IAM
			EncryptionConfig:    cluster.EncryptionConfig,
			DisableSubnetTags:   cluster.DisableSubnetTags,
			UseHostCertificates: cluster.UseHostCertificates,
			SysctlParameters:    cluster.SysctlParameters,
			RollingUpdate:       cluster.RollingUpdate,
		},
	}
	var ig []*kops.InstanceGroup
	for _, instanceGroup := range cluster.InstanceGroup {
		ig = append(ig, ToKopsInstanceGroup(instanceGroup))
	}
	return &c, ig
}

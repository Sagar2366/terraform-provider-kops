package utils

import (
	"context"
	"time"

	"k8s.io/kops/pkg/client/simple"
	"k8s.io/kops/pkg/kubeconfig"
	"k8s.io/kops/upup/pkg/fi/cloudup"
)

func GetKubeConfigBuilder(clientset simple.Clientset, clusterName string, admin *time.Duration, internal bool) (*kubeconfig.KubeconfigBuilder, error) {
	cluster, err := clientset.GetCluster(context.Background(), clusterName)
	if err != nil {
		return nil, err
	}
	keyStore, err := clientset.KeyStore(cluster)
	if err != nil {
		return nil, err
	}
	secretStore, err := clientset.SecretStore(cluster)
	if err != nil {
		return nil, err
	}
	duration := kubeconfig.DefaultKubecfgAdminLifetime
	if admin != nil {
		duration = *admin
	}
	cloud, err := cloudup.BuildCloud(cluster)
	if err != nil {
		return nil, err
	}
	conf, err := kubeconfig.BuildKubecfg(cluster, keyStore, secretStore, cloud, duration, "", internal, "", false)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

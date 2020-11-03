# kops_kubelet_config_spec

| attribute | type | optional/required | computed |
| --- | --- | --- | --- |
| `api_servers` | String | Optional |  |
| `anonymous_auth` | Bool | Optional |  |
| `authorization_mode` | String | Optional |  |
| `bootstrap_kubeconfig` | String | Optional |  |
| `client_ca_file` | String | Optional |  |
| `tls_cert_file` | String | Optional |  |
| `tls_private_key_file` | String | Optional |  |
| `tls_cipher_suites` | List(String) | Optional |  |
| `tls_min_version` | String | Optional |  |
| `kubeconfig_path` | String | Optional |  |
| `require_kubeconfig` | Bool | Optional |  |
| `log_level` | Int | Optional |  |
| `pod_manifest_path` | String | Optional |  |
| `hostname_override` | String | Optional |  |
| `pod_infra_container_image` | String | Optional |  |
| `seccomp_profile_root` | String | Optional |  |
| `allow_privileged` | Bool | Optional |  |
| `enable_debugging_handlers` | Bool | Optional |  |
| `register_node` | Bool | Optional |  |
| `node_status_update_frequency` | Duration | Optional |  |
| `cluster_domain` | String | Optional |  |
| `cluster_dns` | String | Optional |  |
| `network_plugin_name` | String | Optional |  |
| `cloud_provider` | String | Optional |  |
| `kubelet_cgroups` | String | Optional |  |
| `runtime_cgroups` | String | Optional |  |
| `read_only_port` | Int | Optional |  |
| `system_cgroups` | String | Optional |  |
| `cgroup_root` | String | Optional |  |
| `configure_cbr_00` | Bool | Optional |  |
| `hairpin_mode` | String | Optional |  |
| `babysit_daemons` | Bool | Optional |  |
| `max_pods` | Int | Optional |  |
| `nvidia_gp_uss` | Int | Optional |  |
| `pod_cidr` | String | Optional |  |
| `resolver_config` | String | Optional |  |
| `reconcile_cidr` | Bool | Optional |  |
| `register_schedulable` | Bool | Optional |  |
| `serialize_image_pulls` | Bool | Optional |  |
| `node_labels` | Map(String) | Optional |  |
| `non_masquerade_cidr` | String | Optional |  |
| `enable_custom_metrics` | Bool | Optional |  |
| `network_plugin_mtu` | Int | Optional |  |
| `image_gc_high_threshold_percent` | Int | Optional |  |
| `image_gc_low_threshold_percent` | Int | Optional |  |
| `image_pull_progress_deadline` | Duration | Optional |  |
| `eviction_hard` | String | Optional |  |
| `eviction_soft` | String | Optional |  |
| `eviction_soft_grace_period` | String | Optional |  |
| `eviction_pressure_transition_period` | Duration | Optional |  |
| `eviction_max_pod_grace_period` | Int | Optional |  |
| `eviction_minimum_reclaim` | String | Optional |  |
| `volume_plugin_directory` | String | Optional |  |
| `taints` | List(String) | Optional |  |
| `feature_gates` | Map(String) | Optional |  |
| `kube_reserved` | Map(String) | Optional |  |
| `kube_reserved_cgroup` | String | Optional |  |
| `system_reserved` | Map(String) | Optional |  |
| `system_reserved_cgroup` | String | Optional |  |
| `enforce_node_allocatable` | String | Optional |  |
| `runtime_request_timeout` | Duration | Optional |  |
| `volume_stats_agg_period` | Duration | Optional |  |
| `fail_swap_on` | Bool | Optional |  |
| `experimental_allowed_unsafe_sysctls` | List(String) | Optional |  |
| `allowed_unsafe_sysctls` | List(String) | Optional |  |
| `streaming_connection_idle_timeout` | Duration | Optional |  |
| `docker_disable_shared_pid` | Bool | Optional |  |
| `root_dir` | String | Optional |  |
| `authentication_token_webhook` | Bool | Optional |  |
| `authentication_token_webhook_cache_ttl` | Duration | Optional |  |
| `cpucfs_quota` | Bool | Optional |  |
| `cpucfs_quota_period` | Duration | Optional |  |
| `cpu_manager_policy` | String | Optional |  |
| `registry_pull_qps` | Int | Optional |  |
| `registry_burst` | Int | Optional |  |
| `topology_manager_policy` | String | Optional |  |
| `rotate_certificates` | Bool | Optional |  |
| `protect_kernel_defaults` | Bool | Optional |  |
| `cgroup_driver` | String | Optional |  |
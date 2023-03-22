// Code generated by monitor-code-gen. DO NOT EDIT.

package openstack

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "collectd/openstack"

var groupSet = map[string]bool{}

const (
	counterOpenstackNovaServerCPUTime                 = "counter.openstack.nova.server.cpu_time"
	counterOpenstackNovaServerRx                      = "counter.openstack.nova.server.rx"
	counterOpenstackNovaServerRxPackets               = "counter.openstack.nova.server.rx_packets"
	counterOpenstackNovaServerTx                      = "counter.openstack.nova.server.tx"
	counterOpenstackNovaServerTxPackets               = "counter.openstack.nova.server.tx_packets"
	gaugeOpenstackCinderLimitMaxTotalBackupGigabytes  = "gauge.openstack.cinder.limit.maxTotalBackupGigabytes"
	gaugeOpenstackCinderLimitMaxTotalBackups          = "gauge.openstack.cinder.limit.maxTotalBackups"
	gaugeOpenstackCinderLimitMaxTotalSnapshots        = "gauge.openstack.cinder.limit.maxTotalSnapshots"
	gaugeOpenstackCinderLimitMaxTotalVolumeGigabytes  = "gauge.openstack.cinder.limit.maxTotalVolumeGigabytes"
	gaugeOpenstackCinderLimitMaxTotalVolumes          = "gauge.openstack.cinder.limit.maxTotalVolumes"
	gaugeOpenstackCinderLimitTotalBackupGigabytesUsed = "gauge.openstack.cinder.limit.totalBackupGigabytesUsed"
	gaugeOpenstackCinderLimitTotalBackupsUsed         = "gauge.openstack.cinder.limit.totalBackupsUsed"
	gaugeOpenstackCinderLimitTotalGigabytesUsed       = "gauge.openstack.cinder.limit.totalGigabytesUsed"
	gaugeOpenstackCinderLimitTotalSnapshotsUsed       = "gauge.openstack.cinder.limit.totalSnapshotsUsed"
	gaugeOpenstackCinderLimitTotalVolumesUsed         = "gauge.openstack.cinder.limit.totalVolumesUsed"
	gaugeOpenstackCinderSnapshotCount                 = "gauge.openstack.cinder.snapshot.count"
	gaugeOpenstackCinderSnapshotSize                  = "gauge.openstack.cinder.snapshot.size"
	gaugeOpenstackCinderVolumeCount                   = "gauge.openstack.cinder.volume.count"
	gaugeOpenstackCinderVolumeSize                    = "gauge.openstack.cinder.volume.size"
	gaugeOpenstackNeutronFloatingipCount              = "gauge.openstack.neutron.floatingip.count"
	gaugeOpenstackNeutronNetworkCount                 = "gauge.openstack.neutron.network.count"
	gaugeOpenstackNeutronRouterCount                  = "gauge.openstack.neutron.router.count"
	gaugeOpenstackNeutronSecuritygroupCount           = "gauge.openstack.neutron.securitygroup.count"
	gaugeOpenstackNeutronSubnetCount                  = "gauge.openstack.neutron.subnet.count"
	gaugeOpenstackNovaHypervisorCurrentWorkload       = "gauge.openstack.nova.hypervisor.current_workload"
	gaugeOpenstackNovaHypervisorDiskAvailableLeast    = "gauge.openstack.nova.hypervisor.disk_available_least"
	gaugeOpenstackNovaHypervisorFreeDiskGb            = "gauge.openstack.nova.hypervisor.free_disk_gb"
	gaugeOpenstackNovaHypervisorFreeRAMMb             = "gauge.openstack.nova.hypervisor.free_ram_mb"
	gaugeOpenstackNovaHypervisorLoadAverage           = "gauge.openstack.nova.hypervisor.load_average"
	gaugeOpenstackNovaHypervisorLocalGb               = "gauge.openstack.nova.hypervisor.local_gb"
	gaugeOpenstackNovaHypervisorLocalGbUsed           = "gauge.openstack.nova.hypervisor.local_gb_used"
	gaugeOpenstackNovaHypervisorMemoryMb              = "gauge.openstack.nova.hypervisor.memory_mb"
	gaugeOpenstackNovaHypervisorMemoryMbUsed          = "gauge.openstack.nova.hypervisor.memory_mb_used"
	gaugeOpenstackNovaHypervisorRunningVms            = "gauge.openstack.nova.hypervisor.running_vms"
	gaugeOpenstackNovaHypervisorVcpus                 = "gauge.openstack.nova.hypervisor.vcpus"
	gaugeOpenstackNovaHypervisorVcpusUsed             = "gauge.openstack.nova.hypervisor.vcpus_used"
	gaugeOpenstackNovaLimitMaxImageMeta               = "gauge.openstack.nova.limit.maxImageMeta"
	gaugeOpenstackNovaLimitMaxSecurityGroups          = "gauge.openstack.nova.limit.maxSecurityGroups"
	gaugeOpenstackNovaLimitMaxTotalCores              = "gauge.openstack.nova.limit.maxTotalCores"
	gaugeOpenstackNovaLimitMaxTotalFloatingIps        = "gauge.openstack.nova.limit.maxTotalFloatingIps"
	gaugeOpenstackNovaLimitMaxTotalInstances          = "gauge.openstack.nova.limit.maxTotalInstances"
	gaugeOpenstackNovaLimitMaxTotalKeypairs           = "gauge.openstack.nova.limit.maxTotalKeypairs"
	gaugeOpenstackNovaLimitMaxTotalRAMSize            = "gauge.openstack.nova.limit.maxTotalRAMSize"
	gaugeOpenstackNovaLimitTotalCoresUsed             = "gauge.openstack.nova.limit.totalCoresUsed"
	gaugeOpenstackNovaLimitTotalFloatingIpsUsed       = "gauge.openstack.nova.limit.totalFloatingIpsUsed"
	gaugeOpenstackNovaLimitTotalInstancesUsed         = "gauge.openstack.nova.limit.totalInstancesUsed"
	gaugeOpenstackNovaLimitTotalRAMUsed               = "gauge.openstack.nova.limit.totalRAMUsed"
	gaugeOpenstackNovaLimitTotalSecurityGroupsUsed    = "gauge.openstack.nova.limit.totalSecurityGroupsUsed"
	gaugeOpenstackNovaServerMemory                    = "gauge.openstack.nova.server.memory"
	gaugeOpenstackNovaServerMemoryActual              = "gauge.openstack.nova.server.memory-actual"
	gaugeOpenstackNovaServerMemoryRss                 = "gauge.openstack.nova.server.memory-rss"
	gaugeOpenstackNovaServerVdaRead                   = "gauge.openstack.nova.server.vda_read"
	gaugeOpenstackNovaServerVdaReadReq                = "gauge.openstack.nova.server.vda_read_req"
	gaugeOpenstackNovaServerVdaWrite                  = "gauge.openstack.nova.server.vda_write"
	gaugeOpenstackNovaServerVdaWriteReq               = "gauge.openstack.nova.server.vda_write_req"
)

var metricSet = map[string]monitors.MetricInfo{
	counterOpenstackNovaServerCPUTime:                 {Type: datapoint.Count},
	counterOpenstackNovaServerRx:                      {Type: datapoint.Count},
	counterOpenstackNovaServerRxPackets:               {Type: datapoint.Count},
	counterOpenstackNovaServerTx:                      {Type: datapoint.Count},
	counterOpenstackNovaServerTxPackets:               {Type: datapoint.Count},
	gaugeOpenstackCinderLimitMaxTotalBackupGigabytes:  {Type: datapoint.Gauge},
	gaugeOpenstackCinderLimitMaxTotalBackups:          {Type: datapoint.Gauge},
	gaugeOpenstackCinderLimitMaxTotalSnapshots:        {Type: datapoint.Gauge},
	gaugeOpenstackCinderLimitMaxTotalVolumeGigabytes:  {Type: datapoint.Gauge},
	gaugeOpenstackCinderLimitMaxTotalVolumes:          {Type: datapoint.Gauge},
	gaugeOpenstackCinderLimitTotalBackupGigabytesUsed: {Type: datapoint.Gauge},
	gaugeOpenstackCinderLimitTotalBackupsUsed:         {Type: datapoint.Gauge},
	gaugeOpenstackCinderLimitTotalGigabytesUsed:       {Type: datapoint.Gauge},
	gaugeOpenstackCinderLimitTotalSnapshotsUsed:       {Type: datapoint.Gauge},
	gaugeOpenstackCinderLimitTotalVolumesUsed:         {Type: datapoint.Gauge},
	gaugeOpenstackCinderSnapshotCount:                 {Type: datapoint.Gauge},
	gaugeOpenstackCinderSnapshotSize:                  {Type: datapoint.Gauge},
	gaugeOpenstackCinderVolumeCount:                   {Type: datapoint.Gauge},
	gaugeOpenstackCinderVolumeSize:                    {Type: datapoint.Gauge},
	gaugeOpenstackNeutronFloatingipCount:              {Type: datapoint.Gauge},
	gaugeOpenstackNeutronNetworkCount:                 {Type: datapoint.Gauge},
	gaugeOpenstackNeutronRouterCount:                  {Type: datapoint.Gauge},
	gaugeOpenstackNeutronSecuritygroupCount:           {Type: datapoint.Gauge},
	gaugeOpenstackNeutronSubnetCount:                  {Type: datapoint.Gauge},
	gaugeOpenstackNovaHypervisorCurrentWorkload:       {Type: datapoint.Gauge},
	gaugeOpenstackNovaHypervisorDiskAvailableLeast:    {Type: datapoint.Gauge},
	gaugeOpenstackNovaHypervisorFreeDiskGb:            {Type: datapoint.Gauge},
	gaugeOpenstackNovaHypervisorFreeRAMMb:             {Type: datapoint.Gauge},
	gaugeOpenstackNovaHypervisorLoadAverage:           {Type: datapoint.Gauge},
	gaugeOpenstackNovaHypervisorLocalGb:               {Type: datapoint.Gauge},
	gaugeOpenstackNovaHypervisorLocalGbUsed:           {Type: datapoint.Gauge},
	gaugeOpenstackNovaHypervisorMemoryMb:              {Type: datapoint.Gauge},
	gaugeOpenstackNovaHypervisorMemoryMbUsed:          {Type: datapoint.Gauge},
	gaugeOpenstackNovaHypervisorRunningVms:            {Type: datapoint.Gauge},
	gaugeOpenstackNovaHypervisorVcpus:                 {Type: datapoint.Gauge},
	gaugeOpenstackNovaHypervisorVcpusUsed:             {Type: datapoint.Gauge},
	gaugeOpenstackNovaLimitMaxImageMeta:               {Type: datapoint.Gauge},
	gaugeOpenstackNovaLimitMaxSecurityGroups:          {Type: datapoint.Gauge},
	gaugeOpenstackNovaLimitMaxTotalCores:              {Type: datapoint.Gauge},
	gaugeOpenstackNovaLimitMaxTotalFloatingIps:        {Type: datapoint.Gauge},
	gaugeOpenstackNovaLimitMaxTotalInstances:          {Type: datapoint.Gauge},
	gaugeOpenstackNovaLimitMaxTotalKeypairs:           {Type: datapoint.Gauge},
	gaugeOpenstackNovaLimitMaxTotalRAMSize:            {Type: datapoint.Gauge},
	gaugeOpenstackNovaLimitTotalCoresUsed:             {Type: datapoint.Gauge},
	gaugeOpenstackNovaLimitTotalFloatingIpsUsed:       {Type: datapoint.Gauge},
	gaugeOpenstackNovaLimitTotalInstancesUsed:         {Type: datapoint.Gauge},
	gaugeOpenstackNovaLimitTotalRAMUsed:               {Type: datapoint.Gauge},
	gaugeOpenstackNovaLimitTotalSecurityGroupsUsed:    {Type: datapoint.Gauge},
	gaugeOpenstackNovaServerMemory:                    {Type: datapoint.Gauge},
	gaugeOpenstackNovaServerMemoryActual:              {Type: datapoint.Gauge},
	gaugeOpenstackNovaServerMemoryRss:                 {Type: datapoint.Gauge},
	gaugeOpenstackNovaServerVdaRead:                   {Type: datapoint.Gauge},
	gaugeOpenstackNovaServerVdaReadReq:                {Type: datapoint.Gauge},
	gaugeOpenstackNovaServerVdaWrite:                  {Type: datapoint.Gauge},
	gaugeOpenstackNovaServerVdaWriteReq:               {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	counterOpenstackNovaServerRx:                      true,
	counterOpenstackNovaServerRxPackets:               true,
	counterOpenstackNovaServerTx:                      true,
	counterOpenstackNovaServerTxPackets:               true,
	gaugeOpenstackCinderLimitMaxTotalBackupGigabytes:  true,
	gaugeOpenstackCinderLimitMaxTotalBackups:          true,
	gaugeOpenstackCinderLimitMaxTotalSnapshots:        true,
	gaugeOpenstackCinderLimitMaxTotalVolumeGigabytes:  true,
	gaugeOpenstackCinderLimitMaxTotalVolumes:          true,
	gaugeOpenstackCinderLimitTotalBackupGigabytesUsed: true,
	gaugeOpenstackCinderLimitTotalBackupsUsed:         true,
	gaugeOpenstackCinderLimitTotalGigabytesUsed:       true,
	gaugeOpenstackCinderLimitTotalSnapshotsUsed:       true,
	gaugeOpenstackCinderLimitTotalVolumesUsed:         true,
	gaugeOpenstackCinderSnapshotCount:                 true,
	gaugeOpenstackCinderSnapshotSize:                  true,
	gaugeOpenstackCinderVolumeCount:                   true,
	gaugeOpenstackCinderVolumeSize:                    true,
	gaugeOpenstackNeutronFloatingipCount:              true,
	gaugeOpenstackNeutronNetworkCount:                 true,
	gaugeOpenstackNeutronRouterCount:                  true,
	gaugeOpenstackNeutronSecuritygroupCount:           true,
	gaugeOpenstackNeutronSubnetCount:                  true,
	gaugeOpenstackNovaHypervisorCurrentWorkload:       true,
	gaugeOpenstackNovaHypervisorDiskAvailableLeast:    true,
	gaugeOpenstackNovaHypervisorFreeDiskGb:            true,
	gaugeOpenstackNovaHypervisorFreeRAMMb:             true,
	gaugeOpenstackNovaHypervisorLocalGb:               true,
	gaugeOpenstackNovaHypervisorLocalGbUsed:           true,
	gaugeOpenstackNovaHypervisorMemoryMb:              true,
	gaugeOpenstackNovaHypervisorMemoryMbUsed:          true,
	gaugeOpenstackNovaHypervisorRunningVms:            true,
	gaugeOpenstackNovaHypervisorVcpus:                 true,
	gaugeOpenstackNovaHypervisorVcpusUsed:             true,
	gaugeOpenstackNovaLimitMaxImageMeta:               true,
	gaugeOpenstackNovaLimitMaxSecurityGroups:          true,
	gaugeOpenstackNovaLimitMaxTotalCores:              true,
	gaugeOpenstackNovaLimitMaxTotalFloatingIps:        true,
	gaugeOpenstackNovaLimitMaxTotalInstances:          true,
	gaugeOpenstackNovaLimitMaxTotalKeypairs:           true,
	gaugeOpenstackNovaLimitMaxTotalRAMSize:            true,
	gaugeOpenstackNovaLimitTotalCoresUsed:             true,
	gaugeOpenstackNovaLimitTotalFloatingIpsUsed:       true,
	gaugeOpenstackNovaLimitTotalInstancesUsed:         true,
	gaugeOpenstackNovaLimitTotalRAMUsed:               true,
	gaugeOpenstackNovaLimitTotalSecurityGroupsUsed:    true,
	gaugeOpenstackNovaServerMemory:                    true,
	gaugeOpenstackNovaServerMemoryActual:              true,
	gaugeOpenstackNovaServerMemoryRss:                 true,
	gaugeOpenstackNovaServerVdaRead:                   true,
	gaugeOpenstackNovaServerVdaReadReq:                true,
	gaugeOpenstackNovaServerVdaWrite:                  true,
	gaugeOpenstackNovaServerVdaWriteReq:               true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "collectd/openstack",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
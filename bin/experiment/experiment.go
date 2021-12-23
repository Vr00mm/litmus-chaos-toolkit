package main

import (
	"flag"
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"

	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"

	awsSSMChaosByID "github.com/Vr00mm/litmus-chaos-toolkit/experiments/aws-ssm/aws-ssm-chaos-by-id/experiment"
	awsSSMChaosByTag "github.com/Vr00mm/litmus-chaos-toolkit/experiments/aws-ssm/aws-ssm-chaos-by-tag/experiment"
	azureDiskLoss "github.com/Vr00mm/litmus-chaos-toolkit/experiments/azure/azure-disk-loss/experiment"
	azureInstanceStop "github.com/Vr00mm/litmus-chaos-toolkit/experiments/azure/instance-stop/experiment"
	redfishNodeRestart "github.com/Vr00mm/litmus-chaos-toolkit/experiments/baremetal/redfish-node-restart/experiment"
	cassandraPodDelete "github.com/Vr00mm/litmus-chaos-toolkit/experiments/cassandra/pod-delete/experiment"
	gcpVMDiskLoss "github.com/Vr00mm/litmus-chaos-toolkit/experiments/gcp/gcp-vm-disk-loss/experiment"
	gcpVMInstanceStop "github.com/Vr00mm/litmus-chaos-toolkit/experiments/gcp/gcp-vm-instance-stop/experiment"
	chaosToolkit "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/chaos-toolkit/experiment"
	containerKill "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/container-kill/experiment"
	diskFill "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/disk-fill/experiment"
	dockerServiceKill "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/docker-service-kill/experiment"
	kubeletServiceKill "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/kubelet-service-kill/experiment"
	nodeCPUHog "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/node-cpu-hog/experiment"
	nodeDrain "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/node-drain/experiment"
	nodeIOStress "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/node-io-stress/experiment"
	nodeMemoryHog "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/node-memory-hog/experiment"
	nodeRestart "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/node-restart/experiment"
	nodeTaint "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/node-taint/experiment"
	podAutoscaler "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-autoscaler/experiment"
	podCPUHogExec "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-cpu-hog-exec/experiment"
	podCPUHog "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-cpu-hog/experiment"
	podDelete "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-delete/experiment"
	podDNSError "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-dns-error/experiment"
	podDNSSpoof "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-dns-spoof/experiment"
	podFioStress "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-fio-stress/experiment"
	podIOStress "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-io-stress/experiment"
	podMemoryHogExec "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-memory-hog-exec/experiment"
	podMemoryHog "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-memory-hog/experiment"
	podNetworkCorruption "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-network-corruption/experiment"
	podNetworkDuplication "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-network-duplication/experiment"
	podNetworkLatency "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-network-latency/experiment"
	podNetworkLoss "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-network-loss/experiment"
	podNetworkPartition "github.com/Vr00mm/litmus-chaos-toolkit/experiments/generic/pod-network-partition/experiment"
	kafkaBrokerPodFailure "github.com/Vr00mm/litmus-chaos-toolkit/experiments/kafka/kafka-broker-pod-failure/experiment"
	ebsLossByID "github.com/Vr00mm/litmus-chaos-toolkit/experiments/kube-aws/ebs-loss-by-id/experiment"
	ebsLossByTag "github.com/Vr00mm/litmus-chaos-toolkit/experiments/kube-aws/ebs-loss-by-tag/experiment"
	ec2TerminateByID "github.com/Vr00mm/litmus-chaos-toolkit/experiments/kube-aws/ec2-terminate-by-id/experiment"
	ec2TerminateByTag "github.com/Vr00mm/litmus-chaos-toolkit/experiments/kube-aws/ec2-terminate-by-tag/experiment"
	vmpoweroff "github.com/Vr00mm/litmus-chaos-toolkit/experiments/vmware/vm-poweroff/experiment"

	"github.com/Vr00mm/litmus-chaos-toolkit/pkg/clients"
	"github.com/Vr00mm/litmus-chaos-toolkit/pkg/log"
	"github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		DisableSorting:         true,
		DisableLevelTruncation: true,
	})
}

func main() {

	clients := clients.ClientSets{}

	// parse the experiment name
	experimentName := flag.String("name", "pod-delete", "name of the chaos experiment")

	//Getting kubeConfig and Generate ClientSets
	if err := clients.GenerateClientSetFromKubeConfig(); err != nil {
		log.Errorf("Unable to Get the kubeconfig, err: %v", err)
		return
	}

	log.Infof("Experiment Name: %v", *experimentName)

	// invoke the corresponding experiment based on the the (-name) flag
	switch *experimentName {
	case "chaos-toolkit":
		chaosToolkit.Chaos (clients)
	case "container-kill":
		containerKill.ContainerKill(clients)
	case "disk-fill":
		diskFill.DiskFill(clients)
	case "kafka-broker-pod-failure":
		kafkaBrokerPodFailure.KafkaBrokerPodFailure(clients)
	case "kubelet-service-kill":
		kubeletServiceKill.KubeletServiceKill(clients)
	case "docker-service-kill":
		dockerServiceKill.DockerServiceKill(clients)
	case "node-cpu-hog":
		nodeCPUHog.NodeCPUHog(clients)
	case "node-drain":
		nodeDrain.NodeDrain(clients)
	case "node-io-stress":
		nodeIOStress.NodeIOStress(clients)
	case "node-memory-hog":
		nodeMemoryHog.NodeMemoryHog(clients)
	case "node-taint":
		nodeTaint.NodeTaint(clients)
	case "pod-autoscaler":
		podAutoscaler.PodAutoscaler(clients)
	case "pod-cpu-hog-exec":
		podCPUHogExec.PodCPUHogExec(clients)
	case "pod-delete":
		podDelete.PodDelete(clients)
	case "pod-io-stress":
		podIOStress.PodIOStress(clients)
	case "pod-memory-hog-exec":
		podMemoryHogExec.PodMemoryHogExec(clients)
	case "pod-network-corruption":
		podNetworkCorruption.PodNetworkCorruption(clients)
	case "pod-network-duplication":
		podNetworkDuplication.PodNetworkDuplication(clients)
	case "pod-network-latency":
		podNetworkLatency.PodNetworkLatency(clients)
	case "pod-network-loss":
		podNetworkLoss.PodNetworkLoss(clients)
	case "pod-network-partition":
		podNetworkPartition.PodNetworkPartition(clients)
	case "pod-memory-hog":
		podMemoryHog.PodMemoryHog(clients)
	case "pod-cpu-hog":
		podCPUHog.PodCPUHog(clients)
	case "cassandra-pod-delete":
		cassandraPodDelete.CasssandraPodDelete(clients)
	case "aws-ssm-chaos-by-id":
		awsSSMChaosByID.AWSSSMChaosByID(clients)
	case "aws-ssm-chaos-by-tag":
		awsSSMChaosByTag.AWSSSMChaosByTag(clients)
	case "ec2-terminate-by-id":
		ec2TerminateByID.EC2TerminateByID(clients)
	case "ec2-terminate-by-tag":
		ec2TerminateByTag.EC2TerminateByTag(clients)
	case "ebs-loss-by-id":
		ebsLossByID.EBSLossByID(clients)
	case "ebs-loss-by-tag":
		ebsLossByTag.EBSLossByTag(clients)
	case "node-restart":
		nodeRestart.NodeRestart(clients)
	case "pod-dns-error":
		podDNSError.PodDNSError(clients)
	case "pod-dns-spoof":
		podDNSSpoof.PodDNSSpoof(clients)
	case "vm-poweroff":
		vmpoweroff.VMPoweroff(clients)
	case "azure-instance-stop":
		azureInstanceStop.AzureInstanceStop(clients)
	case "azure-disk-loss":
		azureDiskLoss.AzureDiskLoss(clients)
	case "gcp-vm-disk-loss":
		gcpVMDiskLoss.VMDiskLoss(clients)
	case "pod-fio-stress":
		podFioStress.PodFioStress(clients)
	case "gcp-vm-instance-stop":
		gcpVMInstanceStop.VMInstanceStop(clients)
	case "redfish-node-restart":
		redfishNodeRestart.NodeRestart(clients)

	default:
		log.Errorf("Unsupported -name %v, please provide the correct value of -name args", *experimentName)
		return
	}
}

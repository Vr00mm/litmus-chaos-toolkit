package duplication

import (
	"strconv"

	network_chaos "github.com/Vr00mm/litmus-chaos-toolkit/chaoslib/pumba/network-chaos/lib"
	clients "github.com/Vr00mm/litmus-chaos-toolkit/pkg/clients"
	experimentTypes "github.com/Vr00mm/litmus-chaos-toolkit/pkg/generic/network-chaos/types"
	"github.com/Vr00mm/litmus-chaos-toolkit/pkg/types"
)

//PodNetworkDuplicationChaos contains the steps to prepare and inject chaos
func PodNetworkDuplicationChaos(experimentsDetails *experimentTypes.ExperimentDetails, clients clients.ClientSets, resultDetails *types.ResultDetails, eventsDetails *types.EventDetails, chaosDetails *types.ChaosDetails) error {

	args, err := getContainerArguments(experimentsDetails)
	if err != nil {
		return err
	}
	return network_chaos.PrepareAndInjectChaos(experimentsDetails, clients, resultDetails, eventsDetails, chaosDetails, args)
}

// getContainerArguments derives the args for the pumba pod
func getContainerArguments(experimentsDetails *experimentTypes.ExperimentDetails) ([]string, error) {
	baseArgs := []string{
		"pumba",
		"netem",
		"--tc-image",
		experimentsDetails.TCImage,
		"--interface",
		experimentsDetails.NetworkInterface,
		"--duration",
		strconv.Itoa(experimentsDetails.ChaosDuration) + "s",
	}

	args := baseArgs
	args, err := network_chaos.AddTargetIpsArgs(experimentsDetails.DestinationIPs, experimentsDetails.DestinationHosts, args)
	if err != nil {
		return args, err
	}
	args = append(args, "duplicate", "--percent", strconv.Itoa(experimentsDetails.NetworkPacketDuplicationPercentage))

	return args, nil
}

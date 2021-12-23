package latency

import (
	"strconv"

	network_chaos "github.com/Vr00mm/litmus-chaos-toolkit/chaoslib/litmus/network-chaos/lib"
	clients "github.com/Vr00mm/litmus-chaos-toolkit/pkg/clients"
	experimentTypes "github.com/Vr00mm/litmus-chaos-toolkit/pkg/generic/network-chaos/types"
	"github.com/Vr00mm/litmus-chaos-toolkit/pkg/types"
)

//PodNetworkLatencyChaos contains the steps to prepare and inject chaos
func PodNetworkLatencyChaos(experimentsDetails *experimentTypes.ExperimentDetails, clients clients.ClientSets, resultDetails *types.ResultDetails, eventsDetails *types.EventDetails, chaosDetails *types.ChaosDetails) error {

	args := "delay " + strconv.Itoa(experimentsDetails.NetworkLatency) + "ms"
	return network_chaos.PrepareAndInjectChaos(experimentsDetails, clients, resultDetails, eventsDetails, chaosDetails, args)
}

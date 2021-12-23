package duplication

import (
	"strconv"

	network_chaos "github.com/Vr00mm/litmus-chaos-toolkit/chaoslib/litmus/network-chaos/lib"
	clients "github.com/Vr00mm/litmus-chaos-toolkit/pkg/clients"
	experimentTypes "github.com/Vr00mm/litmus-chaos-toolkit/pkg/generic/network-chaos/types"
	"github.com/Vr00mm/litmus-chaos-toolkit/pkg/types"
)

//PodNetworkDuplicationChaos contains the steps to prepare and inject chaos
func PodNetworkDuplicationChaos(experimentsDetails *experimentTypes.ExperimentDetails, clients clients.ClientSets, resultDetails *types.ResultDetails, eventsDetails *types.EventDetails, chaosDetails *types.ChaosDetails) error {

	args := "duplicate " + strconv.Itoa(experimentsDetails.NetworkPacketDuplicationPercentage)
	return network_chaos.PrepareAndInjectChaos(experimentsDetails, clients, resultDetails, eventsDetails, chaosDetails, args)
}

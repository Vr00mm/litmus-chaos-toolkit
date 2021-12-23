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

	containerKill "github.com/Vr00mm/litmus-chaos-toolkit/chaoslib/litmus/container-kill/helper"
	diskFill "github.com/Vr00mm/litmus-chaos-toolkit/chaoslib/litmus/disk-fill/helper"
	networkChaos "github.com/Vr00mm/litmus-chaos-toolkit/chaoslib/litmus/network-chaos/helper"
	dnsChaos "github.com/Vr00mm/litmus-chaos-toolkit/chaoslib/litmus/pod-dns-chaos/helper"
	stressChaos "github.com/Vr00mm/litmus-chaos-toolkit/chaoslib/litmus/stress-chaos/helper"

	"github.com/Vr00mm/litmus-chaos-toolkit/pkg/clients"
	"github.com/Vr00mm/litmus-chaos-toolkit/pkg/log"
	"github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		DisableSorting:         true,
		DisableLevelTruncation: true,
	})
}

func main() {

	clients := clients.ClientSets{}

	// parse the helper name
	helperName := flag.String("name", "", "name of the helper pod")

	//Getting kubeConfig and Generate ClientSets
	if err := clients.GenerateClientSetFromKubeConfig(); err != nil {
		log.Errorf("Unable to Get the kubeconfig, err: %v", err)
		return
	}

	log.Infof("Helper Name: %v", *helperName)

	// invoke the corresponding helper based on the the (-name) flag
	switch *helperName {
	case "container-kill":
		containerKill.Helper(clients)
	case "disk-fill":
		diskFill.Helper(clients)
	case "dns-chaos":
		dnsChaos.Helper(clients)
	case "stress-chaos":
		stressChaos.Helper(clients)
	case "network-chaos":
		networkChaos.Helper(clients)

	default:
		log.Errorf("Unsupported -name %v, please provide the correct value of -name args", *helperName)
		return
	}
}

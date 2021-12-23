package types

import (
	exp "github.com/Vr00mm/litmus-chaos-toolkit/pkg/generic/pod-delete/types"
)

// ExperimentDetails is for collecting all the experiment-related details
type ExperimentDetails struct {
	ChaoslibDetail         *exp.ExperimentDetails
	CassandraServiceName   string
	KeySpaceReplicaFactor  string
	CassandraPort          int
	LivenessServicePort    int
	CassandraLivenessImage string
	CassandraLivenessCheck string
	RunID                  string
	Sequence               string
}

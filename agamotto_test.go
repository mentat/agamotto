package agamotto

import (
	"testing"
)

func TestDocs(t *testing.T) {

	eye, _ := NewEye()

	op, err := eye.ManageCluster(&Cluster{
		Topography:       Flat, // Can be Flat, MasterSlave,
		CorrectiveAction: Heal, // Can be Heal, Ignore, Shutdown
		ControlNodes: []Node{
			Node{
				BaseImage: "sentinal",
			},
		},
		DataNodes: []Node{
			Node{
				BaseImage: "redis",
			},
		},
		DataNodeCount: 3,
	})

	channel, err := eye.Observe()
	if err != nil {
		t.Fatalf("Count not observe infrastructure: %s", err)
	}

	event := <-channel

	if event.ClusterID != op.Cluster.ID {
		t.Fatalf("Weird.")
	}

}

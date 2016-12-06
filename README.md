# Agamotto is a Clairvoyant System to Manage Clusters

Agamotto manages an infrastructure by observing the health of all the nodes in 
each of the defined clusters of the datacenter.  When a node in a cluster has an issue,
Agamotto will take a corrective action.  Agamotto is not a provisioning layer. Agamotto
will notify the provisioning layer of the suggested corrective action to take based
on the cluster configuration.

# Usage
```go
import "agamotto"

eye, err := agamotto.NewEye()

op, err := eye.ManageCluster(agamotto.Cluster{
    Topography: agamotto.Flat, // Can be Flat, MasterSlave,
    CorrectiveAction: agamoto.Heal, // Can be Heal, Ignore, Shutdown
    ControlNodes: []agamotto.Node{
		Node{
			BaseImage: "sentinal",
		},
	},
	DataNodes: []agamotto.Node{
		Node{
			BaseImage: "redis",
		},
	},
	DataNodeCount: 3,
})

channel, err := eye.Observe()

// Wait for actions to come through

action := <-channel

// Provision something based on that action.
```

## Leader Election

When the cluster topography is set to __agamotto.MasterSlave__ and a master fails, a new leader must
be elected.  Agamotto supports a per-cluster election heuristic that is mandatory for all
__agamotto.MasterSlave__ topographies.

```go
op, err := eye.ManageCluster(agamotto.Cluster{
    Election: ElectionHeuristic{
        Field: "replication_lag",
        Operation: agamotto.Lowest,
    },
    // ...
})
```

The value in ElectionHeuristic.Field is the name of a key returned by the nodes in the cluster
during a health check event.  Nested fields are supported through "." (dot) notation.  I.e.:

```go
op, err := eye.ManageCluster(agamotto.Cluster{
    Election: ElectionHeuristic{
        Field: "slave.replication.lag",
        Operation: agamotto.Lowest,
    },
    // ...
})
```
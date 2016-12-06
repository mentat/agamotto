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

op, err := eye.ManageCluster(agamotoo.Cluster{
    Topography: agamotto.Flat, // Can be Flat, MasterSlave,
    CorrectiveAction: agamoto.Heal, // Can be Heal, Ignore, Shutdown
    ControlNodes: []agamotto.Node{
        agamotto.Node{
            Image: "blah",
            Count: 1,
        },
    },
    DataNodes: []agamotto.Node{
        agamotto.Node{
            Image: "blah",
            Count: 4,
        },
    },
})
```
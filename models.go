package agamotto

import (
	"time"
)

// Topography - how the cluster is arranged.
type Topography string

// Topographies
const (
	Flat        Topography = "flat"
	MasterSlave Topography = "master/slave"
)

// Correction - types of corrections to make.
type Correction string

// Corrective actions
const (
	Heal     Correction = "heal"
	Ignore   Correction = "ignore"
	Shutdown Correction = "shutdown"
)

// OperationStatus - the status of an operation.
type OperationStatus string

// Operation Statuses
const (
	Pending OperationStatus = "pending"
	Running OperationStatus = "running"
	Done    OperationStatus = "done"
)

// ActionType - an action that the provisioning system should take.
type ActionType string

const (
	Create ActionType = "create"
	Remove ActionType = "remove"
)

type ElectionOperator string

const (
	Lowest  ElectionOperator = "lowest"  // The lowest value
	Highest ElectionOperator = "highest" // The highest value
	Random  ElectionOperator = "random"  // Pick a random value
)

// ElectionHeuristic - a simple heuristic for choosing a new leader.
type ElectionHeuristic struct {
	Field    string
	Operator ElectionOperator
}

// Node - a node in a cluster.
type Node struct {
	ID          string
	Address     string
	ServiceType string
	BaseImage   string
	CreatedAt   string
	IsError     bool
	Error       string
	Metadata    map[string]interface{}
	IsMaster    bool
}

// Cluster - an independant collection of nodes that interact.
type Cluster struct {
	ID               uint
	Topography       Topography
	CorrectiveAction Correction
	ControlNodes     []Node
	DataNodes        []Node
	DataNodeCount    int // The number of data nodes to maintain in the cluster.
	Election         ElectionHeuristic
	Callback         string // An HTTP callback to fire when an action is need, optional.
}

// Action - an action that the provisioning layer should take.
type Action struct {
	ID           string
	Type         ActionType
	Count        int
	BaseImage    string
	DataNodes    []Node
	ControlNodes []Node
	ClusterID    uint
	Metadata     map[string]interface{}
}

// Operation - an internal long running operation. Generally tracks the
// status of an Action.
type Operation struct {
	ID          uint
	StartedAt   time.Time
	CompletedAt time.Time
	ConfirmedAt time.Time
	Status      OperationStatus
	WasSuccess  bool
	Error       string
	Action      *Action
	Cluster     *Cluster
	RetryCount  int
	Confirmed   bool // Vision has confirmed the operation success.
}

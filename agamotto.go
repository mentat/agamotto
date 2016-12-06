package agamotto

import "github.com/jinzhu/gorm"

// Eye - the clairvoyant piece.
type Eye struct {
	observation chan Action
	db          *gorm.DB
}

// NewEye - create a new Eye of Agamotto.
func NewEye(db *gorm.DB) (*Eye, error) {
	eye := &Eye{
		db: db,
	}
	return eye, nil
}

// ManageCluster - manage the lifecycle of a cluster.
func (eye *Eye) ManageCluster(cluster *Cluster) (*Operation, error) {

	return nil, nil
}

// ScaleCluster - scale the data nodes in a cluster.  Returns an Operation.
func (eye *Eye) ScaleCluster(clusterID uint, newCount int) (*Operation, error) {
	// Check the current size of the data cluster and send actions
	// as appropriate to scale up the cluster.  Each node to be added to the cluster
	// will be it's own idempotent operation/action pair.
	// Important note: operations on a cluster should happen in order, so we need to
	// ensure no other operations are running for the cluster before we notify out
	// a new one.
	return nil, nil
}

// ForgetCluster - forget that a cluster exists and stop managing it.
func (eye *Eye) ForgetCluster(clusterID uint) error {

	return nil
}

// CompleteAction - notify Agamotto that another layer has completed an action.
func (eye Eye) CompleteAction(actionID int) {

}

// Observe - creates a channel from which events flow to the calling
// system.  Agamotto will request that the observing system create
// resources as requested in order to maintain cluster health.
func (eye *Eye) Observe() (chan Action, error) {

	if eye.observation == nil {
		eye.observation = make(chan Action, 50)
	}

	return eye.observation, nil
}

func (eye Eye) notify(action *Action) {
	if eye.observation != nil {
		eye.observation <- *action
	} else {
		// Check if the cluster has a configured callback.
	}
}

// Initialize - initialize database and connect to required services.
func Initialize(db *gorm.DB) error {
	// Initialize Agamotto.  Create tables and such.
	return nil
}

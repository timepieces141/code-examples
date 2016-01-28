package rackspace

type Rackspace struct {
	name string
}

func NewRackspace() *Rackspace {
	return &Rackspace{"rackspace"}
}

func (cloud Rackspace) ServiceName() string {
	return cloud.name
}

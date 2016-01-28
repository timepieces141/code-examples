package rackspace

type RackspaceFlavor struct {
	flavorID string
	name string
	cpus int
	ram int
	disk int
}

func NewRackspaceFlavor(flavorID string, name string, cpus int, ram int, disk int) RackspaceFlavor {
	return RackspaceFlavor{flavorID, name, cpus, ram, disk}
}

func (rf RackspaceFlavor) FlavorID() string {
	return rf.flavorID
}

func (rf RackspaceFlavor) Name() string {
	return rf.name
}

func (rf RackspaceFlavor) CPUs() int {
	return rf.cpus
}

func (rf RackspaceFlavor) RAM() int {
	return rf.ram
}

func (rf RackspaceFlavor) Disk() int {
	return rf.disk
}
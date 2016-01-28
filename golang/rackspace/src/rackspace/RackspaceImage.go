package rackspace

type RackspaceImage struct {
	imageID string
	name string
}

func NewRackspaceImage(imageID string, name string) RackspaceImage {
	return RackspaceImage{imageID, name}
}

func (ri RackspaceImage) ImageID() string {
	return ri.imageID
}

func (ri RackspaceImage) Name() string {
	return ri.name
}
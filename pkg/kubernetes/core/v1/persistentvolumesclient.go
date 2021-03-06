// Code generated by helmit-generate. DO NOT EDIT.

package v1

import (
	"github.com/onosproject/helmit/pkg/kubernetes/resource"
)

type PersistentVolumesClient interface {
	PersistentVolumes() PersistentVolumesReader
}

func NewPersistentVolumesClient(resources resource.Client, filter resource.Filter) PersistentVolumesClient {
	return &persistentVolumesClient{
		Client: resources,
		filter: filter,
	}
}

type persistentVolumesClient struct {
	resource.Client
	filter resource.Filter
}

func (c *persistentVolumesClient) PersistentVolumes() PersistentVolumesReader {
	return NewPersistentVolumesReader(c.Client, c.filter)
}

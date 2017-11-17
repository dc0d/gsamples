package node

import (
	_ "github.com/dc0d/gsamples/gmap"
)

type Node struct {
	Name     string
	Children GMap
}

// New should return pointer? BTW this is just a demonstration.
func New() Node {
	var n Node
	n.Children = make(map[TGMap]UGMap)
	return n
}

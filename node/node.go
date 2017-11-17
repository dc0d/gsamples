package node

import (
	_ "github.com/dc0d/gsamples/gmapdual"
)

type Type struct{}

type Node struct {
	Name   string
	ByName Gmap1
	ByType Gmap2
}

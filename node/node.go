package node

import (
	_ "github.com/dc0d/gsamples/gmapprime"
)

type Type struct{}

type Node struct {
	Name   string
	ByName GMap
	ByType GMapPrime
}

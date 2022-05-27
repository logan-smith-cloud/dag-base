package tdag

import (
	"github.com/galaxy-digital/lachesis-base/hash"
	"github.com/galaxy-digital/lachesis-base/inter/dag"
)

type TestEvent struct {
	dag.MutableBaseEvent
	Name string
}

func (e *TestEvent) AddParent(id hash.Event) {
	parents := e.Parents()
	parents.Add(id)
	e.SetParents(parents)
}

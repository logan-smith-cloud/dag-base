package tdag

import (
	"github.com/logan-smith-cloud/dag-base/hash"
	"github.com/logan-smith-cloud/dag-base/inter/dag"
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

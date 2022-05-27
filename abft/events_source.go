package abft

import (
	"github.com/logan-smith-cloud/dag-base/hash"
	"github.com/logan-smith-cloud/dag-base/inter/dag"
)

// EventSource is a callback for getting events from an external storage.
type EventSource interface {
	HasEvent(hash.Event) bool
	GetEvent(hash.Event) dag.Event
}

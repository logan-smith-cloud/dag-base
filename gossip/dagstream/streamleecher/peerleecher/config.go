package peerleecher

import (
	"time"

	"github.com/logan-smith-cloud/dag-base/inter/dag"
)

type EpochDownloaderConfig struct {
	RecheckInterval        time.Duration
	DefaultChunkSize       dag.Metric
	ParallelChunksDownload int
}

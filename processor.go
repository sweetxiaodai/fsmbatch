package fsmbatch

import (
	"context"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type resourceProcessor struct {
	logger *zap.Logger
}

func (rp *resourceProcessor) processTraces(ctx context.Context, td ptrace.Traces) (ptrace.Traces, error) {
	return td, nil
}

func (rp *resourceProcessor) processMetrics(ctx context.Context, md pmetric.Metrics) (pmetric.Metrics, error) {
	return md, nil
}

func (rp *resourceProcessor) processLogs(ctx context.Context, ld plog.Logs) (plog.Logs, error) {
	return ld, nil
}

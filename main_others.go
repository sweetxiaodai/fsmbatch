// Code generated by "go.opentelemetry.io/collector/cmd/builder". DO NOT EDIT.

//go:build !windows
// +build !windows

package fsmbatch

import "go.opentelemetry.io/collector/otelcol"

func run(params otelcol.CollectorSettings) error {
	return runInteractive(params)
}

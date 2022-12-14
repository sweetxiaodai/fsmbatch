package main

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
)

const (
	typeStr   = "fsmbatch"
	stability = component.StabilityLevelBeta
)

var processorCapabilities = consumer.Capabilities{MutatesData: true}

func NewFactory() component.ProcessorFactory {
	return component.NewProcessorFactory(
		typeStr,
		createDefaultConfig,
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		ProcessorSettings: config.NewProcessorSettings(component.NewID(typeStr)),
	}
}

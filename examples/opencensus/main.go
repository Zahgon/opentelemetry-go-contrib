// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Opencensus exemplifies the use of the OpenCensus to OpenTelemetry bridge.
package main

import (
	"fmt"
	"log"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var (
	// instrumenttype differentiates between our gauge and view metrics.
	keyType = tag.MustNewKey("instrumenttype")
	// Counts the number of lines read in from standard input.
	countMeasure = stats.Int64("test_count", "A count of something", stats.UnitDimensionless)
	countView    = &view.View{
		Name:        "test_count",
		Measure:     countMeasure,
		Description: "A count of something",
		Aggregation: view.Count(),
		TagKeys:     []tag.Key{keyType},
	}
)

func main() {
	log.Println("Using OpenTelemetry stdout exporters.")
	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(fmt.Errorf("error creating trace exporter: %w", err))
	}
	metricsExporter, err := stdoutmetric.New()
	if err != nil {
		log.Fatal(fmt.Errorf("error creating metric exporter: %w", err))
	}
	tracing(traceExporter)
	if err := monitoring(metricsExporter); err != nil {
		log.Fatal(err)
	}
}

// tracing demonstrates overriding the OpenCensus DefaultTracer to send spans
// to the OpenTelemetry exporter by calling OpenCensus APIs.
func tracing(otExporter sdktrace.SpanExporter) { _ = "STUB: not implemented"; return }

// monitoring demonstrates creating an IntervalReader using the OpenTelemetry
// exporter to send metrics to the exporter by using either an OpenCensus
// registry or an OpenCensus view.
func monitoring(exporter metric.Exporter) error { _ = "STUB: not implemented"; return nil }

// Register the OpenCensus metric Producer to add metrics from OpenCensus to the output.

// update stats for our gauge

// update stats for our view

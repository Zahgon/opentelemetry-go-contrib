// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Example exemplifies the use of the otellambda instrumentation.
package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	lambdadetector "go.opentelemetry.io/contrib/detectors/aws/lambda"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
)

func lambdaHandler(ctx context.Context) error {
	_ = "STUB: not implemented"
	// init aws config
	return nil
}

// instrument all aws clients

// S3

// HTTP

func main() {
	ctx := context.Background()

	exp, err := stdouttrace.New()
	if err != nil {
		log.Printf("failed to initialize stdout exporter %v\n", err)
		return
	}

	detector := lambdadetector.NewResourceDetector()
	res, err := detector.Detect(ctx)
	if err != nil {
		log.Fatalf("failed to detect lambda resources: %v\n", err)
		return
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSyncer(exp),
		sdktrace.WithResource(res),
	)

	// Downstream spans use global tracer provider
	otel.SetTracerProvider(tp)

	lambda.Start(otellambda.InstrumentHandler(lambdaHandler, otellambda.WithTracerProvider(tp)))
}

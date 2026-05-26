// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelaws // import "go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws"

import (
	"context"

	"github.com/aws/smithy-go/middleware"
	"go.opentelemetry.io/otel/attribute"
)

// DynamoDBAttributeBuilder sets DynamoDB specific attributes depending on the DynamoDB operation being performed.
func DynamoDBAttributeBuilder(_ context.Context, in middleware.InitializeInput, _ middleware.InitializeOutput) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

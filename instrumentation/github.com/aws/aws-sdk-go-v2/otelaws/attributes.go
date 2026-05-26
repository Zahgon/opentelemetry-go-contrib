// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelaws // import "go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws"

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/smithy-go/middleware"
	"go.opentelemetry.io/otel/attribute"
)

// AWS attributes.
const (
	RegionKey    attribute.Key = "aws.region"
	RequestIDKey attribute.Key = "aws.request_id"
	AWSSystemVal string        = "aws-api"
)

var servicemap = map[string]AttributeBuilder{
	dynamodb.ServiceID: DynamoDBAttributeBuilder,
	sqs.ServiceID:      SQSAttributeBuilder,
	sns.ServiceID:      SNSAttributeBuilder,
}

// SystemAttr return the AWS RPC system attribute.
func SystemAttr() attribute.KeyValue { _ = "STUB: not implemented"; return *new(attribute.KeyValue) }

// MethodAttr returns the RPC method attribute for the AWS service and operation.
func MethodAttr(service, operation string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

// OperationAttr returns the AWS operation attribute.
//
// Deprecated: use [MethodAttr] instead.
func OperationAttr(operation string) attribute.KeyValue {
	_ = "STUB: not implemented"
	// rpc.service has been merged into rpc.method in semconv v1.39.0
	return *new(attribute.KeyValue)
}

// RegionAttr returns the AWS region attribute.
func RegionAttr(region string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

// ServiceAttr returns the AWS service attribute.
//
// Deprecated: use [MethodAttr] instead.
func ServiceAttr(service string) attribute.KeyValue {
	_ = "STUB: not implemented"
	// rpc.service has been merged into rpc.method in semconv v1.39.0
	return *new(attribute.KeyValue)
}

// RequestIDAttr returns the AWS request ID attribute.
func RequestIDAttr(requestID string) attribute.KeyValue {
	_ = "STUB: not implemented"
	return *new(attribute.KeyValue)
}

// DefaultAttributeBuilder checks to see if there are service specific attributes available to set for the AWS service.
// If there are service specific attributes available then they will be included.
func DefaultAttributeBuilder(ctx context.Context, in middleware.InitializeInput, out middleware.InitializeOutput) []attribute.KeyValue {
	_ = "STUB: not implemented"
	return nil
}

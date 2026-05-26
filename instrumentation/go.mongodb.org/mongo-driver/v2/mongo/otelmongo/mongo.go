// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelmongo // import "go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/v2/mongo/otelmongo"

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/event"
	"go.opentelemetry.io/otel/semconv/v1.41.0/dbconv"
	"go.opentelemetry.io/otel/trace"
)

type spanKey struct {
	ConnectionID string
	RequestID    int64
}

type monitor struct {
	ClientOperationDuration *dbconv.ClientOperationDuration

	sync.Mutex
	spans map[spanKey]trace.Span
	cfg   config
}

func (m *monitor) Started(ctx context.Context, evt *event.CommandStartedEvent) {
	_ = "STUB: not implemented"
	return
}

func (m *monitor) Succeeded(ctx context.Context, evt *event.CommandSucceededEvent) {
	_ = "STUB: not implemented"
	return
}

// No need to add semconv.DBSystemMongoDB, it will be added by metrics recorder.

// `db.response.status_code` is excluded for succeeded events.
// Succeeded processes an [go.mongodb.org/mongo-driver/v2/event.CommandSucceededEvent] for OTel,
// including collecting metrics. The status code metric is excluded since MongoDB server indicates
// a successful operation with {ok: 1}, which doesn't map to a traditional status code.

// TODO: db.query.text attribute is currently disabled by default.
// Because event does not provide the query text directly.
// command := m.extractCommand(evt)
// attrs = append(attrs, semconv.DBQueryText(sanitizeCommand(evt.Command)))

func (m *monitor) Failed(ctx context.Context, evt *event.CommandFailedEvent) {
	_ = "STUB: not implemented"
	return
}

// TODO: The status code should not be static, but reflect server behavior.
// Assert the error as [go.mongodb.org/mongo-driver/v2/x/mongo/driver.Error] and pull the code from there.
// ref. https://jira.mongodb.org/browse/GODRIVER-3690

// TODO: db.query.text attribute is currently disabled by default.
// Because event does not provide the query text directly.
// command := m.extractCommand(evt)
// attrs = append(attrs, semconv.DBQueryText(sanitizeCommand(evt.Command)))

func (m *monitor) Finished(evt *event.CommandFinishedEvent, err error) {
	_ = "STUB: not implemented"
	return
}

// TODO sanitize values where possible, then re-enable `db.statement` span attributes default.
// TODO limit maximum size.
func sanitizeCommand(command bson.Raw) string { _ = "STUB: not implemented"; return "" }

// extractCollection extracts the collection for the given mongodb command event.
// For CRUD operations, this is the first key/value string pair in the bson
// document where key == "<operation>" (e.g. key == "insert").
// For database meta-level operations, such a key may not exist.
// This function returns the collection name or an error if no collection can be determined.
func extractCollection(evt *event.CommandStartedEvent) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// NewMonitor creates a new mongodb event CommandMonitor.
func NewMonitor(opts ...Option) *event.CommandMonitor { _ = "STUB: not implemented"; return nil }

// peerInfo will parse the hostname and port from the mongo connection ID.
func peerInfo(connectionID string) (hostname string, port int) {
	_ = "STUB: not implemented"
	return "", 0
}

// If parsing fails, assume default MongoDB port and return the entire ConnectionID as hostname

// If port parsing fails, fallback to default MongoDB port

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Copyright 2017, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zpages // import "go.opentelemetry.io/contrib/zpages"

import (
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

const (
	// spanNameQueryField is the header for span name.
	spanNameQueryField = "zspanname"
	// spanTypeQueryField is the header for type (running = 0, latency = 1, error = 2) to display.
	spanTypeQueryField = "ztype"
	// spanLatencyBucketQueryField is the header for latency based samples.
	// Default is [0, 8] representing the latency buckets, where 0 is the first one.
	spanLatencyBucketQueryField = "zlatencybucket"
	// maxTraceMessageLength is the maximum length of a message in tracez output.
	maxTraceMessageLength = 1024

	maxRequestBodySize = 1 << 20 // 1MB
)

type summaryTableData struct {
	Header             []string
	LatencyBucketNames []string
	Links              bool
	TracesEndpoint     string
	Rows               []summaryTableRowData
}

type summaryTableRowData struct {
	Name    string
	Active  int
	Latency []int
	Errors  int
}

// traceTableData contains data for the trace data template.
type traceTableData struct {
	Name string
	Num  int
	Rows []spanRow
}

var _ http.Handler = (*tracezHandler)(nil)

type tracezHandler struct {
	sp *SpanProcessor
}

// NewTracezHandler returns an http.Handler that can be used to serve HTTP requests for trace zpages.
func NewTracezHandler(sp *SpanProcessor) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// ServeHTTP implements the http.Handler and is capable of serving "tracez" HTTP requests.
func (th *tracezHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (th *tracezHandler) getTraceTableData(spanName string, spanType, latencyBucket int) traceTableData {
	_ = "STUB: not implemented"
	return *new(traceTableData)
}

// active

// latency

// error

func (th *tracezHandler) getSummaryTableData() summaryTableData {
	_ = "STUB: not implemented"
	return *new(summaryTableData)
}

// An implicit 0 lower bound latency bucket is always present.

type spanRow struct {
	Fields [3]string
	trace.SpanContext
	ParentSpanContext trace.SpanContext
}

type events []sdktrace.Event

func (e events) Len() int           { _ = "STUB: not implemented"; return 0 }
func (e events) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (e events) Swap(i, j int) { _ = "STUB: not implemented"; return }

type attributes []attribute.KeyValue

func (e attributes) Len() int           { _ = "STUB: not implemented"; return 0 }
func (e attributes) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (e attributes) Swap(i, j int) { _ = "STUB: not implemented"; return }

func spanRows(s sdktrace.ReadOnlySpan) []spanRow { _ = "STUB: not implemented"; return nil }

// There are five cases for duration printing:
// -1234567890s
// -1234.123456
//      .123456
// 12345.123456
// 12345678901s

//nolint:staticcheck // Use deprecated method for formatting backward compatibility.

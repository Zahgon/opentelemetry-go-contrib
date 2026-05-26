// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Copyright (c) 2021 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
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

// Package testutils provides testing utilities for the jaegerremote sampler
// package.
package testutils // import "go.opentelemetry.io/contrib/samplers/jaegerremote/internal/testutils"

import (
	"net/http"
	"net/http/httptest"
)

// MockAgent is a mock representation of Jaeger Agent.
// It has an HTTP endpoint for sampling strategies.
type MockAgent struct {
	samplingMgr *samplingManager
	samplingSrv *httptest.Server
}

// StartMockAgent runs a mock representation of jaeger-agent.
// This function returns a started server.
func StartMockAgent() (*MockAgent, error) { _ = "STUB: not implemented"; return nil, nil }

// Close stops the serving of traffic.
func (s *MockAgent) Close() { _ = "STUB: not implemented"; return }

// SamplingServerAddr returns the host:port of HTTP server exposing sampling strategy endpoint.
func (s *MockAgent) SamplingServerAddr() string { _ = "STUB: not implemented"; return "" }

// AddSamplingStrategy registers a sampling strategy for a service.
func (s *MockAgent) AddSamplingStrategy(service string, strategy any) {
	_ = "STUB: not implemented"
	return
}

type samplingHandler struct {
	manager *samplingManager
}

func (h *samplingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

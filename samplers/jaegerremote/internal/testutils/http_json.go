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

package testutils // import "go.opentelemetry.io/contrib/samplers/jaegerremote/internal/testutils"

import (
	"context"
	"net/http"
)

// getJSON makes an HTTP call to the specified URL and parses the returned JSON into `out`.
func getJSON(ctx context.Context, url string, out any) error { _ = "STUB: not implemented"; return nil }

// readJSON reads JSON from http.Response and parses it into `out`.
func readJSON(resp *http.Response, out any) error { _ = "STUB: not implemented"; return nil }

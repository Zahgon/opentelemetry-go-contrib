// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelhttptrace // import "go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"

import (
	"context"
	"net/http"
)

// W3C client.
func W3C(ctx context.Context, req *http.Request) (context.Context, *http.Request) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}

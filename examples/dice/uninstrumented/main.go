// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Uninstrumented provides an example rolldice service that is not instrumented
// with observability.
package main

import (
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() (err error) {
	_ = "STUB: not implemented"
	// Handle SIGINT (CTRL+C) gracefully.
	return nil
}

// Start HTTP server.

// Wait for interruption.

// Error when starting HTTP server.

// Wait for first CTRL+C.
// Stop receiving signal notifications as soon as possible.

// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.

func newHTTPHandler() http.Handler {
	_ = "STUB: not implemented"
	return *

	// Register handlers.
	new(http.Handler)
}

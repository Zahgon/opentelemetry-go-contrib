// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Instrumented provides an example rolldice service that is instrumented with
// OpenTelemetry.
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

func run() error {
	_ = "STUB: not implemented"
	// Handle SIGINT (CTRL+C) gracefully.
	return nil
}

// Set up OpenTelemetry.

// Handle shutdown properly so nothing leaks.

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

// Add HTTP instrumentation for the whole server.

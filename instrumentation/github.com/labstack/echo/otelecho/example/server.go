// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Example exemplifies the otelecho package.
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
)

var tracer = otel.Tracer("echo-server")

func main() {
	tp, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	r := echo.New()
	r.Use(otelecho.Middleware("my-server"))

	r.GET("/users/:id", func(c echo.Context) error {
		id := c.Param("id")
		name := getUser(c.Request().Context(), id)
		return c.JSON(http.StatusOK, struct {
			ID   string
			Name string
		}{
			ID:   id,
			Name: name,
		})
	})
	_ = r.Start(":8080")
}

func initTracer() (*sdktrace.TracerProvider, error) { _ = "STUB: not implemented"; return nil, nil }

func getUser(ctx context.Context, id string) string { _ = "STUB: not implemented"; return "" }

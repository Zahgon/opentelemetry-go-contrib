// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package autoexport // import "go.opentelemetry.io/contrib/exporters/autoexport"

import (
	"context"
)

type signal[T any] struct {
	envKey   string
	registry *registry[T]
}

func newSignal[T any](envKey string) signal[T] { _ = "STUB: not implemented"; return nil }

func (s signal[T]) create(ctx context.Context, opts ...option[T]) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

type config[T any] struct {
	fallbackFactory func(ctx context.Context) (T, error)
}

type option[T any] interface {
	apply(cfg *config[T])
}

type optionFunc[T any] func(cfg *config[T])

//lint:ignore U1000 https://github.com/dominikh/go-tools/issues/1440
func (fn optionFunc[T]) apply(cfg *config[T]) { _ = "STUB: not implemented"; return }

func withFallbackFactory[T any](fallbackFactory func(ctx context.Context) (T, error)) option[T] {
	_ = "STUB: not implemented"
	return nil
}

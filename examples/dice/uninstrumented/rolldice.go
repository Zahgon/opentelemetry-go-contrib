// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"net/http"
)

func rolldice(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

//nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand) is ignored as this is not security-sensitive.

//nolint:gosec // G706: In a real production environment, you should sanitize the player name.

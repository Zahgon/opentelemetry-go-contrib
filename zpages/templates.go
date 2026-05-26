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
//

package zpages // import "go.opentelemetry.io/contrib/zpages"

import (
	"html/template"
)

var (
	templateFunctions = template.FuncMap{
		"even":    even,
		"spanRow": spanRowFormatter,
	}
	headerTemplate       = parseTemplate("header")
	summaryTableTemplate = parseTemplate("summary")
	tracesTableTemplate  = parseTemplate("traces")
	footerTemplate       = parseTemplate("footer")
)

// headerData contains data for the header template.
type headerData struct {
	Title string
}

func parseTemplate(name string) *template.Template { _ = "STUB: not implemented"; return nil }

//nolint:revive  // Called during initialization.

//nolint:revive  // Called during initialization.

//nolint:revive  // Called during initialization.

func spanRowFormatter(r spanRow) template.HTML {
	_ = "STUB: not implemented"
	return *new(template.HTML)
}

//nolint:gosec // G203: None of the dynamic attributes (TraceID/SpanID) can
// contain characters that need escaping so this lint issue is a false
// positive.

func even(x int) bool { _ = "STUB: not implemented"; return false }

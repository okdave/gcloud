// Copyright 2015 Google Inc. All Rights Reserved.
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

// Package reqtrace contains a very simple request tracing framework.
package reqtrace

import "golang.org/x/net/context"

// A function that must be called exactly once to report the outcome of an
// operation represented by a span.
type ReportFunc func(error)

// Return a context descending from the supplied parent that contains the
// smarts necessary to be used with the other functions in this package. If ctx
// is already the result of calling Trace, do nothing.
//
// This function starts a root span. The returned report function must be
// called when the overall operation completes.
func Trace(
	parent context.Context,
	desc string) (ctx context.Context, report ReportFunc)

// If ctx is the result of calling Trace, begin a span in the trace with the
// supplied description and return a report function that must be called to
// report the outcome of the span. Otherwise return a function that does
// nothing.
func Start(
	ctx context.Context,
	desc string) (report ReportFunc)

// Call Start, then return a function that reports the value of *err at the
// time it is invoked. Intended to be used with defer at the start of a
// function with a named error return value:
//
//     func DoSomething(ctx context.Context) (err error) {
//       defer reqtrace.StartWithError(ctx, &err, "DoSomething")
//       [...]
//     }
//
func StartWithError(
	ctx context.Context,
	err *error,
	desc string) (f func())

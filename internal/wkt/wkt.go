// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package wkt contains the list of the Google Well-Known Types as well
// as the Golang package mappings for the generated code for
// github.com/golang/protobuf and github.com/gogo/protobuf.
//
// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf
package wkt

var (
	// Filenames contains the Google Well-Known Types filenames.
	Filenames = map[string]struct{}{
		"google/protobuf/any.proto":             {},
		"google/protobuf/api.proto":             {},
		"google/protobuf/compiler/plugin.proto": {},
		"google/protobuf/descriptor.proto":      {},
		"google/protobuf/duration.proto":        {},
		"google/protobuf/empty.proto":           {},
		"google/protobuf/field_mask.proto":      {},
		"google/protobuf/source_context.proto":  {},
		"google/protobuf/struct.proto":          {},
		"google/protobuf/timestamp.proto":       {},
		"google/protobuf/type.proto":            {},
		"google/protobuf/wrappers.proto":        {},
	}

	// FilenameToGoModifierMap is a map from filename to package for google.golang.org/protobuf.
	FilenameToGoModifierMap = map[string]string{
		"google/protobuf/any.proto":             "google.golang.org/protobuf/types/known/anypb",
		"google/protobuf/api.proto":             "google.golang.org/protobuf/types/known/apipb",
		"google/protobuf/compiler/plugin.proto": "google.golang.org/protobuf/types/pluginpb",
		"google/protobuf/descriptor.proto":      "google.golang.org/protobuf/types/descriptorpb",
		"google/protobuf/duration.proto":        "google.golang.org/protobuf/types/known/durationpb",
		"google/protobuf/empty.proto":           "google.golang.org/protobuf/types/known/emptypb",
		"google/protobuf/field_mask.proto":      "google.golang.org/protobuf/types/known/field_maskpb",
		"google/protobuf/source_context.proto":  "google.golang.org/protobuf/types/known/source_contextpb",
		"google/protobuf/struct.proto":          "google.golang.org/protobuf/types/known/structpb",
		"google/protobuf/timestamp.proto":       "google.golang.org/protobuf/types/known/timestamppb",
		"google/protobuf/type.proto":            "google.golang.org/protobuf/types/known/ptypepb",
		"google/protobuf/wrappers.proto":        "google.golang.org/protobuf/types/known/wrapperspb",
	}

	// FilenameToGogoModifierMap is a map from filename to package for github.com/gogo/protobuf.
	FilenameToGogoModifierMap = map[string]string{
		"google/protobuf/any.proto":             "github.com/gogo/protobuf/types",
		"google/protobuf/api.proto":             "google.golang.org/genproto/protobuf/api",
		"google/protobuf/compiler/plugin.proto": "github.com/gogo/protobuf/protoc-gen-gogo/plugin",
		"google/protobuf/descriptor.proto":      "github.com/gogo/protobuf/protoc-gen-gogo/descriptor",
		"google/protobuf/duration.proto":        "github.com/gogo/protobuf/types",
		"google/protobuf/empty.proto":           "github.com/gogo/protobuf/types",
		"google/protobuf/field_mask.proto":      "github.com/gogo/protobuf/types",
		"google/protobuf/source_context.proto":  "google.golang.org/genproto/protobuf/source_context",
		"google/protobuf/struct.proto":          "github.com/gogo/protobuf/types",
		"google/protobuf/timestamp.proto":       "github.com/gogo/protobuf/types",
		"google/protobuf/type.proto":            "google.golang.org/genproto/protobuf/ptype",
		"google/protobuf/wrappers.proto":        "github.com/gogo/protobuf/types",
	}
)

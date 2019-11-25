// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aip0158

import (
	"fmt"

	"github.com/googleapis/api-linter/lint"
	"github.com/googleapis/api-linter/locations"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/builder"
)

var requestPaginationPageToken = &descrule.MessageRule{
	RuleName:   lint.NewRuleName(158, "request-page-token-field"),
	OnlyIf: isPaginatedRequestMessage,
	LintMessage: func(m *desc.MessageDescriptor) (problems []lint.Problem) {
		// Rule check: Establish that a page_size field is present.
		pageToken := m.FindFieldByName("page_token")
		if pageToken == nil {
			return []lint.Problem{{
				Message:    fmt.Sprintf("Message %q has no `page_token` field.", m.GetName()),
				Descriptor: m,
			}}
		}

		// Rule check: Ensure that the name page_size is the correct type.
		if pageToken.GetType() != builder.FieldTypeString().GetType() {
			return []lint.Problem{{
				Message:    "`page_token` field on List RPCs should be a string",
				Suggestion: "string",
				Descriptor: pageToken,
				Location:   locations.FieldType(pageToken),
			}}
		}

		return nil
	},
}

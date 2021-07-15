//
// Author:: Salim Afiune Maya (<afiune@lacework.net>)
// Copyright:: Copyright 2020, Lacework Inc.
// License:: Apache License, Version 2.0
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

package api

import "github.com/pkg/errors"

type QueryValidateResponse struct {
	Data    []map[string]interface{} `json:"data"`
	Ok      bool                     `json:"ok"`
	Message string                   `json:"message"`
}

func (svc *QueryService) Validate(queryText string) (
	response QueryValidateResponse,
	err error,
) {
	if queryText == "" {
		err = errors.New("query text must be provided")
		return
	}
	query := map[string]string{"query_text": queryText}
	err = svc.client.RequestEncoderDecoder("POST", apiLQLCompile, query, &response)
	return
}
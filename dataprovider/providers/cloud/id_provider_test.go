// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cloud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_k8sIdProvider_GetId(t *testing.T) {
	tests := []struct {
		name     string
		want     string
		id       string
		resource string
	}{
		{
			name:     "should return the raw id",
			want:     "metadata_id",
			id:       "metadata_id",
			resource: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewIdProvider()
			data := p.GetId(tt.resource, tt.id)
			assert.Equal(t, tt.want, data)
		})
	}
}

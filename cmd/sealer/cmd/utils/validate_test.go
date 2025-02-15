// Copyright © 2021 Alibaba Group Holding Ltd.
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

package utils

import (
	"testing"
)

func Test_validateIPStr(t *testing.T) {
	tests := []struct {
		name     string
		inputStr string
		wantErr  bool
	}{
		{
			"test empty string",
			"",
			true,
		},
		{
			"single IP",
			"1.1.1.1",
			false,
		},
		{
			"IP list format",
			"1.1.1.1,2.2.2.2,3.3.3.3",
			false,
		},
		{
			"IP range format",
			"1.1.1.1-1.1.1.255",
			false,
		},
		{
			"invalid IP list format 1",
			"1.1.1.1,2.2.2",
			true,
		},
		{
			"invalid IP list format 2",
			"1.1.1.1,",
			true,
		},
		{
			"invalid IP list format 3",
			"1.1.1.1,2.2.2.345",
			true,
		},
		{
			"invalid IP list format 4",
			",",
			true,
		},
		{
			"invalid IP range format 1",
			"1.1.1.1-1.1.1.3-1.1.1.5",
			true,
		},
		{
			"invalid IP range format 2",
			"1.1.1.1-",
			true,
		},
		{
			"invalid IP range format 3",
			"-",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateIPStr(tt.inputStr); err != nil {
				if tt.wantErr != true {
					t.Errorf("test name(%s) does not want error, but got non-nil error(%v)", tt.name, err)
				}
			} else if tt.wantErr == true {
				t.Errorf("test name(%s) wants error, but got nil error", tt.name)
			}
		})
	}
}

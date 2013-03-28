// Copyright 2013 Tumblr, Inc.
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

// Importing package worker has the side effect of turning your program into a circuit worker executable
package worker

import (
	_ "circuit/load"
	_ "circuit/kit/debug/kill"
)

func init() {
	// After package load installs and activates all circuit-related logic,
	// this function blocks forever, never allowing execution of main.
	<-(chan struct{})(nil)
}
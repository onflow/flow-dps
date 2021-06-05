// Copyright 2021 Alvalor S.A.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package storage

// TODO: Reorder the prefixes so first has 1 and last 2, before we tag the first
// release candidate:
// => https://github.com/optakt/flow-dps/issues/134

const (
	prefixFirst   = 1
	prefixLast    = 2
	prefixHeader  = 3
	prefixCommit  = 4
	prefixEvents  = 5
	prefixPayload = 6
)
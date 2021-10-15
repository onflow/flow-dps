// Copyright 2021 Optakt Labs OÜ
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

package zbor

// eventsDictionary is a byte slice that contains the result of running the Zstandard training mode
// on the events of the DPS index. This allows zstandard to achieve a better compression ratio, specifically for
// small data.
// See http://facebook.github.io/zstd/#small-data
// See https://github.com/facebook/zstd/blob/master/doc/zstd_compression_format.md#dictionary-format
var eventsDictionary = []byte{
	55, 164, 48, 236, 177, 199, 29, 56, 87, 16, 80, 150, 181, 9, 224, 65, 84, 225, 230, 221, 133, 63, 75, 3, 72, 177, 232, 55, 248, 23, 195, 240, 53, 82, 10, 91, 130, 175, 162, 245, 136, 222, 3, 248, 46, 242, 160, 63, 129, 72, 72, 91, 2, 173, 123, 69, 182, 114, 176, 61, 165, 76, 70, 23, 86, 122, 138, 151, 213, 32, 13, 135, 81, 84, 116, 236, 238, 238, 157, 210, 121, 142, 18, 22, 37, 0, 62, 9, 59, 81, 240, 16, 41, 225, 14, 9, 83, 3, 0, 0, 10, 143, 141, 200, 118, 91, 0, 0, 4, 192, 128, 130, 4, 0, 4, 19, 8, 5, 36, 13, 14, 12, 36, 18, 2, 19, 6, 20, 25, 2, 195, 97, 64, 216, 65, 101, 26, 97, 16, 4, 68, 64, 18, 198, 56, 160, 36, 86, 0, 0, 0, 84, 205, 110, 118, 161, 28, 70, 81, 48, 128, 33, 6, 56, 133, 20, 112, 160, 9, 204, 0, 0, 0, 1, 0, 0, 0, 4, 0, 0, 0, 8, 0, 0, 0, 25, 156, 133, 149, 108, 213, 112, 177, 221, 241, 162, 246, 101, 53, 97, 56, 98, 55, 102, 50, 51, 101, 56, 98, 53, 52, 56, 102, 46, 70, 108, 111, 119, 70, 101, 101, 115, 46, 84, 111, 107, 101, 110, 115, 68, 101, 112, 111, 115, 105, 116, 101, 100, 103, 80, 97, 121, 108, 111, 97, 100, 88, 117, 101, 34, 58, 34, 48, 34, 125, 125, 44, 123, 34, 110, 97, 109, 101, 34, 58, 34, 116, 111, 107, 101, 110, 115, 82, 101, 113, 117, 101, 115, 116, 101, 100, 84, 111, 85, 110, 115, 116, 97, 107, 101, 34, 44, 34, 118, 97, 108, 117, 116, 121, 112, 101, 34, 58, 34, 65, 114, 114, 97, 121, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 91, 93, 125, 125, 44, 123, 34, 110, 97, 109, 101, 34, 58, 34, 100, 101, 108, 101, 103, 97, 116, 111, 114, 73, 68, 67, 111, 117, 34, 110, 101, 116, 119, 111, 114, 107, 105, 110, 103, 65, 100, 100, 114, 101, 115, 115, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 123, 34, 116, 121, 112, 101, 34, 58, 34, 83, 116, 114, 105, 110, 103, 34, 44, 34, 118, 97, 108, 117, 48, 34, 125, 125, 44, 123, 34, 110, 97, 109, 101, 34, 58, 34, 116, 111, 34, 44, 34, 118, 97, 108, 117, 101, 34, 58, 123, 34, 116, 121, 112, 101, 34, 58, 34, 79, 112, 116, 105, 111, 110, 97, 108, 34, 44, 34, 118, 97, 108, 117, 100, 54, 101, 48, 53, 56, 54, 98, 48, 97, 50, 48, 99, 55, 34, 125, 125, 125, 93, 125, 125, 10, 106, 69, 118, 101, 110, 116, 73, 110, 100, 101, 120, 0, 109, 84, 114, 97, 110, 115, 97, 99, 116, 105, 111, 110, 73, 68, 88, 32, 121, 112, 101, 34, 58, 34, 69, 118, 101, 110, 116, 34, 44, 34, 118,
}

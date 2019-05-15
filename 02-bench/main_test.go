// Copyright 2019 Yoshi Yamaguchi
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

package main

import (
	"testing"
)

func benchmarkfindAllLines(q string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		findAllLines(q, dataFiles)
	}
}

func benchmarkfindLines(q string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, f := range dataFiles {
			findLines(q, f)
		}
	}
}

func BenchmarkFindAllLines1(b *testing.B) {
	benchmarkfindAllLines("健康", b)
}

func BenchmarkFindAllLines2(b *testing.B) {
	benchmarkfindAllLines("タイ", b)
}

func BenchmarkFindAllLines3(b *testing.B) {
	benchmarkfindAllLines("当時[0-9]+歳", b)
}

func BenchmarkFindLines1(b *testing.B) {
	benchmarkfindLines("健康", b)
}

func BenchmarkFindLines2(b *testing.B) {
	benchmarkfindLines("タイ", b)
}

func BenchmarkFindLines3(b *testing.B) {
	benchmarkfindLines("当時[0-9]+歳", b)
}

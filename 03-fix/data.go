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
	"log"
	"os"
	"path/filepath"
)

// dataFiles is the result of getDataFiles in init() and
// is used as const in this pacakge.
var dataFiles []string

func init() {
	var err error
	if dataFiles, err = getDataFiles(); err != nil {
		log.Fatalf("error on getting data files: %v", err)
	}
}

// getDataFiles returns the slice of filenames of sample data files.
func getDataFiles() ([]string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	parent := filepath.Dir(cwd)
	dataDir := filepath.Join(parent, "data")
	dir, err := os.Open(dataDir)
	if err != nil {
		return nil, err
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		return nil, err
	}
	files := make([]string, 0, len(fis))
	for _, fi := range fis {
		if fi.IsDir() {
			continue
		}
		files = append(files, filepath.Join(dataDir, fi.Name()))
	}
	return files, nil
}

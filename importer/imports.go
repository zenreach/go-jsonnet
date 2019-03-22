/*
Copyright 2017 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package jsonnet

// An Importer imports data from a path.
// TODO(sbarzowski) caching of errors (may require breaking changes)
type Importer interface {
	// Import fetches data from a given path. It may be relative
	// to the file where we do the import. What "relative path"
	// means depends on the importer.
	//
	// It is required that:
	// a) for given (importedFrom, importedPath) the same
	//    (contents, foundAt) are returned on subsequent calls.
	// b) for given foundAt, the contents are always the same
	//
	// It is recommended that if there are multiple locations that
	// need to be probed (e.g. relative + multiple library paths)
	// then all results of all attempts will be cached separately,
	// both nonexistence and contents of existing ones.
	// FileImporter may serve as an example.
	Import(importedFrom, importedPath string) (contents Contents, foundAt string, err error)
}

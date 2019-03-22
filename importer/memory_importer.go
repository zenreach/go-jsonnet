package importer

import "fmt"

// MemoryImporter "imports" data from an in-memory map.
type MemoryImporter struct {
	Data map[string]Contents
}

// Import fetches data from a map entry.
// All paths are treated as absolute keys.
func (importer *MemoryImporter) Import(importedFrom, importedPath string) (contents Contents, foundAt string, err error) {
	if content, ok := importer.Data[importedPath]; ok {
		return content, importedPath, nil
	}
	return Contents{}, "", fmt.Errorf("import not available %v", importedPath)
}

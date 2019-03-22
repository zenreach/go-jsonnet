package importer

import "fmt"

// Cache represents a cache of imported data.
//
// While the user-defined Importer implementations
// are required to cache file content, this cache
// is an additional layer of optimization that caches values
// (i.e. the result of executing the file content).
// It also verifies that the content pointer is the same for two foundAt values.
type Cache struct {
	foundAtVerification map[string]Contents
	codeCache           map[string]potentialValue
	importer            Importer
}

// MakeCache creates an Cache using an Importer.
func MakeCache(importer Importer) *Cache {
	return &Cache{
		importer:            importer,
		foundAtVerification: make(map[string]Contents),
		codeCache:           make(map[string]potentialValue),
	}
}

func (cache *Cache) importData(importedFrom, importedPath string) (contents Contents, foundAt string, err error) {
	contents, foundAt, err = cache.importer.Import(importedFrom, importedPath)
	if err != nil {
		return Contents{}, "", err
	}
	if cached, importedBefore := cache.foundAtVerification[foundAt]; importedBefore {
		if cached != contents {
			panic(fmt.Sprintf("importer problem: a different instance of Contents returned when importing %#v again", foundAt))
		}
	} else {
		cache.foundAtVerification[foundAt] = contents
	}
	return
}

// ImportString imports a string, caches it and then returns it.
func (cache *Cache) ImportString(importedFrom, importedPath string, i *interpreter, trace TraceElement) (*valueString, error) {
	data, _, err := cache.importData(importedFrom, importedPath)
	if err != nil {
		return nil, i.Error(err.Error(), trace)
	}
	return makeValueString(data.String()), nil
}

// ImportCode imports code from a path.
func (cache *Cache) ImportCode(importedFrom, importedPath string, i *interpreter, trace TraceElement) (value, error) {
	contents, foundAt, err := cache.importData(importedFrom, importedPath)
	if err != nil {
		return nil, i.Error(err.Error(), trace)
	}
	var pv potentialValue
	if cachedPV, isCached := cache.codeCache[foundAt]; !isCached {
		// File hasn't been parsed and analyzed before, update the cache record.
		pv = codeToPV(i, foundAt, contents.String())
		cache.codeCache[foundAt] = pv
	} else {
		pv = cachedPV
	}
	return i.evaluatePV(pv, trace)
}

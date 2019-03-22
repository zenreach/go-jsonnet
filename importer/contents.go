package importer

// Contents is a representation of imported data. It is a simple
// string wrapper, which makes it easier to enforce the caching policy.
type Contents struct {
	data *string
}

func (c Contents) String() string {
	return *c.data
}

// MakeContents creates Contents from a string.
func MakeContents(s string) Contents {
	return Contents{
		data: &s,
	}
}

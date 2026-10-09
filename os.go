package udetect

// OS base information structure
type OS struct {
	ID        uint   `json:"id"`            // Internal system ID (catalog root)
	VersionID uint   `json:"vid,omitempty"` // Catalog version row ID
	Name      string `json:"name,omitempty"`
	Version   string `json:"ver,omitempty"`
}

// OSDefault value
var OSDefault OS

// Clone a new OS
func (o *OS) Clone() *OS {
	return &OS{
		ID:        o.ID,
		VersionID: o.VersionID,
		Name:      o.Name,
		Version:   o.Version,
	}
}

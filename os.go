package udetect

// OS base information structure
type OS struct {
	ID      uint   `json:"id"` // Internal system ID
	Name    string `json:"name,omitempty"`
	Version string `json:"ver,omitempty"`
}

// OSDefault value
var OSDefault OS

//go:build !linux

package find

import "errors"

var ErrorFSType error = errors.New("filesystem type")

// FsTypeMap provides filesystem type names (stub for non-Linux platforms)
var FsTypeMap = map[int64]string{}

type mountPredicate struct {
	path   []string
	fsType []string
}

// Match always returns false on non-Linux platforms as filesystem type detection is not supported
func (p *mountPredicate) Match(_ string, path string) (bool, error) {
	return false, PredicateError{errType: ErrorFSType, errMessage: "filesystem type matching not supported on this platform"}
}

// getFileSystemType returns empty on non-Linux platforms
func getFileSystemType(paths ...string) []string {
	return nil
}

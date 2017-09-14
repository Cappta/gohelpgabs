package gohelpgabs

import "github.com/Jeffail/gabs"

// Container encapsulates gabs.Container
type Container struct {
	*gabs.Container
}

// New - Create a new gohelpgabs JSON object.
func New() *Container {
	return &Container{gabs.New()}
}

// ParseJSON parses a string into a representation of the parsed JSON in gabs and returns our container
func ParseJSON(sample []byte) (container *Container, err error) {
	gabsContainer, err := gabs.ParseJSON(sample)
	if err != nil {
		return
	}

	container = &Container{gabsContainer}
	return
}

// GetMissingPaths returns a list of specified paths not found in the specified container
func (container *Container) GetMissingPaths(paths ...string) (missingPaths []string) {
	for _, path := range paths {
		if container.ExistsP(path) == false {
			missingPaths = append(missingPaths, path)
		}
	}
	return
}

// SetValueIfPathExists sets the value to the specified path if it exists
func (container *Container) SetValueIfPathExists(path string, value interface{}) {
	if container.ExistsP(path) {
		container.SetP(value, path)
	}
}

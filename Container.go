package gohelpgabs

import (
	"fmt"
	"log"
	"strings"

	"github.com/Jeffail/gabs"
)

var (
	defaultErrorPath = "Errors"
)

// Container encapsulates gabs.Container
type Container struct {
	*gabs.Container
	errorPath string
}

// New - Create a new gohelpgabs JSON object.
func New() *Container {
	return &Container{gabs.New(), defaultErrorPath}
}

// ParseJSON parses a string into a representation of the parsed JSON in gabs and returns our container
func ParseJSON(sample []byte) (container *Container, err error) {
	var gabsContainer *gabs.Container
	if gabsContainer, err = gabs.ParseJSON(sample); err == nil {
		container = &Container{gabsContainer, defaultErrorPath}
	}
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

// PopPath will search the Path for the value then delete it if it exists
func (container *Container) PopPath(path string) *Container {
	popped := container.Path(path)
	if popped == nil {
		return nil
	}
	container.DeleteP(path)
	return &Container{popped, container.errorPath}
}

// Search - Attempt to find and return an object within the JSON structure by specifying the
// hierarchy of field names to locate the target. If the search encounters an array and has not
// reached the end target then it will iterate each object of the array for the target and return
// all of the results in a JSON array.
func (container *Container) Search(hierarchy ...string) *Container {
	gabsContainer := container.Container.Search(hierarchy...)
	if gabsContainer == nil {
		return nil
	}
	return &Container{gabsContainer, container.errorPath}
}

// ArrayAppendOrCreate - Append a value onto a JSON array or create one with the provided value
// if the array does not yet exists.
func (container *Container) ArrayAppendOrCreate(value interface{}, path ...string) (err error) {
	if err = container.ArrayAppend(value, path...); err == gabs.ErrNotArray {
		if _, err = container.Array(path...); err != nil {
			return
		}
		err = container.ArrayAppend(value, path...)
	}
	return
}

// ArrayAppendOrCreateP - Append a value onto a JSON array using a dot notation JSON path or
// create one with the provided value if the array does not yet exists.
func (container *Container) ArrayAppendOrCreateP(value interface{}, path string) error {
	return container.ArrayAppendOrCreate(value, strings.Split(path, ".")...)
}

// LogAndSetError - Formats the message, log it and appends to the container's error path
func (container *Container) LogAndSetError(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	log.Println(message)
	container.ArrayAppendOrCreateP(message, container.errorPath)
}

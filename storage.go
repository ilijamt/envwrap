package envwrap

import (
	"errors"
	"os"
	"strings"
	"sync"
)

var (
	// ErrEntryAlreadyExists is returned when an entry already exists
	ErrEntryAlreadyExists = errors.New("ErrEntryAlreadyExists")
	// ErrEntryDoesNotExists is returned when an entry doesn't exists
	ErrEntryDoesNotExists = errors.New("ErrEntryDoesNotExists")
)

// Storage is a container used to store the environment variables that we override
type Storage struct {
	env map[string]string

	mu sync.Mutex
}

// NewStorage creates a new environment storage instance, only used for debugging purposes, so we can test various combination of environments from the single environment
func NewStorage() *Storage {
	return &Storage{env: make(map[string]string)}
}

// NewCleanStorage removes all env variables and returns a new environment storage instance, only used for debugging purposes, so we can test various combination of environments from the single environment
func NewCleanStorage() *Storage {
	storage := &Storage{env: make(map[string]string)}
	// fetch all the envs and put in the storage
	for _, s := range os.Environ() {
		parts := strings.SplitN(s, "=", 2)
		k, v := parts[0], parts[1]
		_ = storage.Store(k, v)
	}
	os.Clearenv()
	return storage
}

// List returns the list of entries from environment storage
func (e *Storage) List() map[string]string {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.env
}

// Store stores a value for an entry in environment storage
func (e *Storage) Store(envar string, value string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, ok := e.env[envar]; ok {
		return ErrEntryAlreadyExists
	}

	if enval := os.Getenv(envar); enval != "" {
		e.env[envar] = enval
		_ = os.Setenv(envar, value)
	} else {
		e.env[envar] = ""
		_ = os.Setenv(envar, value)
	}

	return nil
}

// Release releases all values for an entry and returns an error is entry doesn't exists
func (e *Storage) Release(envar string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if val, ok := e.env[envar]; ok {
		if val == "" {
			_ = os.Unsetenv(envar)
		} else {
			_ = os.Setenv(envar, val)
		}
		delete(e.env, envar)
		return nil
	}

	return ErrEntryDoesNotExists

}

// ReleaseAll releases all environment entries we have stored, restoring the environment to the original state as before the call was made
func (e *Storage) ReleaseAll() (err error) {
	for envar := range e.env {
		_ = e.Release(envar)
	}
	return err
}

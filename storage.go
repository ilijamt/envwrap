package envwrap

import (
	"os"
	"strings"
	"sync"
	"testing"
)

type KV struct {
	Key, Value string
}

type Storage interface {
	Setenv(kv ...KV)
}

// storage is a container used to store the environment variables that we override
type storage struct {
	env      map[string]string
	mu       *sync.Mutex
	t        *testing.T
	cleanEnv bool
}

func (s *storage) init() {
	s.env = make(map[string]string)
	s.mu = new(sync.Mutex)

	if s.cleanEnv {
		s.collectEnv()
		os.Clearenv()
		s.t.Cleanup(s.restoreEnv)
	}
}

func (s *storage) collectEnv() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range os.Environ() {
		p := strings.SplitN(v, "=", 2)
		s.env[p[0]] = p[1]
	}
}

func (s *storage) restoreEnv() {
	s.mu.Lock()
	defer s.mu.Unlock()
	for k, v := range s.env {
		_ = os.Setenv(k, v)
	}
}

func (s *storage) Setenv(kv ...KV) {
	for _, v := range kv {
		key := v.Key
		value := v.Value
		prevValue, ok := os.LookupEnv(key)
		if err := os.Setenv(key, value); err != nil {
			s.t.Fatalf("cannot set environment variable: %v", err)
		}
		if ok {
			s.t.Cleanup(func() { _ = os.Setenv(key, prevValue) })
		} else {
			s.t.Cleanup(func() { _ = os.Unsetenv(key) })
		}
	}
}

func New(t *testing.T) Storage {
	t.Helper()
	s := &storage{t: t}
	s.init()
	return s
}

func NewClean(t *testing.T) Storage {
	t.Helper()
	s := &storage{t: t, cleanEnv: true}
	s.init()
	return s
}

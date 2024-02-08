package envwrap_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ilijamt/envwrap"
)

func TestNewCleanStorage(t *testing.T) {
	envs := os.Environ()
	ncs := envwrap.NewCleanStorage()
	assert.EqualValues(t, len(envs), len(ncs.List()))
	assert.Empty(t, os.Environ())
	assert.NoError(t, ncs.ReleaseAll())
	assert.EqualValues(t, len(envs), len(os.Environ()))
	assert.Empty(t, ncs.List())
}

func TestEnvNotDefinedMap(t *testing.T) {
	nes1 := envwrap.NewStorage()
	defer func() {
		assert.NoError(t, nes1.ReleaseAll())
		assert.Error(t, nes1.Release("test_nes_handler"))
		assert.Empty(t, os.Getenv("test_nes_handler"))
	}()
	assert.NoError(t, nes1.Store("test_nes_handler", "test"))
	nes2 := envwrap.NewStorage()
	assert.Error(t, nes2.Release("test_nes_handler"))
	assert.NotNil(t, nes1.List())
	assert.NoError(t, nes2.ReleaseAll())

}
func TestEnvHandler(t *testing.T) {

	env := envwrap.NewStorage()

	assert.EqualValues(t, envwrap.ErrEntryDoesNotExists, env.Release("doesntexist"))

	// set a non existing one
	assert.Empty(t, os.Getenv("testenvhandler"))
	assert.NoError(t, env.Store("testenvhandler", "yes1"))
	assert.EqualValues(t, "yes1", os.Getenv("testenvhandler"))
	assert.EqualValues(t, envwrap.ErrEntryAlreadyExists, env.Store("testenvhandler", "yes2"))
	assert.EqualValues(t, "yes1", os.Getenv("testenvhandler"))
	assert.NoError(t, env.Release("testenvhandler"))
	assert.Empty(t, os.Getenv("testenvhandler"))
	assert.Error(t, env.Release("testenvhandler"))

	// a env already exists
	assert.Empty(t, os.Getenv("testenvhandler"))
	os.Setenv("testenvhandler", "test")
	assert.EqualValues(t, "test", os.Getenv("testenvhandler"))
	assert.NoError(t, env.Store("testenvhandler", "yes1"))
	assert.EqualValues(t, "yes1", os.Getenv("testenvhandler"))
	assert.NoError(t, env.Release("testenvhandler"))
	assert.EqualValues(t, "test", os.Getenv("testenvhandler"))

	assert.NoError(t, env.ReleaseAll())
}

func ExampleNewStorage() {
	env := envwrap.NewStorage()
	oldVariable := os.Getenv("A_VARIABLE")
	_ = env.Store("A_VARIABLE", "test")
	fmt.Println(oldVariable, os.Getenv("A_VARIABLE"))
	_ = env.ReleaseAll()
	// Output: test
}

func ExampleNewCleanStorage() {
	env := envwrap.NewCleanStorage()
	fmt.Println(len(os.Environ()))
	_ = env.ReleaseAll()
	// Output: 0
}

package envwrap_test

import (
	"fmt"
	"github.com/ilijamt/envwrap"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnvNotDefinedMap(t *testing.T) {
	nes1 := envwrap.NewStorage()
	defer func() {
		nes1.ReleaseAll()
		assert.Error(t, nes1.Release("test_nes_handler"))
		assert.Empty(t, os.Getenv("test_nes_handler"))
	}()
	nes1.Store("test_nes_handler", "test")

	nes2 := envwrap.NewStorage()
	defer nes2.ReleaseAll()
	nes2.Release("test_nes_handler")

	assert.NotNil(t, nes1.List())

}
func TestEnvHandler(t *testing.T) {

	env := envwrap.NewStorage()
	defer env.ReleaseAll()

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

}

func ExampleNewStorage() {

	env := envwrap.NewStorage()
	defer env.ReleaseAll()
	oldVariable := os.Getenv("A_VARIABLE")
	env.Store("A_VARIABLE", "test")
	fmt.Println(oldVariable, os.Getenv("A_VARIABLE"))
	// Output: test
}

package envwrap_test

import (
	"os"
	"testing"

	"github.com/ilijamt/envwrap/v2"
)

func TestNew(t *testing.T) {
	// Test that New returns a non-nil Storage
	s := envwrap.New(t)
	if s == nil {
		t.Fatal("New returned nil")
	}
}

func TestNewClean(t *testing.T) {
	// Set a test environment variable
	testKey := "TEST_ENV_VAR_FOR_CLEAN"
	_ = os.Setenv(testKey, "test_value")

	// Verify the variable is set
	if val, exists := os.LookupEnv(testKey); !exists || val != "test_value" {
		t.Fatalf("Failed to set up test environment variable: %s", testKey)
	}

	// Create a clean environment
	s := envwrap.NewClean(t)
	if s == nil {
		t.Fatal("NewClean returned nil")
	}

	// Verify the environment is clean (our test variable should be gone)
	if _, exists := os.LookupEnv(testKey); exists {
		t.Errorf("NewClean did not clear environment variables")
	}

	// After the test completes, the environment should be restored
}

func TestSetenv(t *testing.T) {
	// Test setting a single environment variable
	s := envwrap.New(t)
	testKey := "TEST_SETENV_SINGLE"
	testValue := "test_value_single"

	s.Setenv(envwrap.KV{Key: testKey, Value: testValue})

	// Verify the variable was set
	if val, exists := os.LookupEnv(testKey); !exists || val != testValue {
		t.Errorf("Setenv failed to set environment variable %s", testKey)
	}

	// Test setting multiple environment variables
	testKeys := []string{"TEST_SETENV_MULTI_1", "TEST_SETENV_MULTI_2"}
	testValues := []string{"test_value_multi_1", "test_value_multi_2"}

	var kvs []envwrap.KV
	for i := range testKeys {
		kvs = append(kvs, envwrap.KV{Key: testKeys[i], Value: testValues[i]})
	}

	s.Setenv(kvs...)

	// Verify all variables were set
	for i, key := range testKeys {
		if val, exists := os.LookupEnv(key); !exists || val != testValues[i] {
			t.Errorf("Setenv failed to set environment variable %s", key)
		}
	}
}

func TestCleanupAfterTest(t *testing.T) {
	// Create a sub-test to verify cleanup works
	t.Run("SubTest", func(t *testing.T) {
		s := envwrap.New(t)
		testKey := "TEST_CLEANUP"
		testValue := "test_value_cleanup"

		// Set initial state - ensure the variable doesn't exist
		_ = os.Unsetenv(testKey)

		// Set the variable through our wrapper
		s.Setenv(envwrap.KV{Key: testKey, Value: testValue})

		// Verify it was set
		if val, exists := os.LookupEnv(testKey); !exists || val != testValue {
			t.Errorf("Setenv failed to set environment variable %s", testKey)
		}

		// When this sub-test ends, cleanup should happen automatically
	})

	// Verify the variable was cleaned up after the sub-test
	if _, exists := os.LookupEnv("TEST_CLEANUP"); exists {
		t.Error("Cleanup failed to remove environment variable")
	}
}

func TestOverrideAndRestore(t *testing.T) {
	// Test that overriding an existing variable works and is restored
	testKey := "TEST_OVERRIDE"
	originalValue := "original_value"
	newValue := "new_value"

	// Set initial value
	_ = os.Setenv(testKey, originalValue)

	t.Run("SubTest", func(t *testing.T) {
		s := envwrap.New(t)

		// Override the variable
		s.Setenv(envwrap.KV{Key: testKey, Value: newValue})

		// Verify it was overridden
		if val, exists := os.LookupEnv(testKey); !exists || val != newValue {
			t.Errorf("Setenv failed to override environment variable %s", testKey)
		}

		// When this sub-test ends, the original value should be restored
	})

	// Verify the original value was restored
	if val, exists := os.LookupEnv(testKey); !exists || val != originalValue {
		t.Errorf("Cleanup failed to restore original environment variable value for %s", testKey)
	}

	// Clean up
	_ = os.Unsetenv(testKey)
}

func TestNewCleanRestoresEnvironment(t *testing.T) {
	// Set a test environment variable
	testKey := "TEST_RESTORE_ENV"
	testValue := "test_restore_value"
	_ = os.Setenv(testKey, testValue)

	t.Run("SubTest", func(t *testing.T) {
		// Create a clean environment
		_ = envwrap.NewClean(t)

		// Verify the environment is clean
		if _, exists := os.LookupEnv(testKey); exists {
			t.Errorf("NewClean did not clear environment variables")
		}

		// When this sub-test ends, the environment should be restored
	})

	// Verify the original environment was restored
	if val, exists := os.LookupEnv(testKey); !exists || val != testValue {
		t.Errorf("NewClean failed to restore original environment")
	}

	// Clean up
	_ = os.Unsetenv(testKey)
}

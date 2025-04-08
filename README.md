# envwrap - Environment Variable Management for Go Tests
===================

[![Go Report Card](https://goreportcard.com/badge/github.com/ilijamt/envwrap)](https://goreportcard.com/report/github.com/ilijamt/envwrap)
[![Codecov](https://img.shields.io/codecov/c/gh/ilijamt/envwrap)](https://app.codecov.io/gh/ilijamt/envwrap)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ilijamt/envwrap)](go.mod)
[![GitHub](https://img.shields.io/github/license/ilijamt/envwrap)](LICENSE)
[![Release](https://img.shields.io/github/release/ilijamt/envwrap.svg)](https://github.com/ilijamt/envwrap/releases/latest)

`envwrap` is a lightweight Go package that simplifies environment variable management in tests. It provides a clean way to set, override, and restore environment variables, ensuring that your tests don't interfere with each other or leave behind unwanted state.

## Features

- **Isolated Environment Changes**: Safely modify environment variables for tests without affecting other tests
- **Automatic Cleanup**: All changes are automatically reverted when tests complete
- **Clean Environment Option**: Start with a completely clean environment for maximum isolation
- **Simple API**: Intuitive interface for setting environment variables

### Important Limitation
⚠️ Not for Parallel Tests: This package modifies global environment variables and is not suitable for parallel tests. Using it with `t.Parallel()` will lead to race conditions and unpredictable behavior.

## Installation

```bash
go get github.com/ilijamt/envwrap
```

## Usage

### Basic Usage

```go
func TestMyFunction(t *testing.T) {
    // Create a new environment wrapper
    env := envwrap.New(t)

    // Set environment variables for this test
    env.Setenv(
        envwrap.KV{Key: "API_KEY", Value: "test-api-key"},
        envwrap.KV{Key: "DEBUG", Value: "true"},
    )

    // Run your test - environment variables are set
    result := MyFunctionThatUsesEnvVars()

    // No need to clean up - envwrap handles it automatically
    // when the test completes
}
```

### Starting with a Clean Environment

If you want to start with a completely clean environment (no inherited environment variables):

```go
func TestWithCleanEnv(t *testing.T) {
    // Create a clean environment wrapper
    env := envwrap.NewClean(t)

    // Set only the environment variables you need
    env.Setenv(envwrap.KV{Key: "HOME", Value: "/tmp/fakehome"})

    // Your test runs with only the environment variables you explicitly set
    result := MyFunction()

    // Original environment is restored when test completes
}
```

### Overriding Existing Variables

When you override existing environment variables, `envwrap` automatically restores them to their original values after the test:

```go
func TestOverride(t *testing.T) {
    // HOME is typically set in the environment
    originalHome := os.Getenv("HOME")

    env := envwrap.New(t)
    env.Setenv(envwrap.KV{Key: "HOME", Value: "/tmp/testhome"})

    // HOME is now "/tmp/testhome"
    // Run your test...

    // After test completes, HOME will be restored to originalHome
}
```

## How It Works

`envwrap` uses Go's testing cleanup mechanism (`t.Cleanup()`) to register functions that will restore the environment to its original state after each test. This ensures that environment changes are isolated to individual tests and don't affect other tests or the system.

The `NewClean` function takes this a step further by saving the entire environment, clearing it completely, and then restoring everything when the test completes.

## Use Cases

- Testing code that relies on environment variables
- Isolating tests from the system environment
- Simulating different environment configurations
- Testing error handling for missing environment variables

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
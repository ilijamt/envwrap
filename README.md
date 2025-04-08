Environment Wrapper
===================

[![Go Report Card](https://goreportcard.com/badge/github.com/ilijamt/envwrap)](https://goreportcard.com/report/github.com/ilijamt/envwrap)
[![Codecov](https://img.shields.io/codecov/c/gh/ilijamt/envwrap)](https://app.codecov.io/gh/ilijamt/envwrap)
[![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ilijamt/envwrap)](go.mod)
[![GitHub](https://img.shields.io/github/license/ilijamt/envwrap)](LICENSE)
[![Release](https://img.shields.io/github/release/ilijamt/envwrap.svg)](https://github.com/ilijamt/envwrap/releases/latest)

A small and useful utility for tests so you can run your tests with multiple environments from inside the test functions.

The environment is set for the whole application while you use the wrapper, so running tests in parallel may have unexpected problems.

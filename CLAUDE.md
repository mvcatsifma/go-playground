# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this repo is

A Go playground — a flat collection of self-contained experiments and learning exercises. Each top-level directory is an independent program (its own `package main`) or test package exploring a specific Go feature, pattern, or library.

## Commands

```bash
# Run a single experiment
go run ./generics/

# Run all tests across the whole module
go test ./...

# Run tests for a single package
go test ./errors/
go test ./timeout/

# Build a single experiment
go build ./embed/
```

## Structure

Each subdirectory is standalone — a `main.go` (or a small set of `.go` files) that can be run directly with `go run ./<dir>/`. There is no shared internal library; packages do not import each other (the `timeout/` package is the only exception, which has a `system/` sub-package).

Notable experiments:
- `generics`, `generics3`, `generics5` — progressively more complex use of Go generics (type parameters, constraints, generic HTTP response handling)
- `timeout/` — context cancellation, semaphore limiting, OS signal handling, directory walking
- `job-queue/`, `integer-heap/`, `priority-queue/` — heap-based queues via `container/heap`
- `errors/` — sentinel errors, `errors.Is`/`errors.As`, custom error types
- `embed/` — `//go:embed` directives
- `viper/` — config management
- `worker_pool/`, `maxgoroutines/` — goroutine limiting patterns
- `ctrl-c/` — OS signal handling and graceful shutdown
- `unitofwork/` — Unit of Work pattern with `go-cmp`
- `godirwalk2/` — directory walking with error handling
- `options/` — functional options pattern
- `httphandler/` — func-as-`io.Writer` pattern
- `stringer/` — `go:generate` with `stringer`

## Adding a new experiment

Create a new directory with a `main.go` using `package main`. The module path is `dsen.nl/go-playground`, so sub-packages import as `dsen.nl/go-playground/<dir>/<subpkg>`.

package main

import (
	"context"
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"
	"testing/synctest"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_runTask_foundTarget(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		task := &Task{id: 1, path: "testdata"}
		testFS := fstest.MapFS{
			"testdata/sentinel-deadbeef-marker.txt": &fstest.MapFile{
				Data: []byte{},
				Mode: 644,
			},
		}

		result := runTask(ctx, testFS, task)

		assert.True(t, result.foundTarget)
	})
}

func Test_runTask_skipDir(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		task := &Task{id: 1, path: "testdata"}
		testFS := fstest.MapFS{
			"testdata/restricted/file_a.txt": &fstest.MapFile{
				Data: []byte{},
				Mode: 644,
			},
			"testdata/file_b.txt": &fstest.MapFile{
				Data: []byte{},
				Mode: 644,
			},
		}

		result := runTask(ctx, testFS, task)

		assert.Equal(t, int64(2), result.visited)
	})
}

func Test_runTask_walkErr(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		task := &Task{id: 1, path: "doesnotexist"}
		testFS := fstest.MapFS{
			"foobar": &fstest.MapFile{
				Data: []byte{},
				Mode: 644,
			},
		}

		result := runTask(ctx, testFS, task)

		var pathErr *fs.PathError
		assert.True(t, errors.As(result.err, &pathErr))
	})
}

func Test_runTask_Timeout(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()
		time.Sleep(time.Second + time.Millisecond)

		task := &Task{id: 1, path: "testdata"}
		testFS := fstest.MapFS{
			"testdata/foo.txt": &fstest.MapFile{
				Data: []byte{},
				Mode: 644,
			},
		}

		result := runTask(ctx, testFS, task)

		assert.True(t, result.timeout)
	})
}

func Test_runTask_Cancel(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		cancel()

		task := &Task{id: 1, path: "testdata"}
		testFS := fstest.MapFS{
			"testdata/foo.txt": &fstest.MapFile{
				Data: []byte{},
				Mode: 644,
			},
		}

		result := runTask(ctx, testFS, task)

		assert.True(t, result.canceled)
	})
}

// brokenFS is a test filesystem that always returns permission denied errors.
// Implements fs.FS interface for testing error handling.
type brokenFS struct{}

func (b brokenFS) Open(name string) (fs.File, error) {
	return nil, &fs.PathError{
		Op:   "open",
		Path: name,
		Err:  fs.ErrPermission,
	}
}

func Test_runTask_PermissionError(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
		defer cancel()

		task := &Task{id: 1, path: "restricted"}
		testFS := brokenFS{}

		result := runTask(ctx, testFS, task)

		// Verify we got a permission error
		assert.NotNil(t, result.err)
		var pathErr *fs.PathError
		if assert.True(t, errors.As(result.err, &pathErr)) {
			assert.ErrorIs(t, pathErr.Err, fs.ErrPermission)
		}
	})
}


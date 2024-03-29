// Copyright (c) 2023  The Go-Curses Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package path

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/djherbis/times"
)

var (
	ErrExistingFile = errors.New("path is an existing file")
)

// Abs is a convenience wrapper around filepath.Abs
func Abs(path string) (absolute string, err error) {
	absolute, err = filepath.Abs(path)
	return
}

// Clean is a convenience wrapper around filepath.Clean
func Clean(path string) (cleaned string) {
	cleaned = filepath.Clean(path)
	return
}

// Dir is a convenience wrapper around filepath.Dir
func Dir(path string) (name string) {
	name = filepath.Dir(path)
	return
}

// Walk is a convenience wrapper around filepath.Walk
func Walk(root string, fn filepath.WalkFunc) (err error) {
	err = filepath.Walk(root, fn)
	return
}

// ReadDir is a convenience wrapper around os.ReadDir
func ReadDir(path string) (paths []fs.DirEntry, err error) {
	paths, err = os.ReadDir(path)
	return
}

// ReadFile is a convenience wrapper around os.ReadFile
func ReadFile(path string) (content []byte, err error) {
	content, err = os.ReadFile(path)
	return
}

// Stat is a convenience wrapper around github.com/djherbis/times.Stat
func Stat(path string) (spec times.Timespec, err error) {
	spec, err = times.Stat(path)
	return
}

// Which is a convenience wrapper around exec.LookPath and filepath.Abs
func Which(name string) (path string) {
	var err error
	var p, abs string
	if p, err = exec.LookPath(name); err == nil {
		if abs, err = filepath.Abs(p); err == nil {
			path = abs
		}
	}
	return
}

// Pwd is a wrapper around os.Getwd
func Pwd() (pwd string) {
	pwd, _ = os.Getwd()
	return
}

// MkdirAll is a wrapper around os.MkdirAll with DefaultPathPerms (0770)
func MkdirAll(path string) (err error) {
	if IsFile(path) {
		err = ErrExistingFile
		return
	} else if IsDir(path) {
		// nop
		return
	}
	err = os.MkdirAll(path, DefaultPathPerms)
	return
}

// ChmodAll uses Walk to traverse the filesystem and change all regular file
// permissions to the DefaultFilePerms (0660) and all directory
// permissions to the DefaultPathPerms (0770).
//
// Stops on the first error
func ChmodAll(src string) error {
	return filepath.WalkDir(src, func(path string, entry fs.DirEntry, err error) error {
		if err == nil {
			if entry.IsDir() {
				err = os.Chmod(path, DefaultPathPerms)
			} else if info, ee := entry.Info(); ee == nil && info.Mode().IsRegular() {
				err = os.Chmod(path, DefaultFilePerms)
			}
		}
		return err
	})
}

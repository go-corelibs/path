// Copyright (c) 2024  The Go-Curses Authors
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
	"io/fs"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWrappers(t *testing.T) {
	Convey("(all)", t, func() {
		cwd, _ := os.Getwd()
		// Abs
		path, err := Abs("./thing/..")
		So(err, ShouldEqual, nil)
		So(path, ShouldEqual, cwd)
		// Clean
		So(Clean("/thing/.."), ShouldEqual, "/")
		// Dir
		So(Dir("/thing/file.txt"), ShouldEqual, "/thing")
		// Walk
		So(Walk(".", func(path string, info fs.FileInfo, err error) error {
			return nil
		}), ShouldEqual, nil)
		// ReadDir
		_, err = ReadDir(".")
		So(err, ShouldEqual, nil)
		// ReadFile
		_, err = ReadFile("wrappers_test.go")
		So(err, ShouldEqual, nil)
		// Stat
		_, err = Stat(".")
		So(err, ShouldEqual, nil)
		// Which
		path = Which("go")
		So(path, ShouldNotEqual, "")
	})

	Convey("Pwd", t, func() {
		cwd, err := os.Getwd()
		So(err, ShouldBeNil)
		So(cwd, ShouldNotEqual, "")
		So(Pwd(), ShouldEqual, cwd)
	})

	Convey("MkdirAll", t, func() {
		tempDir, err := os.MkdirTemp("", "corelibs-path.*.d")
		So(err, ShouldBeNil)
		So(tempDir, ShouldNotEqual, "")

		Convey("path is an existing file", func() {
			So(os.WriteFile(tempDir+"/file.txt", []byte("nope"), DefaultFilePerms), ShouldBeNil)
			So(MkdirAll(tempDir+"/file.txt"), ShouldEqual, ErrExistingFile)
		})

		Convey("path is an existing directory", func() {
			So(MkdirAll(tempDir), ShouldBeNil)
		})

		Convey("make all dirs", func() {
			So(MkdirAll(tempDir+"/one/two"), ShouldBeNil)
			So(IsDir(tempDir+"/one"), ShouldBeTrue)
			So(IsDir(tempDir+"/one/two"), ShouldBeTrue)
		})

		So(os.RemoveAll(tempDir), ShouldEqual, nil)
	})

	Convey("ChmodAll", t, func() {
		tempDir, err := os.MkdirTemp("", "corelibs-path.*.d")
		So(err, ShouldBeNil)
		So(tempDir, ShouldNotEqual, "")

		So(os.WriteFile(tempDir+"/file.txt", []byte("nope"), 0440), ShouldBeNil)
		So(IsPermission(tempDir+"/file.txt", 0440), ShouldBeTrue)

		So(os.MkdirAll(tempDir+"/one/two", 0750), ShouldBeNil)
		So(IsPermission(tempDir+"/one/two", 0750), ShouldBeTrue)

		So(ChmodAll(tempDir), ShouldBeNil)
		So(IsPermission(tempDir+"/file.txt", DefaultFilePerms), ShouldBeTrue)
		So(IsPermission(tempDir+"/one", DefaultPathPerms), ShouldBeTrue)
		So(IsPermission(tempDir+"/one/two", DefaultPathPerms), ShouldBeTrue)

		So(os.RemoveAll(tempDir), ShouldEqual, nil)
	})
}

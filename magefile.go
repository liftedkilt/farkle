//+build mage

package main

import (
	"os"

	"github.com/magefile/mage/sh"
)

var Aliases = map[string]interface{}{
	"b": Build,
}

func init() {
	// Until GOPATH is deprecated completely, we want to make sure we are always using modules
	os.Setenv("GO111MODULE", "on")
}

func Build() error {
	return sh.Run("go", "build", "-o", "bin/farkle")
}

func Test() error {
	return sh.RunV("go", "test", "./...")
}

func Tidy() error {
	return sh.RunV("go", "mod", "tidy")
}

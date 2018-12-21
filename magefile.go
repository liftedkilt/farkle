//+build mage

package main

import (
	"errors"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Aliases = map[string]interface{}{
	"b": Build,
	"t": Test,
}

type Mod mg.Namespace

func init() {
	// Until GOPATH is deprecated completely, we want to make sure we are always using modules
	os.Setenv("GO111MODULE", "on")
}

// output a farkle binary in ./bin
func Build() error {
	// We want to run from vendored deps until proxies become widespread
	return sh.Run("go", "build", "-o", "bin/farkle", "-mod", "vendor")
}

// run all tests and if all tests are passing, build
func Release() error {
	err := Test()
	if err != nil {
		return errors.New("\n\ntests are not passing. Aborting release process")
	}
	return Build()
}

// run all package tests
func Test() error {
	return sh.RunV("go", "test", "./...")
}

// remove unused deps
func (Mod) Tidy() error {
	return sh.RunV("go", "mod", "tidy")
}

// ensure all deps are vendored
func (Mod) Vendor() error {
	return sh.RunV("go", "mod", "vendor")
}

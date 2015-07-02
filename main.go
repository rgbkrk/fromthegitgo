package main

import (
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
	pkg := PkgForPath(".")

	out, err := exec.Command("git", "rev-parse", "--verify", "HEAD").Output()
	if err != nil {
		fmt.Errorf("Failed to run git: %v\n", err)
	}
	commitHash := strings.TrimSpace(string(out))

	const hashGo = `package %s

const HASH = "%s"`

	contents := fmt.Sprintf(hashGo, pkg.Name, commitHash)

	err = ioutil.WriteFile("hash.go", []byte(contents), 0644)

	if err != nil {
		panic(fmt.Sprintf("Couldn't write hash.go: %v", err))
	}

}

// PkgForPath gets a *build.Package for a given path
func PkgForPath(path string) *build.Package {
	// get pwd for relative imports
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("error getting working directory (required for relative imports): %s\n", err)
		os.Exit(1)
	}

	// read full package information
	pkg, err := build.Import(path, pwd, 0)
	if err != nil {
		fmt.Printf("error reading package: %s\n", err)
		os.Exit(1)
	}

	return pkg
}

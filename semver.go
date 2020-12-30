package main

import (
	"fmt"
	"sort"

	"github.com/Masterminds/semver"
)

// Script to test parsing semantic version
// This script will look to:
// * Parse Sem: Parse Semantic Version
// * SortSem: Sort Semantic Version
// * Check if a semver fits within a constraint
// * Work with a specific prefix ?
// GoDoc Reference: https://godoc.org/github.com/Masterminds/semver && https://blog.gopheracademy.com/advent-2015/semver/
func parsesem() {
	v, err := semver.NewVersion("1.2.3-beta.1+b345")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(v.Major())
	}
}

func sortsem() {
	raw := []string{"1.2.3", "1.0", "1.3", "2", "0.4.2"}
	vs := make([]*semver.Version, len(raw))

	// Loop through
	for i, r := range raw {
		v, err := semver.NewVersion(r)
		if err != nil {
			fmt.Errorf("Semver Error version: %s", err)
		}
		vs[i] = v
	}
	sort.Sort(semver.Collection(vs))
	fmt.Println(vs)
}

func main() {
	fmt.Println("Sorting semver")
	sortsem()
}

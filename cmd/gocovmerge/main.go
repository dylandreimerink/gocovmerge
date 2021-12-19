// gocovmerge takes the results from multiple `go test -coverprofile` runs and
// merges them into one profile
package main

import (
	"flag"
	"log"
	"os"

	"github.com/dylandreimerink/gocovmerge"
	"golang.org/x/tools/cover"
)

func main() {
	flag.Parse()

	var merged []*cover.Profile

	for _, file := range flag.Args() {
		profiles, err := cover.ParseProfiles(file)
		if err != nil {
			log.Fatalf("failed to parse profiles: %v", err)
		}
		for _, p := range profiles {
			merged = gocovmerge.AddProfile(merged, p)
		}
	}

	gocovmerge.DumpProfiles(merged, os.Stdout)
}

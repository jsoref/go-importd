package main

import "fmt"

// This is overwritten by the ldflag "-X main.version=<blah>"
// See script/build for more info.
var (
	version = "master"
	commit  = "HEAD"
	date    = "unknown"
)

func versionString() string {
	return fmt.Sprintf("%s (%s) @ %s", version, commit, date)
}

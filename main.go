package main

import (
	"flag"
	"fmt"

	"github.com/lunny/axmlParser"
)

func main() {
	// Read flag
	showVersionCode := flag.Bool("version-code", false, "Show apk version code")
	showVersionName := flag.Bool("version-name", false, "Show apk version name")
	flag.Parse()

	listener := new(axmlParser.AppNameListener)
	axmlParser.ParseApk(flag.Args()[0], listener)

	if *showVersionCode {
		fmt.Println(listener.VersionCode)
	}

	if *showVersionName {
		fmt.Println(listener.VersionName)
	}
}

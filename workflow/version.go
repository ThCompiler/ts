package main

import "fmt"

// Version - version of app.
const Version = "v0.0.1-alpha"

func main() {
	_, _ = fmt.Printf("%s\n", Version) //nolint:forbidigo // it is bin for print version
}

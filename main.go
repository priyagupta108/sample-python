//+build linux darwin windows
package main

import (
	"C"
)

//export cNoop
func cNoop() {
}

func main() {}

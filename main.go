// go:build (darwin && cgo) || linux
// +build darwin,cgo linux
package main

import (
	"C"
)

//export cNoop
func cNoop() {
}

func main() {}

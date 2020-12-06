// Package organya is a Go implementation of the Organya sound format. It can be
// used to parse Organya files and makes it possible to play them back using
// your own implementation.
//
// Usage
//
// Reading an Organya file is very simple. Just use the `Open` method and
// provide the path to the Organya file:
//	package main
//
//	import (
//		"pkg.nimblebun.works/go-organya"
//	)
//
//	func main() {
//		org, err := organya.Open("/path/to/file.org")
//		if err != nil {
//			panic(err)
//		}
//
//		// ...
//	}
//
// You can now take advantage of the methods and properties that are exposed to
// you.
package organya

const (
	// OrgTrackCount defines the number of possible Organya tracks in a file (16).
	OrgTrackCount int = 0x10

	// OrgNoChange is the byte that specifies that a current resource property
	// must not be different from the one before it.
	OrgNoChange uint8 = 0xFF
)

// +build !cgo

// Loaded on WASM
package db

import (
	"github.com/k0kubun/pp"
)

func Run () {
	// Do nothing for WASM env.
	pp.Println("Skip execution of db.Run() because ENV looks like WASM")
}
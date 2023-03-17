package main

import (
	"howett.net/plist"
	"os"
)

func mainss() {
	encoder := plist.NewEncoder(os.Stdout)
	encoder.Encode(map[string]string{"hello": "world"})
}

package main

import (
	"fmt"
	"log"

	versions "./versions"
	_ "./versions/all"
)

type FieldSystem versions.FieldSystem

func main() {
	attachfs, err = versions.GetInstalled()
	if err {
		log.Fatalf(err)
	}
	fs := attachfs()
	fmt.Println(fs.Version())
	fmt.Println(fs.Log())
}

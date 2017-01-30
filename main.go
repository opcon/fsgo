package main

import (
	"fmt"
	"log"

	versions "fs/versions"
	_ "fs/versions/all"
)

type FieldSystem versions.FieldSystem

func main() {
	fs, err := versions.Attach()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(fs.Version())
	fmt.Println(fs.Log())
}

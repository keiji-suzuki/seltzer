package main

import (
	"log"

	"github.com/keiji-suzuki/seltzer/cmd"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"os"

	"github.com/threatresponse/cloudtrailbeat/beater"
	"github.com/elastic/beats/libbeat/beat"
)

var Version = "0.0.1"
var Name = "cloudtrailbeat"

func main() {
	if err := beat.Run(Name, Version, beater.New()); err != nil {
		os.Exit(1)
	}
}

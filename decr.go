package main

import (
	"decr/lib/app"
	"os"
)

func main() {
	os.Exit(app.App.Run(os.Args, os.Stdout))
}

package main

import (
	"flag"
	"fmt"

	"github.com/vijaykiran225/playlist-generator/controller"
)

func main() {
	fmt.Println("started")
	var path string
	flag.StringVar(&path, "path", ".", "give something!")
	flag.Parse()
	controller.ManageApp(path)

}

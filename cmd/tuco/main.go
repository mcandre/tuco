// Package main implements a tool for automating Go crosscompilation.
package main

import (
	"github.com/mcandre/tuco"

	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

var flagClean = flag.Bool("clean", false, "remove artifacts")
var flagVersion = flag.Bool("version", false, "show version")
var flagHelp = flag.Bool("help", false, "show usage menu")

func usage() {
	program, err := os.Executable()

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Usage: %v [OPTION]\n", program)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	switch {
	case *flagVersion:
		fmt.Println(tuco.Version)
		os.Exit(0)
	case *flagHelp:
		usage()
		os.Exit(0)
	}

	tc, err := tuco.Load()

	if err != nil {
		log.Fatal(err)
	}

	if *flagClean {
		if err2 := tc.Clean(); err2 != nil {
			log.Fatal(err2)
		}

		os.Exit(0)
	}

	if errs := tc.Run(); len(errs) != 0 {
		sort.Slice(errs, func(i, j int) bool {
			return errs[i].Error() < errs[j].Error()
		})

		for _, err2 := range errs {
			log.Println(err2)
		}

		os.Exit(1)
	}
}

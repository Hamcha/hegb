package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hamcha/hegb"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <romfile.gb>\n", os.Args[0])
		flag.PrintDefaults()
	}

	romdata := flag.Bool("rominfo", false, "Print ROM info and exit")
	flag.Parse()

	// Must be at least one non-flag argument (ROM file)
	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	data, err := ioutil.ReadFile(flag.Arg(0))
	assert(err)

	rom := hegb.LoadROM(data)

	if *romdata {
		fmt.Println(rom)
	}
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

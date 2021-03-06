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
	usebs := flag.Bool("bootrom", true, "Use boot ROM")
	dumpcode := flag.Bool("dumpcode", false, "Dump code as its being executed")
	flag.Parse()

	// Must be at least one non-flag argument (ROM file)
	if flag.NArg() < 1 {
		flag.Usage()
		return
	}

	data, err := ioutil.ReadFile(flag.Arg(0))
	assert(err)

	if *romdata {
		header, err := hegb.GetROMHeader(data)
		assert(err)
		fmt.Println(header)
		return
	}

	rom, err := hegb.LoadROM(data)
	assert(err)

	gb := hegb.MakeGB(rom, hegb.EmulatorOptions{
		UseBootstrap: *usebs,
		DumpCode:     *dumpcode,
	})
	gb.Run()
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

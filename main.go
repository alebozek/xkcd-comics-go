package main

import (
	"flag"
	"fmt"
	"os"
	"unsafe"
	"xkcd-comics/xkcd-comics"
)

var programName string = "xkcd-comics"

func main() {
	randomComic := flag.Bool("r", false, "Gets a random comic.")
	numberComic := flag.Int("n", 0, "Gets the comic with that number if exists.")
	noDisplay := flag.Bool("no-display", false, "If true, doesn't display with feh the downloaded comic.")
	//flag.Usage() = usage
	flag.Parse()

	if len(os.Args) == 1 {
		usage()
		os.Exit(2)
	}

	var b bool
	b = *((*bool)(unsafe.Pointer(noDisplay)))

	if *randomComic == true {
		xkcd_comics.GetRandomComic(b)
	}

	if numberComic != nil {
		var i int
		i = *((*int)(unsafe.Pointer(numberComic)))
		xkcd_comics.GetComicByNumber(i, b)
	}
}
func usage() {
	msg := fmt.Sprintf("USAGE: %s [OPTIONS]\n%s gets a comic from xkcd.com and displays it to you.", programName, programName)
	fmt.Println(msg)
	flag.PrintDefaults()
}

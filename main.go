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
	numberComic := flag.Int("n", 0, "Gets the comic with that number if exists")

	//flag.Usage() = usage
	flag.Parse()

	if len(os.Args) == 1 {
		usage()
		os.Exit(2)
	}

	if *randomComic == true {
		xkcd_comics.GetRandomComic()
	}

	if numberComic != nil {
		var i int
		i = *((*int)(unsafe.Pointer(numberComic)))
		xkcd_comics.GetComicByNumber(i)
	}
}
func usage() {
	msg := fmt.Sprintf("USAGE: %s [OPTIONS]\n%s gets a comic from xkcd.com and displays it to you.", programName, programName)
	fmt.Println(msg)
	flag.PrintDefaults()
}

/*
$ ./flags a bc def
-> a bc def
$ ./flags -s / a bc def
-> a/bc/def
$ ./flags -n a bc def
-> a bc def$
$ ./flags -help

	Usage of ./flags:
	-n
		omit trailing newline
	-s string
		separator (default " ")
*/
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")

var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

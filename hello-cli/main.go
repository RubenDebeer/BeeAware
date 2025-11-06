package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	nameFlag    string
	shoutFlag   bool
	versionFlag bool
)

var version = "0.1.0"

func init() {
	flag.StringVar(&nameFlag, "name", "world", "name to greet")
	flag.BoolVar(&shoutFlag, "shout", false, "print in ALL CAPS")
	flag.BoolVar(&versionFlag, "version", false, "print version and exit")

	initShortHandCommands()
}

func initShortHandCommands() {
	flag.StringVar(&nameFlag, "n", "world", "short for -name")
	flag.BoolVar(&shoutFlag, "s", false, "short for -shout")
	flag.BoolVar(&versionFlag, "v", false, "short for -version")
}

func main() {
	flag.Parse()

	if versionFlag {
		fmt.Println(version)
		return
	}

	args := flag.Args()

	if len(args) > 0 {

		switch args[0] {
		case "help":
			usage()
			return
		default:
			fmt.Fprintf(os.Stderr, "unknown command: %q\n\n", args[0])
			usage()
			os.Exit(2)
		}
	}

	msg := fmt.Sprintf("Hello, %s!", nameFlag)

	if shoutFlag {
		msg = toUpper(msg)
	}
	fmt.Println(msg)
}

func toUpper(s string) string {

	b := []rune(s)
	for i, r := range b {
		if r >= 'a' && r <= 'z' {
			b[i] = r - ('a' - 'A')
		}
	}
	return string(b)
}

func usage() {
	fmt.Fprintf(os.Stderr, `hello-cli %s

	Usage:
	hello-cli [flags]           # default action (greet)
	hello-cli [flags] help      # show help

	Flags:
	`, version)
	flag.PrintDefaults()
}

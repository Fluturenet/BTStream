package main

import (
	"flag"
	_ "fmt"
	"os"

	_ "github.com/fluturenet/ed25519"
)

var (
	dataDir string
)

func main() {
	flag.StringVar(&dataDir, "datadir", ".", "Path where store key and temporary data")
	flag.StringVar(&dataDir, "d", ".", "Path where store key and temporary data")

	genkey := flag.Bool("genkey", false, "generates key dentity. ex -genkey")
	flag.Parse()

	if *genkey {
		generateKey()
		os.Exit(0)
	}

	gentlydie()
}

func gentlydie() {
	flag.Usage()
	os.Exit(0)
}

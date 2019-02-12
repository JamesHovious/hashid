package main

// Software to identify the different types of hashes
// A Go implimentation of https://github.com/psypanda/hashID

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/JamesHovious/hashid"
	flag "github.com/spf13/pflag"
)

var outFilePtr *string
var extended, version, mode, john *bool
var out []byte

func main() {
	lastArg := os.Args[len(os.Args)-1]
	parseArgs()
	checkVersion()
	if lastArg != *outFilePtr && len(lastArg) > 2 {
		getHash(lastArg)
	} else {
		reader := bufio.NewReader(os.Stdin)
		hash, _ := reader.ReadString('\n')
		getHash(hash)
	}
	if *outFilePtr != "" {
		writeOutputFile()
	}
	writeStdOut()
	fmt.Println(string(out[:]))
}

func checkVersion() {
	if *version {
		fmt.Println("version 0.1")
		os.Exit(0)
	}
}

func getHash(hash string) {

	hasheTypes, err := hashid.GetHashType(hash)
	newline := "\n"
	space := " "
	if err != nil {
		panic(err)
	}

	for _, hashType := range hasheTypes {
		for _, m := range hashType.Modes {
			if !*extended && m.Extended {
				break
			}
			out = append(out, m.Name...)
			if *mode {
				if m.Hashcat != nil {
					hashcatBanner := fmt.Sprintf("[HashcatMode: %d]", *m.Hashcat)
					out = append(out, space...)
					out = append(out, hashcatBanner...)
				}
			}
			if *john {
				jtrBanner := fmt.Sprintf("[JtR Format: %v]", m.John)
				out = append(out, space...)
				out = append(out, jtrBanner...)
			}
			out = append(out, newline...)
		}
	}
}

func writeOutputFile() {
	err := ioutil.WriteFile(*outFilePtr, out, 0644)
	if err != nil {
		fmt.Println("[!] Error writing file")
		fmt.Println(err)
		os.Exit(0)
	}
	fmt.Println("[+] Successfully wrote out file")
}

func writeStdOut() {

}

func parseArgs() {
	mode = flag.BoolP("mode", "m", false, "show corresponding hashcat mode in output")
	john = flag.BoolP("john", "j", false, "show corresponding JohnTheRipper format in output")
	version = flag.BoolP("version", "v", false, "show program's version number and exit")
	outFilePtr = flag.StringP("outfile", "o", "", "write output to file (default: STDOUT)")
	extended = flag.BoolP("extended", "e", false, "list all hash algorithms including salted passwords")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Identify the different types of hashes used to encrypt data and especially passwords.
hashid is a tool written in Go which supports the identification of over 220 unique hash types using regular expressions.
hashid is written as a Go library so that it can easily be inlcuded in other tools. 
It is able to identify a single hash. hashid is also capable of including the corresponding hashcat mode and/or JohnTheRipper format in its output.
Note: When identifying a hash on *nix operating systems use single quotes to prevent interpolation.

Usage:
`)
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()
}

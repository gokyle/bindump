// bindump is a simple utility to dump binary files.
package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	flDump := flag.Bool("b", false, "binary dump")
	flCDump := flag.Bool("c", false, "C-struct-style dump")
	flHDump := flag.Bool("hex", false, "hex string dump")
	flag.Parse()

	if !*flHDump && !*flCDump {
		*flDump = true
	}

	nfiles := flag.NArg()
	for i, name := range flag.Args() {
		fmt.Printf("%s:\n", name)
		if *flDump {
			dumpFile(name)
		}
		if *flCDump {
			cDumpFile(name)
		}
		if *flHDump {
			hexString(name)
		}
		if i != (nfiles-1) {
			fmt.Printf("\n")
		}
	}
}

func dumpFile(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("[!]", err.Error())
		return
	}

	for {
		var in = make([]byte, 16)
		n, err := file.Read(in)
		if err != nil {
			if err != io.EOF {
				fmt.Println("[!]", err.Error())
			}
			return
		}

		fmt.Printf("\t")
		for i := 0; i < n; i++ {
			fmt.Printf("%02x ", in[i])
			if i == 7 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}

func cDumpFile(name string) {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println("[!]", err.Error())
		return
	}

	for {
		var in = make([]byte, 8)
		n, err := file.Read(in)
		if err != nil {
			if err != io.EOF {
				fmt.Println("[!]", err.Error())
			}
			return
		}

		fmt.Printf("\t")
		for i := 0; i < n; i++ {
			fmt.Printf("0x%02x, ", in[i])
		}
		fmt.Printf("\n")
	}
}

func hexString(name string) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("[!]", err.Error())
		return
	}
	fmt.Printf("%x\n", data)
}

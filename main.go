package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var (
	byteSlice   bool
	trimSpace   bool
	packageName string
	fileMode    int
)

func init() {
	flag.BoolVar(&byteSlice, "bytes", false, "Should the variable be a byte slice rather than a string?")
	flag.StringVar(&packageName, "pkg", "main", "Name of the package for the created go file.")
	flag.IntVar(&fileMode, "mode", 0644, "Mode for the created file.")
	flag.BoolVar(&trimSpace, "trim", false, "Should trim space be called on the resulting string?")
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: filestr [source file] [dest file] [var name]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) != 3 {
		usage()
	}
	srcFile, destFile, varName := args[0], args[1], args[2]
	src, err := ioutil.ReadFile(srcFile)
	if err != nil {
		log.Fatal(err)
	}
	dest, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(fileMode))
	if err != nil {
		log.Fatal(err)
	}
	if trimSpace {
		src = bytes.TrimSpace(src)
	}
	if !byteSlice {
		_, err = fmt.Fprintf(dest, "package %s\n\nvar %s string = %s\n", packageName, varName, strconv.Quote(string(src)))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	_, err = fmt.Fprintf(dest, "package %s\n\nvar %s = []byte{", packageName, varName)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 4)
	buf[0] = '0'
	buf[1] = 'x'
	encodeBuf := make([]byte, 1)
	w := bufio.NewWriter(dest)
	for i, b := range src {
		if i > 0 {
			if err := w.WriteByte(','); err != nil {
				log.Fatal(err)
			}
		}
		encodeBuf[0] = b
		hex.Encode(buf[2:], encodeBuf)
		if _, err := w.Write(buf); err != nil {
			log.Fatal(err)
		}
	}
	if err = w.WriteByte('}'); err != nil {
		log.Fatal(err)
	}
	if err = w.WriteByte('\n'); err != nil {
		log.Fatal(err)
	}
	if err = w.Flush(); err != nil {
		log.Fatal(err)
	}
}

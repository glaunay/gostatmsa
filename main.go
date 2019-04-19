package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/docopt/docopt-go"
	"github.com/msago"
)

func main() {
	usage := `Usage:
	msago <file> [--name=<letters> --word=<subseq>]
	
Options:
	  -h --help     Show this screen.
	  --version     Show version.
	  --name=<n>  	Filter sequence based on name substring [default: None].
	  --word=<w>  	Filter sequence based on sequence substring [default: None].`

	opts, _ := docopt.ParseArgs(usage, os.Args[1:], "rc.O.1.1")
	var conf struct {
		File string
		Name string
		Word string
	}
	opts.Bind(&conf)

	msa := msago.ParseFile([]byte(conf.File))

	if conf.Name != "None" {
		subMsa := msa.MapSearch(
			func(name string, sequence string) bool {
				return strings.Contains(name, conf.Name) //"THE"
			})
		msa = subMsa
	}
	if conf.Word != "None" {
		subMsa := msa.MapSearch(
			func(name string, sequence string) bool {
				return strings.Contains(sequence, conf.Word) //"THE"
			})
		msa = subMsa
	}

	fmt.Println("matching record size", msa.Len())
	// Interator /closure for fun
	var record *msago.MsaRecord
	for next, hasNext := msa.Iterator(); hasNext; {
		record, hasNext = next()
		fmt.Println(record)
	}
}

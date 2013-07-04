package main

import (
	"bufio"
	"os"
)

type context struct {
	reader bufio.Reader
	script []*os.File
	input  []*os.File

	vars                *varTable
	safeMode            bool
	maxField, maxRecord int
}

func newContext() *context {
	con := &context{vars: &varTable{
		tab: map[string]*cell{
			"CONVFMT":  &cell{name: "CONVFMT", sval: `%.6g`},
			"FS":       &cell{name: "FS", sval: " "},
			"NF":       &cell{name: "NF"},
			"NR":       &cell{name: "NR"},
			"FNR":      &cell{name: "FNR"},
			"FILENAME": &cell{name: "FILENAME"},
			"RS":       &cell{name: "RS", sval: `\n`},
			"OFS":      &cell{name: "OFS"},
			"ORS":      &cell{name: "ORS", sval: `\n`},
			"OFMT":     &cell{name: "OFMT", sval: `%.6g`},
			"SUBSEP":   &cell{name: "SUBSEP", sval: "034", nval: 034},
			"ARGC":     &cell{name: "ARGC"},
			"ARGV":     &cell{name: "ARGV"},
			"ENVIRON":  &cell{name: "ENVIRON"},
		}},
	}
	return con
}
package main

import (
	"github.com/Makoto87/checkvarname"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(checkvarname.Analyzer) }

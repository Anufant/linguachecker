package main

import (
	"github.com/anufant/linguachecker"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(linguachecker.Analyzer)
}

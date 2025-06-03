package main

import (
	"github.com/shiron-dev/go-err113"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(err113.NewAnalyzer())
}

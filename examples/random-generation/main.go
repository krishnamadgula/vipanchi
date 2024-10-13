package main

import (
	"log/slog"
	"os"

	"github.com/synth-veena/pkg/generator"
	"github.com/synth-veena/pkg/raagas"
	"github.com/synth-veena/pkg/taalas"
	"github.com/synth-veena/pkg/types"
)

func main() {
	err := generator.Generate(generator.RandomGenerate, *raagas.MohanaRaagam, taalas.Rupaka, types.Kaala2nd, 180, 30)
	if err != nil {
		slog.Error("unable to generate" + err.Error())
		os.Exit(1)
	}	
}

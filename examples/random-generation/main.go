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
	err := generator.Generate(
		generator.RandomGenerate,
		*raagas.ThodiRaaga,
		taalas.Rupaka,
		types.Kaala3rd,
		180,
		5,
	)

	if err != nil {
		slog.Error("unable to generate" + err.Error())
		os.Exit(1)
	}
}

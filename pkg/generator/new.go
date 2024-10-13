package generator

import (
	"os"

	"github.com/google/uuid"
	"github.com/synth-veena/pkg/encoders"
	"github.com/synth-veena/pkg/types"
)

type GenerateFunc func(
	numberOfTaalas int,
	kaalam int,
	scale float32,
	raagam types.Raagam,
	taalam types.Taalam) (noteFrequencies []float32)

// type Generator struct {
// 	fn     GenerateFunc
// 	raagam types.Raagam
// 	taalam types.Taalam
// 	// tempo beats per minute
// 	tempo  int
// 	cycles int
// }

func Generate(
	fn GenerateFunc,
	raagam types.Raagam,
	taalam types.Taalam,
	kaalam int,
	tempo,
	cycles int) error {
	generatedFileName := raagam.Name + "-" + taalam.Name + uuid.NewString() + ".wav"
	outFile, err := os.Create(generatedFileName)
	if err != nil {
		return err
	}

	defer outFile.Close()
	encoder := encoders.NewWAVEncoder(outFile, tempo)
	notes := RandomGenerate(cycles, kaalam, types.CScale, raagam, taalam)

	if err := encoder.Encode(notes); err != nil {
		return err
	}

	return nil
}

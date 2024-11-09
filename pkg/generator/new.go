package generator

import (
	"bytes"
	"io"
	"os"

	"github.com/synth-veena/pkg/encoders"
	"github.com/synth-veena/pkg/types"
)

type GenerateFunc func(
	numberOfTaalas int,
	kaalam int,
	scale string,
	raagam types.Raagam,
	taalam types.Taalam) *types.SwaraLipi

func Generate(
	fn GenerateFunc,
	raagam types.Raagam,
	taalam types.Taalam,
	kaalam int,
	tempo,
	cycles int) error {
	fileName := raagam.Name + "-" + taalam.Name
	generatedWavFileName := fileName + ".wav"
	generatedLipiFileName := fileName + ".txt"
	outFile, err := os.Create(generatedWavFileName)
	if err != nil {
		return err
	}
	defer outFile.Close()

	outDocFile, err := os.Create(generatedLipiFileName)
	if err != nil {
		return err
	}

	defer outDocFile.Close()

	encoder := encoders.NewWAVEncoder(outFile, tempo, kaalam)
	lipi := RandomGenerate(cycles, kaalam, types.CScale, raagam, taalam)

	_, err = io.Copy(outDocFile, bytes.NewBufferString(lipi.Doc()))
	if err != nil {
		return err
	}

	for _, line := range lipi.Notes() {
		for _, freq := range line {
			if err := encoder.Encode(freq); err != nil {
				return err
			}
		}
	}

	return encoder.Close()
}

package generator

import (
	"bytes"
	"io"
	"os"

	"github.com/synth-veena/pkg/encoders"
	"github.com/synth-veena/pkg/instruments"
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
	tempo int,
	cycles int,
	instrument instruments.Instrument,
) error {
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

	encoder := encoders.NewWAVEncoder(outFile, tempo, kaalam, instrument)
	lipi := RandomGenerate(cycles, kaalam, types.CScale, raagam, taalam)

	_, err = io.Copy(outDocFile, bytes.NewBufferString(lipi.Doc()))
	if err != nil {
		return err
	}

	for _, line := range lipi.Notes() {
		if err := encoder.Encode(line); err != nil {
			return err
		}
	}

	return encoder.Close()
}

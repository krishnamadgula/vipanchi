package generator

import (
	"log/slog"
	"math/rand"

	"github.com/synth-veena/pkg/types"
)

/*
rules for generation

Each complete Taalam is a cycle
if it is a first iteration of taalam, starting note is SLower (basic)
else use the last note of last taalam to generate first note for new taalam

within a single taalam

	get first note
	generate the next note as current note -1, current note, current note +1
*/
func RandomGenerate(
	numberOfTaalas int,
	kaalam int,
	scale float32,
	raagam types.Raagam,
	taalam types.Taalam) []float32 {

	slog.Info("Synthesizing music in raagam: " + raagam.Name)

	output := []float32{}
	start := types.SLower
	for i := 0; i < numberOfTaalas*kaalam; i++ {
		output = append(output, randomSingleTaalam(taalam, raagam, start)...)
		start = randomNextNote(raagam, output[len(output)-1])
	}

	return output
}

func randomSingleTaalam(taalam types.Taalam, raagam types.Raagam, start float32) []float32 {
	output := []float32{}
	for i := range taalam.Sequence {
		for j := 0; j < taalam.Sequence[i]; j++ {
			cur := randomNextNote(raagam, start)
			output = append(output, cur)
			start = cur
		}
	}

	return output
}

func randomNextNote(raagam types.Raagam, previous float32) float32 {
	/*
		seed - 1234 (hardcoded)
		0 - same swara
		1 - aaroh / same swara
		2 - avroh / same swara
	*/
	const (
		same     = 0
		aarohOne = 1
		aarohTwo = 2
		avrohOne = 3
		avrohTwo = 4
	)

	num := rand.Intn(30000) % 5
	switch num {
	case same:
		return previous
	case aarohOne:
		if v, err := raagam.Aaroh(previous); err != nil {
			return previous
		} else {
			return v
		}

	case aarohTwo:
		if v, err := raagam.Aaroh(previous); err != nil {
			return previous
		} else {
			out, err := raagam.Aaroh(v)
			if err != nil {
				return v
			}

			return out
		}

	case avrohOne:
		if v, err := raagam.Avroh(previous); err != nil {
			return previous
		} else {
			return v
		}

	case avrohTwo:
		if v, err := raagam.Avroh(previous); err != nil {
			return previous
		} else {
			out, err := raagam.Avroh(v)
			if err != nil {
				return v
			}

			return out
		}

	default:
		return 0
	}
}

// func randomGenerate()

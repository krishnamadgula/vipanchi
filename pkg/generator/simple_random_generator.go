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
	scale string,
	raagam types.Raagam,
	taalam types.Taalam) *types.SwaraLipi {

	slog.Info("Synthesizing music in raagam: " + raagam.Name)

	lipi := types.NewSwaraLipi(&raagam, &taalam, kaalam)

	start := types.SLower
	for i := 0; i < numberOfTaalas*kaalam; i++ {
		last := randomSingleTaalam(lipi, taalam, raagam, start)

		start = randomNextNote(raagam, last)
	}

	return lipi
}

func randomSingleTaalam(
	lipi *types.SwaraLipi,
	taalam types.Taalam,
	raagam types.Raagam,
	start string) (endNote string) {
	for i := range taalam.Sequence {
		for j := 0; j < taalam.Sequence[i]; j++ {
			cur := randomNextNote(raagam, start)
			lipi.AppendMatra(cur)
			start = cur
		}
	}

	return start
}

func randomNextNote(raagam types.Raagam, previous string) string {
	direction := rand.Intn(5000) % 3
	step := rand.Intn(5000)%2 + 1
	switch direction {
	case 1:
		if v, err := raagam.Aaroh(previous, step); err != nil {
			return previous
		} else {
			return v
		}

	case 2:
		if v, err := raagam.Avroh(previous, step); err != nil {
			return previous
		} else {
			return v
		}

	default:
		return previous
	}
}

// func randomGenerate()

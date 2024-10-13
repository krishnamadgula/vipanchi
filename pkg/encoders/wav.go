package encoders

import (
	"io"
	"math"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

const KHz44 = 44_100
const Amplitude = 32767

type WAVEncoder struct {
	// tempo - beats per minute
	Tempo int
	e     *wav.Encoder
}

func NewWAVEncoder(output io.WriteSeeker, tempo int) *WAVEncoder {
	// TODO revisit encoder initialization, currently it is one time use
	e := wav.NewEncoder(output, KHz44, 16, 1, 1)
	return &WAVEncoder{e: e, Tempo: tempo}
}

func (w *WAVEncoder) Encode(frequencies []float32) error {
	for _, freq := range frequencies {
		data := generateSineWave(freq, w.Tempo, KHz44)
		err := w.e.Write(&audio.IntBuffer{Data: data, Format: audio.FormatMono44100, SourceBitDepth: 16})
		if err != nil {
			return err
		}
	}

	return w.e.Close()
}

func generateSineWave(freq float32, tempo, rate int) []int {
	secondsPerBeat := 60 / float32(tempo)
	numSamples := int(secondsPerBeat * float32(rate))
	data := make([]int, numSamples)
	for i := 0; i < numSamples; i++ {
		sampleVal := int(math.Sin(float64(2*math.Pi*freq*float32(i)/float32(rate))) * (Amplitude / 2))
		harmony2Val := int(math.Sin(float64(2*math.Pi*2*freq*float32(i)/float32(rate))) * (Amplitude / 5))
		harmony3Val := int(math.Sin(float64(2*math.Pi*3*freq*float32(i)/float32(rate))) * (Amplitude / 10))
		harmony4Val := int(math.Sin(float64(2*math.Pi*4*freq*float32(i)/float32(rate))) * (Amplitude / 20))
		harmony5Val := int(math.Sin(float64(2*math.Pi*5*freq*float32(i)/float32(rate))) * (Amplitude / 40))
		sampleVal += harmony2Val + harmony3Val + harmony4Val+harmony5Val
		if sampleVal > Amplitude {
			sampleVal = Amplitude
		} else if sampleVal < -Amplitude {
			sampleVal = -Amplitude
		}
		data[i] = sampleVal
	}

	return data
}

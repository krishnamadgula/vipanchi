package encoders

import (
	"io"
	"math"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"github.com/synth-veena/pkg/instruments"
)

const KHz44 = 44_100
const Amplitude = 32767

type WAVEncoder struct {
	// tempo - beats per minute
	Tempo      int
	Kaalam     int
	Instrument instruments.Instrument
	e          *wav.Encoder
}

func NewWAVEncoder(output io.WriteSeeker,
	tempo, Kaalam int) *WAVEncoder {
	// TODO revisit encoder initialization, currently it is one time use
	e := wav.NewEncoder(output, KHz44, 16, 1, 1)
	return &WAVEncoder{e: e, Tempo: tempo, Kaalam: Kaalam}
}

func (w *WAVEncoder) Close() error {
	return w.e.Close()
}

func (w *WAVEncoder) Encode(freq float32) error {

	// Generate sine wave data for each frequency in the line
	data := generateSineWave(freq, w.Tempo*w.Kaalam, KHz44)

	// Write the audio data to the encoder
	return w.e.Write(&audio.IntBuffer{
		Data:           data,
		Format:         audio.FormatMono44100,
		SourceBitDepth: 16,
	})
}

// generateSineWave generates a sine wave with harmonics for the given frequency and duration
// tempo is in beats per minute, rate is sample rate in Hz
func generateSineWave(freq float32, tempo, rate int) []int {
	// Calculate duration in seconds based on tempo (beats per minute)
	// For example, if tempo is 120 BPM, each beat is 0.5 seconds
	secondsPerBeat := 60.0 / float32(tempo)

	// Calculate total number of samples needed for this duration
	numSamples := int(secondsPerBeat * float32(rate))
	data := make([]int, numSamples)

	// Angular frequency (ω = 2πf)
	angularFreq := float64(2 * math.Pi * freq)

	// Generate samples
	for i := 0; i < numSamples; i++ {
		t := float64(i) / float64(rate) // Time at this sample

		// Fundamental frequency
		sample := math.Sin(angularFreq*t) * (Amplitude / 2)
		sample *= instruments.GetEnvelopeMultiplier(
			t,
			float64(numSamples)/float64(rate),
			instruments.Veena,
		)
		// Convert to integer and clip to prevent overflow
		sampleInt := int(sample)
		if sampleInt > Amplitude {
			sampleInt = Amplitude
		} else if sampleInt < -Amplitude {
			sampleInt = -Amplitude
		}

		data[i] = sampleInt
	}

	return data
}

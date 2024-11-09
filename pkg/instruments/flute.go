package instruments

// Flute represents the physical characteristics of a flute instrument
var Flute = Instrument{
	Harmonics: [][2]float64{
		{1, 1.0},  // Fundamental
		{2, 0.75}, // Strong 2nd harmonic
		{3, 0.5},  // Moderate 3rd harmonic
		{4, 0.25}, // Weak 4th harmonic
		{5, 0.1}}, // Very weak 5th harmonic
	Attack:  0.05, // Gentle attack
	Decay:   0.1,  // Short decay
	Sustain: 0.8,  // Strong sustain
	Release: 0.15, // Quick release
}

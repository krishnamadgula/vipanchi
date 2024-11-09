package instruments

// Veena represents the physical characteristics of a veena instrument
var Veena = Instrument{
	Harmonics: [][2]float64{
		{1, 1.0},    // Fundamental
		{2, 0.5},    // 2nd harmonic
		{3, 0.33},   // 3rd harmonic
		{4, 0.25},   // 4th harmonic
		{5, 0.2},    // 5th harmonic
		{6, 0.167},  // 6th harmonic
		{7, 0.143},  // 7th harmonic
		{8, 0.125}}, // 8th harmonic
	Attack:  0.01, // Quick attack
	Decay:   0.1,  // Short decay
	Sustain: 0.7,  // Strong sustain
	Release: 0.2,  // Medium release
}

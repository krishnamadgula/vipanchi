package instruments

// Harmonica represents the physical characteristics of a harmonica
var Harmonica = Instrument{
	Harmonics: [][2]float64{
		{1, 1.0},  // Fundamental
		{2, 0.85}, // Very strong 2nd harmonic - gives harmonica its bright, reedy tone
		{3, 0.7},  // Strong 3rd harmonic - adds warmth
		{4, 0.4},  // Moderate 4th harmonic
		{5, 0.3},  // Moderate 5th harmonic
		{6, 0.2},  // Weak 6th harmonic
		{7, 0.1}}, // Very weak 7th harmonic
	Attack:  0.03, // Quick attack due to direct air flow
	Decay:   0.1,  // Short decay
	Sustain: 0.8,  // High sustain while air flows
	Release: 0.15, // Quick release when air stops
	Overlap: 0.4,  // Moderate overlap for legato effect
}

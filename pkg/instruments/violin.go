package instruments

// Violin represents the physical characteristics of a violin instrument
var Violin = Instrument{
	Harmonics: [][2]float64{
		{1, 1.0},  // Fundamental
		{2, 0.85}, // Strong 2nd harmonic - gives violin its bright tone
		{3, 0.7},  // Prominent 3rd harmonic
		{4, 0.5},  // Moderate 4th harmonic
		{5, 0.4},  // Moderate 5th harmonic
		{6, 0.3},  // Weaker 6th harmonic
		{7, 0.2}}, // Weak 7th harmonic
	Attack:  0.03, // Quick attack
	Decay:   0.1,  // Short decay
	Sustain: 0.85, // Strong sustain for long bow strokes
	Release: 0.15, // Medium-quick release
	Overlap: 0.3,  // Moderate overlap for legato bowing effect
}

package instruments

// Shenai represents the physical characteristics of a shenai (Indian double-reed woodwind) instrument
var Shenai = Instrument{
	Harmonics: [][2]float64{
		{1, 1.0},  // Fundamental
		{2, 0.85}, // Very strong 2nd harmonic - gives shenai its piercing quality
		{3, 0.7},  // Strong 3rd harmonic - adds brightness
		{4, 0.5},  // Moderate 4th harmonic
		{5, 0.35}, // Moderate 5th harmonic
		{6, 0.2},  // Weak 6th harmonic
		{7, 0.1}}, // Very weak 7th harmonic
	Attack:  0.04, // Quick attack characteristic of double-reed instruments
	Decay:   0.1,  // Short decay
	Sustain: 0.85, // Strong sustain for long melodic phrases
	Release: 0.12, // Moderate release
	Overlap: 0.25, // Moderate overlap for smooth transitions
}

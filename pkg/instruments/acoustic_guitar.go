package instruments	
// AcousticGuitar represents the physical characteristics of an acoustic guitar
var AcousticGuitar = Instrument{
	Harmonics: [][2]float64{
		{1, 1.0},   // Fundamental
		{2, 0.7},   // Strong 2nd harmonic - adds warmth
		{3, 0.5},   // Moderate 3rd harmonic
		{4, 0.3},   // Moderate 4th harmonic
		{5, 0.25},  // Weak 5th harmonic
		{6, 0.2},   // Weak 6th harmonic
		{7, 0.15}}, // Very weak 7th harmonic
	Attack:  0.02,  // Quick attack
	Decay:   0.15,  // Short-moderate decay
	Sustain: 0.7,   // Medium-high sustain level
	Release: 0.25,  // Moderate release
	Overlap: 0.1,   // Slight overlap between notes for smooth transitions
}

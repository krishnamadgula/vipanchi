package instruments
// ElectricGuitar represents the physical characteristics of an electric guitar
var ElectricGuitar = Instrument{
	Harmonics: [][2]float64{
		{1, 1.0},   // Fundamental
		{2, 0.6},   // Strong 2nd harmonic - adds brightness
		{3, 0.4},   // Moderate 3rd harmonic 
		{4, 0.3},   // Moderate 4th harmonic
		{5, 0.2},   // Weaker 5th harmonic
		{6, 0.15},  // Weak 6th harmonic
		{7, 0.1}},  // Very weak 7th harmonic
	Attack:  0.01,  // Very quick attack
	Decay:   0.2,   // Moderate decay
	Sustain: 0.65,  // Medium sustain level
	Release: 0.3,   // Moderate release
}

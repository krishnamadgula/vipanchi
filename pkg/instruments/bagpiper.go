package instruments

// Bagpiper represents the physical characteristics of a bagpipe instrument
var Bagpiper = Instrument{
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
	Overlap: 0.2,  // Moderate overlap for legato effect
}

// Bagpiper represents the physical characteristics of a bagpipe instrument
// var Bagpiper = Instrument{
// 	Harmonics: [][2]float64{
// 		{1, 1.0},   // Fundamental
// 		{2, 0.8},   // Strong 2nd harmonic - gives bagpipe its distinctive drone
// 		{3, 0.65},  // Strong 3rd harmonic - adds richness
// 		{4, 0.45},  // Moderate 4th harmonic
// 		{5, 0.35},  // Moderate 5th harmonic
// 		{6, 0.25},  // Weak 6th harmonic
// 		{7, 0.15}}, // Very weak 7th harmonic
// 	Attack:  0.08, // Moderate attack due to air pressure buildup
// 	Decay:   0.05, // Very short decay since bagpipes maintain constant pressure
// 	Sustain: 0.95, // Very high sustain - bagpipes known for continuous sound
// 	Release: 0.1,  // Quick release when air pressure stops
// 	Overlap: 0.35, // Moderate overlap for legato effect
// }

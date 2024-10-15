package raagas

import "github.com/synth-veena/pkg/types"

/*
S (C4): 261.63 Hz
R2 (D4): 293.66 Hz
G3 (E4): 329.63 Hz
P (G4): 392.00 Hz
D2 (A4): 440.00 Hz
S (C5): 523.25 Hz
*/
var MohanaRaagam = types.NewSimpleRaagam("Mohana", []string{
	types.SLower,
	types.R2,
	types.G3,
	types.P,
	types.D2,
	types.SUpper,
})

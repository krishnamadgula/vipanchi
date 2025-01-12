package raagas

import "github.com/synth-veena/pkg/types"

/*
S (C4): 261.63 Hz
R2 (D4): 293.66 Hz
G3 (E4): 329.63 Hz
M1 (F4): 349.23 Hz
P (G4): 392.00 Hz
D2 (A4): 440.00 Hz
N2 (B4): 493.88 Hz
S (C5): 523.25 Hz
*/
var ChakravaakaRaagam = types.NewSimpleRaagam("Chakravaaka", []string{
	types.SLower,
	types.R2,
	types.G3,
	types.M1,
	types.P,
	types.D2,
	types.N2,
	types.SUpper,
})

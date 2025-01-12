package raagas
import "github.com/synth-veena/pkg/types"

/*
S (C4): 261.63 Hz
R2 (D4): 293.66 Hz
G3 (E4): 329.63 Hz
P (G4): 392.00 Hz
D1 (Ab4): 415.30 Hz
S (C5): 523.25 Hz
*/
var ShivaranjaniRaagam = types.NewSimpleRaagam("Shivaranjani", []string{
	types.SLower,
	types.R2,
	types.G3,
	types.P,
	types.D1,
	types.SUpper,
})

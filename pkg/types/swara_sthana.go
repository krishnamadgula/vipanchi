package types

/*
S (C4) – 261.63 Hz
R₁ (D♭4) – 277.18 Hz
G₃ (E4) – 329.63 Hz
M₁ (F4) – 349.23 Hz
P (G4) – 392.00 Hz
D₁ (A♭4) – 415.30 Hz
N₃ (B4) – 493.88 Hz
*/

var Swaras = map[string]float32{
	SLower: SLowerFreq,
	R1:     R1Freq,
	R2:     R2Freq,
	G2:     G2Freq,
	G3:     G3Freq,
	M1:     M1Freq,
	M2:     M2Freq,
	P:      PFreq,
	D1:     D1Freq,
	D2:     D2Freq,
	N2:     N2Freq,
	N3:     N3Freq,
	SUpper: SUpperFreq,
}

const (
	SLowerFreq = float32(261.63)
	SLower     = "SLower"
	R1Freq     = float32(277.18)
	R1         = "R1"
	R2Freq     = float32(293.66)
	R2         = "R2"

	G2Freq = float32(311.13)
	G2     = "G2"
	G3Freq = float32(329.63)
	G3     = "G3"

	M1Freq = float32(349.23)
	M1     = "M1"
	M2Freq = float32(369.99)
	M2     = "M2"

	PFreq = float32(392)
	P     = "P"

	D1Freq = float32(415.30)
	D1     = "D1"
	D2Freq = float32(440)
	D2     = "D2"

	N2Freq     = float32(466.16)
	N2         = "N2"
	N3Freq     = float32(493.88)
	N3         = "N3"
	SUpperFreq = float32(523.25)
	SUpper     = "SUpper"
)
